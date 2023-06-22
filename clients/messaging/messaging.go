package messaging

import "net/url"

type Producer interface {
	Send(*url.URL, *Message) error
	SendBatch(*url.URL, ...*Message) error
}

type ConsumeMessage func(msg *Message) error

type Consumer interface {
	OnMessage(*url.URL, ConsumeMessage) error
}

type Messaging interface {
	Producer
	Consumer
	Schemes() []string
}

type ManagerMessaging interface {
	Messaging
	Register(messaging Messaging)
	IsSupported(scheme string) bool
}
