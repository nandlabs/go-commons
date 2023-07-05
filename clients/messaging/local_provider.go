package messaging

//
//import (
//	"errors"
//	"net/url"
//)
//
//var (
//	localMsgChannel = make(chan Message)
//)
//
//// localProvider is an implementation of the Provider interface
//type localProvider struct{}
//
//func (lp *localProvider) AddListener(url *url.URL, listener func(msg Message)) (err error) {
//	return
//}
//
//func (lp *localProvider) Send(url *url.URL, msg Message) (err error) {
//	if url.Scheme != "chan" {
//		err = errors.New("invalid provider url " + url.String())
//	}
//	go func() {
//		localMsgChannel <- msg
//	}()
//	return
//}
//
//func (lp *localProvider) SendBatch(url *url.URL, msgs ...Message) (err error) {
//	if url.Scheme != "chan" {
//		err = errors.New("invalid provider url " + url.String())
//	}
//	for _, message := range msgs {
//		err = lp.Send(url, message)
//		if err != nil {
//			return
//		}
//	}
//	return
//}
//
//func (lp *localProvider) Receive(url *url.URL) (msg Message, err error) {
//	if url.Scheme != "chan" {
//		err = errors.New("invalid provider url " + url.String())
//	}
//	for m := range localMsgChannel {
//		msg = m
//	}
//	return
//}
//
//func (lp *localProvider) ReceiveBatch(url *url.URL) (msgs []Message, err error) {
//	if url.Scheme != "chan" {
//		err = errors.New("invalid provider url " + url.String())
//	}
//	for m := range localMsgChannel {
//		msgs = append(msgs, m)
//	}
//	return
//}
//
//func (lp *localProvider) Setup() {
//	localProvider := &localProvider{}
//	uri, err := url.Parse("chan://localhost:8080")
//	if err == nil {
//		Register(uri, localProvider)
//	}
//}
//
//func init() {
//	lp := &localProvider{}
//	lp.Setup()
//}
