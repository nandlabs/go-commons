package clients

import (
	"sync"
	"time"
)

type Client interface {
	Execute(req any) (res any, err error)
}

type clientInfo struct {
	circuitBreaker *CircuitBreaker
	retryHandler   *RetryInfo
	client         Client
}

func (c *clientInfo) Execute(req any) (res any, err error) {
	if c.circuitBreaker != nil {
		err = c.circuitBreaker.CanExecute()
		if err == nil {
			res, err = c.client.Execute(req)
			c.circuitBreaker.OnExecution(err == nil)
		}
	} else if c.retryHandler == nil {
		res, err = c.client.Execute(req)
		if err != nil {
			for i := 0; i < c.retryHandler.MaxRetries; i++ {
				time.Sleep(time.Second * time.Duration(c.retryHandler.Wait))
				res, err = c.client.Execute(req)
				if err == nil {
					break
				}
			}
		}
	}
	return
}

type ClientManager struct {
	sync.RWMutex
	clients map[string]*clientInfo
}

var Manager = &ClientManager{
	clients: make(map[string]*clientInfo),
}

func (cm *ClientManager) GetClient(id string) Client {
	cm.RLock()
	defer cm.RUnlock()
	return cm.clients[id]
}

func (cm *ClientManager) Register(id string, c Client, retryInfo *RetryInfo, breakerInfo *BreakerInfo) {
	cm.Lock()
	defer cm.Unlock()
	ci := &clientInfo{
		circuitBreaker: NewCB(breakerInfo),
		retryHandler:   retryInfo,
		client:         c,
	}
	cm.clients[id] = ci
}
