package cli

import (
	"errors"
	"flag"
)

type Command struct {
	// command name used to invoke from CLI
	Name string
	// command usage information
	Usage     string
	ArgsUsage string
	// the array of aliases to invoke the commands
	Aliases []string
	// execute on the command invocation
	Action ActionFunc
	Flags  []*FlagBase
	// subcommands of the root command
	Commands []*Command
}

func (command *Command) Run(conTxt *Context, arguments ...string) error {
	a := args(arguments)
	inputArgs := a.FetchArgs()
	if len(inputArgs) > 2 {
		return errors.New("multiple args not supported")
	}
	set, err := command.parseArgs()
	if err != nil {
		return err
	}
	conTxt.flagsSet = set

	if isHelp(conTxt, inputArgs) {
		return helpCommand.Action(conTxt)
	}

	if command.Action == nil {
		command.Action = helpCommand.Action
	}

	if len(inputArgs) > 0 {
		currentCommand, err := command.GetCommand(inputArgs[0])
		if err != nil {
			logger.Error(err)
			command.Action = helpCommand.Action
		} else {
			command.Action = currentCommand.Action
		}
	}

	err = command.Action(conTxt)
	return err
}

func (command *Command) parseArgs() (*flag.FlagSet, error) {
	set, err := command.newFlagSet()
	if err != nil {
		return nil, err
	}
	return set, nil
}

func (command *Command) newFlagSet() (*flag.FlagSet, error) {
	return flagSet(command.Name, command.allFlags())
}

func (command *Command) allFlags() []*FlagBase {
	var flags []*FlagBase
	flags = append(flags, command.Flags...)
	return flags
}

func (command *Command) HasName(name string) bool {
	for _, n := range command.Names() {
		if n == name {
			return true
		}
	}
	return false
}

func (command *Command) Names() []string {
	return append([]string{command.Name}, command.Aliases...)
}

func hasCommand(commands []*Command, command *Command) bool {
	for _, exist := range commands {
		if command == exist {
			return true
		}
	}
	return false
}

func (command *Command) GetCommand(arg string) (*Command, error) {
	for _, c := range command.Commands {
		// TODO : logic can be improved
		if c.Name == arg || c.checkForAlias(arg) {
			return c, nil
		}
	}
	return nil, errors.New("no command present")
}

func (command *Command) checkForAlias(arg string) bool {
	for _, a := range command.Aliases {
		if arg == a {
			return true
		}
	}
	return false
}
