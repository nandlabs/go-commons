package secrets

import (
	"context"
	"sync"
)

type Store interface {
	Get(key string, ctx context.Context) (*Credential, error)
	Write(key string, credential *Credential, ctx context.Context) error
	Provider() string
}

type Manager struct {
	stores map[string]Store
	once   sync.Once
}

func (m *Manager) Register(store Store) {
	if m.stores == nil {
		m.once.Do(func() {
			m.stores = make(map[string]Store)
		})
	}
	m.stores[store.Provider()] = store
}

func (m *Manager) Store(name string) (store Store) {
	if m.stores != nil {
		store = m.stores[name]
	}
	return
}
