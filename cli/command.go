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
	// command specific flags
	Flags []*FlagBase
	// subcommands of the root command
	Commands []*Command
	HelpName string
}

func (command *Command) Run(conTxt *Context, arguments ...string) error {
	a := args(arguments)
	inputArgs := a.FetchArgs()
	command.addUserDefinedFlags()

	isHelpPresent := a.checkForHelp()
	var finalCommand *Command

	parseArgs()

	if len(inputArgs) > 0 {
		finalCommand = command.findCommandPath(conTxt, inputArgs)
		if finalCommand == nil {
			return errors.New("command not found")
		}
		command.Action = finalCommand.Action
	}

	if isHelpPresent {
		conTxt.Command = finalCommand
		return helpCommand.Action(conTxt)
	}

	//if checkHelpFlag(conTxt, inputArgs) {
	//	return helpCommand.Action(conTxt)
	//}

	if command.Action == nil {
		command.Action = helpCommand.Action
	}

	err := command.Action(conTxt)
	return err
}

// with default flag library they can be parsed if they are added before args
func parseArgs() {
	flag.Parse()
	flag.Visit(func(f *flag.Flag) {
		if f.Value != nil {
			mappedFlags[f.Name] = f.Value
		} else {
			mappedFlags[f.Name] = f.DefValue
		}
	})
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

func (command *Command) findCommandPath(conTxt *Context, args []string) *Command {
	for _, c := range command.Commands {
		found := search(c, conTxt, args)
		if found != nil {
			return found
		}
	}
	return nil
}

func search(command *Command, conTxt *Context, args []string) *Command {
	if command == nil {
		return nil
	}
	if command.Name == args[0] {
		if len(args) == 1 {
			return command
		}
		for _, child := range command.Commands {
			if search(child, conTxt, args[1:]) != nil {
				return child
			}
		}
	}
	return nil
}

func (command *Command) checkForAlias(arg string) bool {
	for _, a := range command.Aliases {
		if arg == a {
			return true
		}
	}
	return false
}

func (command *Command) addUserDefinedFlags() {
	setFlags(command.Name, command.Flags)
}
