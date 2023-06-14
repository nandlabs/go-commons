package messaging

var (
	localMessagingSchemes = []string{"abc", "efg"}
)

type LocalMessagingSystem struct {
}

func (lms *LocalMessagingSystem) Send(destination string, msg LocalMessage) error {
	return nil
}

func (lms *LocalMessagingSystem) SendBatch(destination string, msg ...LocalMessage) error {
	return nil
}

func (lms *LocalMessagingSystem) OnMessage() error {
	return nil
}

func (lms *LocalMessagingSystem) Schemes() []string {
	return localMessagingSchemes
}
