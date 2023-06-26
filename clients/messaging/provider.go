package messaging

import "net/url"

type Producer interface {
	Send(*url.URL, Message) error
	SendBatch(*url.URL, ...Message) error
}

type Receiver interface {
	Receive(*url.URL) (Message, error)
	ReceiveBatch(*url.URL) ([]Message, error)
	AddListener(*url.URL, func(msg Message)) error
}

type Provider interface {
	Producer
	Receiver
	Setup()
}
