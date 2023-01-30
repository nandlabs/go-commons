package cli

import (
	"flag"
)

type Command struct {
	Name      string
	Usage     string
	ArgsUsage string
	Aliases   []string
	Action    ActionFunc
	Flags     []Flag
	Commands  []*Command
}

func (command *Command) Run(conTxt *Context, arguments ...string) error {
	a := args(arguments)
	set, err := command.parseFlags(&a, conTxt)
	conTxt.flagsSet = set

	//if isHelp(a[1]) {
	//	return helpCommand.Action(conTxt)
	//}

	if command.Action == nil {
		command.Action = helpCommand.Action
	}

	//err = command.Action(conTxt)
	err = helpCommand.Action(conTxt)
	return err
}

func (command *Command) parseFlags(args Args, conTxt *Context) (*flag.FlagSet, error) {
	set, err := command.newFlagSet()
	if err != nil {
		return nil, err
	}
	return set, nil
}

func (command *Command) newFlagSet() (*flag.FlagSet, error) {
	return flagSet(command.Name, command.allFlags())
}

func (command *Command) allFlags() []Flag {
	var flags []Flag
	flags = append(flags, command.Flags...)
	return flags
}

func (command *Command) HasName(name string) bool {
	return false
}

func (command *Command) Names() []string {
	return append([]string{command.Name}, command.Aliases...)
}
