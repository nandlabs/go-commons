package cli

import "go.nandlabs.io/commons/l3"

var (
	logger = l3.Get()
)

type ActionFunc func(conTxt *Context) error
