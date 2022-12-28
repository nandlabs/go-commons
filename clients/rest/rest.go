package rest

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/base64"
	"net/http"
	"net/url"
	"os"
	"time"

	"go.nandlabs.io/commons/clients"
	"go.nandlabs.io/commons/config"
	"go.nandlabs.io/commons/fnutils"
	"go.nandlabs.io/commons/textutils"
)

const (
	defaultReqTimeout            = 60
	defaultMaxIdleConnections    = 100
	defaultIdleConnTimeout       = 90 * time.Second
	defaultTLSHandshakeTimeout   = 10 * time.Second
	defaultExpectContinueTimeout = 1 * time.Second
	contentTypeHdr               = "Content-Type"
	proxyAuthHdr                 = "Proxy-Authorization"
)

// TODO Add certificate
type Client struct {
	retryInfo      *clients.RetryInfo
	circuitBreaker *clients.CircuitBreaker
	errorOnMap     map[int]int
	proxyBasicAuth string
	httpClient     http.Client
	httpTransport  *http.Transport
	tlsConfig      *tls.Config
	codecOptions   map[string]interface{}
}

// NewClient function generates a new client with default values
func NewClient() *Client {
	transport := &http.Transport{
		MaxIdleConns:          defaultMaxIdleConnections,
		IdleConnTimeout:       defaultIdleConnTimeout,
		ExpectContinueTimeout: defaultExpectContinueTimeout,
		TLSHandshakeTimeout:   defaultTLSHandshakeTimeout,
	}
	httpClient := http.Client{
		Transport: transport,
		Timeout:   time.Duration(defaultReqTimeout) * time.Second,
	}

	return &Client{
		httpClient:    httpClient,
		httpTransport: transport,
	}
}

// ReqTimeout function sets the overall client timeout for a request.
// The default value is 60 seconds
func (c *Client) ReqTimeout(t uint) *Client {
	c.httpClient.Timeout = time.Duration(t) * time.Second
	return c
}

// AddCodecOption sets the option for codec to be used
func (c *Client) AddCodecOption(k string, v interface{}) *Client {
	if c.codecOptions == nil {
		c.codecOptions = make(map[string]interface{})
	}
	c.codecOptions[k] = v

	return c
}

// IdleTimeout sets is the maximum amount of time a conn can stay idle (keep-alive) before closing itself
func (c *Client) IdleTimeout(t uint) *Client {
	c.httpTransport.IdleConnTimeout = time.Duration(t) * time.Second
	return c
}

// ErrorOnHttpStatus sets the list of status codes that can be considered failures. This is useful for
// QualityOfService features like CircuitBreaker
func (c *Client) ErrorOnHttpStatus(statusCodes ...int) *Client {
	if c.errorOnMap == nil {
		c.errorOnMap = make(map[int]int)
	}
	for _, code := range statusCodes {
		c.errorOnMap[code] = code
	}
	return c
}

// MaxIdle sets the max idle connections that can stay idle (keep-alive).
func (c *Client) MaxIdle(maxIdleConn int) *Client {
	c.httpTransport.MaxIdleConns = maxIdleConn
	return c
}

// MaxIdlePerHost sets the max idle connections that can stay idle (keep-alive) for a given hostname
func (c *Client) MaxIdlePerHost(maxIdleConnPerHost int) *Client {
	c.httpTransport.MaxIdleConnsPerHost = maxIdleConnPerHost
	return c
}

// SSlVerify set the ssl verify value
func (c *Client) SSlVerify(verify bool) (*Client, error) {
	conf, err := c.setTlSConfig()
	if err != nil {
		return nil, err
	}
	conf.InsecureSkipVerify = verify
	return c, nil
}

func (c *Client) SetProxy(proxyUrl, user, password string) (err error) {
	var u *url.URL
	if user != textutils.EmptyStr && password != textutils.EmptyStr {
		c.proxyBasicAuth = "Basic " + base64.StdEncoding.EncodeToString([]byte(user+":"+password))
	}
	u, err = url.Parse(proxyUrl)
	if err == nil {
		c.httpTransport = &http.Transport{
			Proxy: http.ProxyURL(u),
		}
	}
	return
}

func (c *Client) SetCACerts(caFilePath ...string) (*Client, error) {
	conf, err := c.setTlSConfig()
	if err != nil {
		return nil, err
	}
	if conf.RootCAs == nil {
		conf.RootCAs = x509.NewCertPool()
	}
	for _, v := range caFilePath {
		caCert, err := os.ReadFile(v)
		if err != nil {
			return nil, err
		}
		conf.RootCAs.AppendCertsFromPEM(caCert)
	}
	c.setSSL(conf)
	return c, nil
}

func (c *Client) SetTLSCerts(certs ...tls.Certificate) (*Client, error) {
	conf, err := c.setTlSConfig()
	if err != nil {
		return nil, err
	}
	for _, v := range certs {
		conf.Certificates = append(conf.Certificates, v)
	}
	c.setSSL(conf)
	return c, nil
}

func (c *Client) setSSL(conf *tls.Config) {
	// Load client cert
	c.tlsConfig = conf
	transport := &http.Transport{
		TLSClientConfig: conf,
	}
	c.httpTransport = transport
}

func (c *Client) UseEnvProxy(urlParam, userParam, passwdParam string) (err error) {
	u := config.GetEnvAsString(urlParam, textutils.EmptyStr)
	user := config.GetEnvAsString(userParam, textutils.EmptyStr)
	pass := config.GetEnvAsString(passwdParam, textutils.EmptyStr)
	err = c.SetProxy(u, user, pass)
	return
}

// Retry sets the maximum number of retries and wait interval in seconds between retries.
// The client does not retry by default. If retry configuration is set along with UseCircuitBreaker then the retry config
// is ignored
func (c *Client) Retry(maxRetries, wait int) *Client {
	c.retryInfo = &clients.RetryInfo{
		MaxRetries: maxRetries,
		Wait:       wait,
	}
	return c
}

// UseCircuitBreaker sets the circuit breaker configuration for this client.
// The circuit breaker pattern has higher precedence than retry pattern. If both are set then the retry configuration is
// ignored.
func (c *Client) UseCircuitBreaker(failureThreshold, successThreshold uint64, maxHalfOpen, timeout uint32) *Client {
	breakerInfo := &clients.BreakerInfo{
		FailureThreshold: failureThreshold,
		SuccessThreshold: successThreshold,
		MaxHalfOpen:      maxHalfOpen,
		Timeout:          timeout,
	}
	c.circuitBreaker = clients.NewCB(breakerInfo)

	return c
}

func (c *Client) NewRequest(url, method string) *Request {
	return &Request{
		url:    url,
		method: method,
		header: map[string][]string{},
		client: c,
	}
}

// Execute the client request and get the response object
func (c *Client) Execute(req *Request) (res *Response, err error) {
	var httpReq *http.Request
	var httpRes *http.Response
	httpReq, err = req.toHttpRequest()
	if c.proxyBasicAuth != "" {
		httpReq.Header.Set(proxyAuthHdr, c.proxyBasicAuth)
	}
	if err == nil {
		if c.circuitBreaker != nil {
			//Use Circuit Breaker
			err = c.circuitBreaker.CanExecute()
			if err == nil {
				httpRes, err = c.httpClient.Do(httpReq)
				c.circuitBreaker.OnExecution(c.isError(err, httpRes))
			}
		} else if c.retryInfo != nil {
			httpRes, err = c.httpClient.Do(httpReq)

			for i := 0; c.isError(err, httpRes) && i < c.retryInfo.MaxRetries; i++ {
				fnutils.ExecuteAfterSecs(func() {
					httpRes, err = c.httpClient.Do(httpReq)
				}, c.retryInfo.Wait)
			}
		} else {
			httpRes, err = c.httpClient.Do(httpReq)
		}
		if err == nil {
			res = &Response{raw: httpRes, client: c}
		}
	}
	return
}

// Check if the response is an error response or an error has been received
func (c *Client) isError(err error, httpRes *http.Response) (isErr bool) {
	isErr = err != nil
	if !isErr && c.errorOnMap != nil {
		_, isErr = c.errorOnMap[httpRes.StatusCode]
	}
	return
}

// Close function with close all idle connections that are available
func (c *Client) Close() (err error) {
	c.httpClient.CloseIdleConnections()
	return
}

func (c *Client) setTlSConfig() (*tls.Config, error) {
	if c.tlsConfig != nil {
		return c.tlsConfig, nil
	}
	return &tls.Config{}, nil
}
