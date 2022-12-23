package secrets

import (
	"bytes"
	"go.nandlabs.io/commons/config"
	"time"
)

type Credential struct {
	Value       *bytes.Buffer
	LastUpdated time.Time
	Version     string
	MetaData    config.Properties
}
