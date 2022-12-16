package clients

import "time"

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

type Manager struct {
	clients map[string]*clientInfo
}

func (m *Manager) Get(id string) Client {
	return m.clients[id]
}

func (m *Manager) Register(id string, c Client, retryInfo *RetryInfo, breakerInfo *BreakerInfo) {
	ci := &clientInfo{
		circuitBreaker: NewCB(breakerInfo),
		retryHandler:   retryInfo,
		client:         c,
	}
	m.clients[id] = ci
}
