package secrets

import (
	"time"

	"go.nandlabs.io/commons/config"
)

type Credential struct {
	Value       []byte
	LastUpdated time.Time
	Version     string
	MetaData    config.Properties
}

func (c *Credential) Str() (s string) {

	if c.Value != nil {
		s = string(c.Value)
	}

	return
}
