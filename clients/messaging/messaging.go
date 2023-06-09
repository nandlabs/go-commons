package messaging

type Producer interface {
	Send(string, Message) error
	SendBatch(string, ...Message) error
}

type Consumer interface {
	// TODO :: function as parameter execution on success or error
	OnMessage() error
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
