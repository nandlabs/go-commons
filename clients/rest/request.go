package rest

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/textproto"
	"net/url"
	"os"
	"path/filepath"
	"strings"

	"go.nandlabs.io/commons/codec"
	"go.nandlabs.io/commons/errutils"
	"go.nandlabs.io/commons/ioutils"
	"go.nandlabs.io/commons/textutils"
)

const (
	pathParamPrefix = "${"
	pathParamSuffix = "}"
)

// Request struct holds the http Request for the rest client
// TODO Add multipart implementation
type Request struct {
	url           string
	method        string
	formData      url.Values
	queryParam    url.Values
	pathParams    map[string]string
	header        http.Header
	body          any
	bodyReader    io.Reader
	contentType   string
	client        *Client
	isMultipart   bool
	multipartFile *MultipartFile
}

type MultipartFile struct {
	ParamName string
	FilePath  string
}

// Method function prints the current method for this Request
func (r *Request) Method() string {
	return r.method
}

// AddFormData function adds the form data with the name specified by k list of values in order as specified in v
// If the key does not exist then it creates a new form data by calling url.Values.Set() function on the first key and
// the value
// Setting form data will have precedence over to setting body directly.
func (r *Request) AddFormData(k string, v ...string) *Request {
	if r.formData == nil {
		r.formData = url.Values{}
	}
	for _, s := range v {
		if r.formData.Has(k) {
			r.formData.Add(k, s)
		} else {
			r.formData.Set(k, s)
		}
	}
	return r
}

// AddQueryParam function adds the query parameter with the name specified by k list of values in order as specified in v
// If the key does not exist then it creates a new form data by calling url.Values.Set() function passing the first key
// and value
func (r *Request) AddQueryParam(k string, v ...string) *Request {
	if r.queryParam == nil {
		r.queryParam = url.Values{}
	}
	for _, s := range v {
		if r.queryParam.Has(k) {
			r.queryParam.Add(k, s)
		} else {
			r.queryParam.Set(k, s)
		}
	}
	return r
}

func (r *Request) AddHeader(k string, v ...string) *Request {
	mh := textproto.MIMEHeader(r.header)
	for i, s := range v {
		if i == 0 {
			if _, ok := mh[k]; !ok {
				mh.Set(k, s)
			}
		} else {
			mh.Add(k, s)
		}

	}
	return r
}

func (r *Request) SetBody(v interface{}) *Request {
	r.body = v
	return r
}
func (r *Request) SeBodyReader(reader io.Reader) *Request {
	r.bodyReader = reader
	return r
}

func (r *Request) SetContentType(contentType string) *Request {
	r.contentType = contentType
	return r
}

func (r *Request) SetMultipartFile(paramName string, path string) *Request {
	r.isMultipart = true
	r.multipartFile = &MultipartFile{
		ParamName: paramName,
		FilePath:  path,
	}
	return r
}

func (r *Request) handleMultipart() (err error) {
	bodyBuf := new(bytes.Buffer)
	w := multipart.NewWriter(bodyBuf)
	err = addFile(w, r.multipartFile.ParamName, r.multipartFile.FilePath)

	err = w.Close()
	return
}

func addFile(w *multipart.Writer, fieldName, path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer closeInterface(file)
	return writeMultipartFormFile(w, fieldName, filepath.Base(path), file)
}

func (r *Request) toHttpRequest() (httpReq *http.Request, err error) {
	var u *url.URL
	u, err = url.Parse(r.url)

	if err == nil {
		//path := u.Path
		if strings.Contains(u.Path, pathParamPrefix) {
			pathValues := strings.Split(u.Path, textutils.ForwardSlashStr)
			for i := range pathValues {
				l := len(pathValues[i])
				if l > 3 && strings.HasPrefix(pathValues[i], pathParamPrefix) &&
					strings.HasSuffix(pathValues[l-1], pathParamSuffix) {
					key := pathValues[i][2:l]
					if v, ok := r.pathParams[key]; ok {
						pathValues[i] = v
					} else {
						err = errutils.FmtError("Path param with name %s is not set in the request ", key)
						break
					}
				}
			}
			path := ""
			for i, pv := range pathValues {
				if i != 0 {
					path = path + textutils.ForwardSlashStr
				}
				path = path + pv
			}
			u.Path = path
		}

		if err == nil {

			if r.formData != nil {
				r.bodyReader = strings.NewReader(r.formData.Encode())
			}

			if r.bodyReader == nil && r.body != nil {
				pr, pw := io.Pipe()
				go func() {
					defer ioutils.CloserFunc(pw)
					var c codec.Codec
					c, err = codec.Get(r.contentType, r.client.codecOptions)
					if err == nil {
						err = c.Write(r.body, pw)
					}
				}()
				r.bodyReader = pr
			}

			if r.isMultipart {
				err = r.handleMultipart()
			}

			if err == nil {
				fmt.Println(r.bodyReader)
				httpReq, err = http.NewRequest(r.method, u.String(), r.bodyReader)
				if r.header != nil {
					if r.contentType != "" {
						r.header.Set(contentTypeHdr, r.contentType)
					}
					httpReq.Header = r.header
				}
			}
		}
	}
	return
}
