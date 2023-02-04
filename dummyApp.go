package main

import (
	"fmt"
	"go.nandlabs.io/commons/cli"
	"log"
	"os"
)

func init() {
}

func main() {
	app := &cli.App{
		Version: "v0.0.1",
		Commands: []*cli.Command{
			{
				Name:    "test",
				Usage:   "this is a test command",
				Aliases: []string{"t"},
				Action: func(ctx *cli.Context) error {
					fmt.Println("hello from test command")
					return nil
				},
			},
			{
				Name:    "run",
				Usage:   "time to run",
				Aliases: []string{"r"},
				Action: func(ctx *cli.Context) error {
					fmt.Println("time to run away")
					return nil
				},
			},
		},
	}

	if err := app.Execute(os.Args); err != nil {
		log.Fatal(err)
	}
}
