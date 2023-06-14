package messaging

type Producer interface {
	Send(string, LocalMessage) error
	SendBatch(string, ...LocalMessage) error
}

type Consumer interface {
	// TODO :: function execution on success or error
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
