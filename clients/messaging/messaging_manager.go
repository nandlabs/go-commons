package messaging

import (
	"go.nandlabs.io/commons/errutils"
	"net/url"
	"sync"
)

var manager ManagerMessaging

type messagingSystems struct {
	mutex            sync.Mutex
	messagingSystems map[string]Messaging
}

func (ms *messagingSystems) Send(destination *url.URL, msg Message) (err error) {
	var messaging Messaging
	messaging, err = ms.getMsFor(destination)
	if err == nil {
		err = messaging.Send(destination, msg)
	}
	return
}

func (ms *messagingSystems) SendBatch(destination *url.URL, msg ...Message) (err error) {
	var messaging Messaging
	messaging, err = ms.getMsFor(destination)
	if err == nil {
		err = messaging.SendBatch(destination, msg...)
	}
	return
}

func (ms *messagingSystems) OnMessage(source *url.URL, onConsume ConsumeMessage) (err error) {
	var messaging Messaging
	messaging, err = ms.getMsFor(source)
	if err == nil {
		err = messaging.OnMessage(source, onConsume)
	}
	return
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

func (ms *messagingSystems) IsSupported(scheme string) (supported bool) {
	_, supported = ms.messagingSystems[scheme]
	return
}

func (ms *messagingSystems) getMsFor(src *url.URL) (msg Messaging, err error) {
	var ok bool
	msg, ok = ms.messagingSystems[src.Scheme]
	if !ok {
		err = errutils.FmtError("unsupported messaging scheme %s for in the url %s", src.Scheme, src.String())
	}
	return
}

func init() {
	manager = &messagingSystems{}
	localMs := newLocalMessagingSystem()
	manager.Register(localMs)
}

func newLocalMessagingSystem() (msg *LocalMessagingSystem) {
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
