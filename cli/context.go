package cli

import (
	"context"
	"flag"
)

type Context struct {
	context.Context
	App           *App
	Command       *Command
	flagsSet      *flag.FlagSet
	parentContext *Context
}

func NewContext(app *App, parentCtx *Context) *Context {
	c := &Context{
		App:           app,
		parentContext: parentCtx,
	}
	if parentCtx != nil {
		c.Context = parentCtx.Context
	}
	c.Command = &Command{}
	if c.Context == nil {
		c.Context = context.Background()
	}
	return c
}

func (conTxt *Context) Args() Args {
	res := args(conTxt.flagsSet.Args())
	return &res
}

func (conTxt *Context) Bool(name string) bool {
	if v, ok := conTxt.Value(name).(bool); ok {
		return v
	}
	return false
}

func (conTxt *Context) Value(name string) interface{} {
	if fs := conTxt.lookupFlagSet(name); fs != nil {
		return fs.Lookup(name).Value.(flag.Getter).Get()
	}
	return nil
}

func (conTxt *Context) lookupFlagSet(name string) *flag.FlagSet {
	if f := conTxt.flagsSet.Lookup(name); f != nil {
		return conTxt.flagsSet
	}
	return nil
}
