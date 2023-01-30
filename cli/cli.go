package cli

import (
	"context"
	"io"
	"os"
	"path/filepath"
)

type App struct {
	Name      string
	Usage     string
	HelpName  string
	ArgsUsage string
	UsageText string
	Action    ActionFunc
	Flags     []Flag
	Commands  []*Command
	Writer    io.Writer

	setupComplete bool
	rootCommand   *Command
}

func (app *App) init() {
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
	app.init()

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
