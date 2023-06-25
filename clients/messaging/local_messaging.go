package messaging

import (
	"fmt"
	"net/url"
)

var (
	localMessagingSchemes = []string{"chan"}
	localMsgChannel       = make(chan Message)
)

type LocalMessagingSystem struct {
	LocalMessage
}

func (lms *LocalMessagingSystem) Send(_ *url.URL, msg Message) (err error) {
	go func() {
		localMsgChannel <- msg
	}()
	return
}

func (lms *LocalMessagingSystem) SendBatch(destination *url.URL, messages ...Message) (err error) {
	for _, message := range messages {
		err = lms.Send(destination, message)
		if err != nil {
			return
		}
	}
	return
}

func (lms *LocalMessagingSystem) OnMessage(_ *url.URL, onConsume ConsumeMessage) (err error) {
	for msg := range localMsgChannel {
		fmt.Println(msg)
		err = onConsume(msg)
		if err != nil {
			return
		}
	}
	return
}

func (lms *LocalMessagingSystem) Schemes() []string {
	return localMessagingSchemes
}
