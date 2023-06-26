package messaging

import (
	"errors"
	"net/url"
)

var (
	providers = make(map[string]Provider)
)

// Messaging is a wrapper on the Provider interface
// TODO :: should Messaging implement the Provider interface?
type Messaging struct {
	// TODO :: this should contain the circuit_breaker and retry info
}

func Register(url *url.URL, provider Provider) {
	if providers == nil {
		providers = make(map[string]Provider)
	}
	providers[url.String()] = provider
}

func (m *Messaging) AddListener(url *url.URL, listener func(msg Message)) (err error) {
	var provider Provider
	provider, err = m.getProvider(url)
	if err == nil {
		err = provider.AddListener(url, listener)
	}
	return
}

func (m *Messaging) getProvider(url *url.URL) (provider Provider, err error) {
	supports := false
	provider, supports = providers[url.String()]
	if !supports {
		err = errors.New("unsupported provider with url " + url.String())
	}
	return
}

func (m *Messaging) Send(url *url.URL, msg Message) (err error) {
	var provider Provider
	provider, err = m.getProvider(url)
	if err == nil {
		err = provider.Send(url, msg)
	}
	return
}

func (m *Messaging) SendBatch(url *url.URL, msg ...Message) (err error) {
	var provider Provider
	provider, err = m.getProvider(url)
	if err == nil {
		err = provider.SendBatch(url, msg...)
	}
	return
}

func (m *Messaging) Receive(url *url.URL) (msg Message, err error) {
	var provider Provider
	provider, err = m.getProvider(url)
	if err == nil {
		msg, err = provider.Receive(url)
	}
	return
}

func (m *Messaging) ReceiveBatch(url *url.URL) (msgs []Message, err error) {
	var provider Provider
	provider, err = m.getProvider(url)
	if err == nil {
		msgs, err = provider.ReceiveBatch(url)
	}
	return
}
