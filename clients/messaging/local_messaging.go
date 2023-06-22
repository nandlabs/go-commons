package messaging

import "net/url"

var (
	localMessagingSchemes = []string{"chan"}
	localMsgChannel       = make(chan *Message)
)

type LocalMessagingSystem struct {
	LocalMessage
}

func (lms *LocalMessagingSystem) Send(_ *url.URL, msg *Message) error {
	localMsgChannel <- msg
	return nil
}

func (lms *LocalMessagingSystem) SendBatch(_ *url.URL, messages ...*Message) error {
	for _, message := range messages {
		localMsgChannel <- message
	}
	return nil
}

func (lms *LocalMessagingSystem) OnMessage(source *url.URL, onConsume ConsumeMessage) (err error) {
	for msg := range localMsgChannel {
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
