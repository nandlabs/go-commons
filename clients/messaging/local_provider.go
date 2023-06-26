package messaging

import (
	"errors"
	"net/url"
)

var (
	localMsgChannel = make(chan Message)
)

// LocalProvider is an implementation of the Provider interface
type LocalProvider struct{}

func (lp *LocalProvider) AddListener(url *url.URL, listener func(msg Message)) (err error) {
	return
}

func (lp *LocalProvider) Send(url *url.URL, msg Message) (err error) {
	if url.Scheme != "chan" {
		err = errors.New("invalid provider url " + url.String())
	}
	go func() {
		localMsgChannel <- msg
	}()
	return
}

func (lp *LocalProvider) SendBatch(url *url.URL, msgs ...Message) (err error) {
	if url.Scheme != "chan" {
		err = errors.New("invalid provider url " + url.String())
	}
	for _, message := range msgs {
		err = lp.Send(url, message)
		if err != nil {
			return
		}
	}
	return
}

func (lp *LocalProvider) Receive(url *url.URL) (msg Message, err error) {
	if url.Scheme != "chan" {
		err = errors.New("invalid provider url " + url.String())
	}
	for m := range localMsgChannel {
		msg = m
	}
	return
}

func (lp *LocalProvider) ReceiveBatch(url *url.URL) (msgs []Message, err error) {
	if url.Scheme != "chan" {
		err = errors.New("invalid provider url " + url.String())
	}
	for m := range localMsgChannel {
		msgs = append(msgs, m)
	}
	return
}

func (lp *LocalProvider) Setup() {
	localProvider := &LocalProvider{}
	uri, err := url.Parse("chan://localhost:8080")
	if err != nil {
		Register(uri, localProvider)
	}
}

func init() {
	lp := &LocalProvider{}
	lp.Setup()
}
