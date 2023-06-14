package messaging

import (
	"sync"
)

var manager ManagerMessaging

type messagingSystems struct {
	mutex            sync.Mutex
	messagingSystems map[string]Messaging
}

func (ms *messagingSystems) Send(destination string, msg LocalMessage) error {
	return nil
}

func (ms *messagingSystems) SendBatch(destination string, msg ...LocalMessage) error {
	return nil
}

func (ms *messagingSystems) OnMessage() error {
	return nil
}

func (ms *messagingSystems) Schemes() (schemes []string) {
	for k := range ms.messagingSystems {
		if k == "" {
			continue
		}
		schemes = append(schemes, k)
	}
	return
}

func init() {
	manager := &messagingSystems{}
	localMs := newLocalMessagingSystem()
	manager.Register(localMs)
}

func newLocalMessagingSystem() (msg Messaging) {
	return &LocalMessagingSystem{}
}

func (ms *messagingSystems) Register(messaging Messaging) {
	ms.mutex.Lock()
	ms.mutex.Unlock()
	for _, s := range messaging.Schemes() {
		if ms.messagingSystems == nil {
			ms.messagingSystems = make(map[string]Messaging)
		}
		ms.messagingSystems[s] = messaging
	}
}

func GetMessagingManager() ManagerMessaging {
	return manager
}
