package cli

import (
	"context"
	"io"
	"os"
	"path/filepath"
)

type App struct {
	// application name
	Name string
	// application usage information
	Usage string
	//
	HelpName    string
	ArgsUsage   string
	UsageText   string
	Version     string
	HideVersion bool
	// the function to be invoked on the default execution
	Action ActionFunc
	// global flags
	Flags []*FlagBase
	// application commands
	Commands        []*Command
	Writer          io.Writer
	HideHelp        bool
	HideHelpCommand bool
	CommandVisible  bool
	setupComplete   bool
	rootCommand     *Command
}

func (app *App) initialize() {
	if app.setupComplete {
		return
	}

	app.setupComplete = true

	if app.Name == "" {
		app.Name = filepath.Base(os.Args[0])
	}

	if app.HelpName == "" {
		app.HelpName = app.Name
	}

	if app.Usage == "" {
		app.Usage = "CLI App 101"
	}

	if app.Version == "" {
		app.HideVersion = true
	}

	var newCommands []*Command
	for _, c := range app.Commands {
		newCommands = append(newCommands, c)
	}
	app.Commands = newCommands

	if app.Command(helpCommand.Name) == nil && !app.HideHelp {
		if !app.HideHelpCommand {
			app.appendCommand(helpCommand)
		}
		if HelpFlag != nil {
			app.appendFlag(HelpFlag)
		}
	}

	if len(app.Commands) > 0 {
		app.CommandVisible = true
	}

	if app.Action == nil {
		app.Action = helpCommand.Action
	}

	if app.Writer == nil {
		app.Writer = os.Stdout
	}
}

func (app *App) Execute(arguments []string) error {
	return app.ExecuteContext(context.Background(), arguments)
}

func (app *App) ExecuteContext(ctx context.Context, arguments []string) error {
	app.initialize()

	// args []

	// create a map of the command [key->command_name, value->{ action: command_action, subcommands: map of commands, flags: [flag_name, flagParser] }]
	// mandatory flags

	conTxt := NewContext(app, &Context{Context: ctx})

	app.rootCommand = app.newRootCommand()
	conTxt.Command = app.rootCommand

	return app.rootCommand.Run(conTxt, arguments...)
}

func (app *App) newRootCommand() *Command {
	return &Command{
		Name:      app.Name,
		Usage:     app.Usage,
		Action:    app.Action,
		Flags:     app.Flags,
		Commands:  app.Commands,
		ArgsUsage: app.ArgsUsage,
	}
}

func (app *App) writer() io.Writer {
	return app.Writer
}

func (app *App) Command(name string) *Command {
	for _, c := range app.Commands {
		if c.HasName(name) {
			return c
		}
	}
	return nil
}

func (app *App) appendCommand(c *Command) {
	if !hasCommand(app.Commands, c) {
		app.Commands = append(app.Commands, c)
	}
}

func (app *App) appendFlag(flag *FlagBase) {
	if !hasFlag(app.Flags, flag) {
		app.Flags = append(app.Flags, flag)
	}
}

func (app *App) VisibleCommands() []*Command {
	return app.Commands
}
