package main

import (
	"fmt"
	"go.nandlabs.io/commons/cli"
	"log"
	"os"
)

const (
	ProjectDir  = "pd"
	ProfileFile = "pf"
)

func main() {
	app := &cli.App{
		Version: "v0.0.1",
		Action: func(ctx *cli.Context) error {
			fmt.Printf("Hello %q", ctx.Args().Get(0))
			fmt.Println(ctx.GetFlag(ProjectDir))
			fmt.Println(ctx.GetFlag(ProfileFile))
			return nil
		},
		Commands: []*cli.Command{
			{
				Name:    "test",
				Usage:   "this is a test command",
				Aliases: []string{"t"},
				Action: func(ctx *cli.Context) error {
					fmt.Println("hello from test command")
					fmt.Println(ctx.GetFlag(ProjectDir))
					fmt.Println(ctx.GetFlag(ProfileFile))
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
				Commands: []*cli.Command{
					{
						Name:  "slow",
						Usage: "run slow",
						Action: func(ctx *cli.Context) error {
							fmt.Println("time to run slow")
							return nil
						},
					},
					{
						Name:  "fast",
						Usage: "run fast",
						Action: func(ctx *cli.Context) error {
							fmt.Println("time to run fast")
							return nil
						},
					},
				},
			},
		},
		Flags: []*cli.FlagBase{
			{
				Name:    ProjectDir,
				Aliases: []string{"pd"},
				Default: "",
				Usage:   "Directory of the project to be built",
			},
			{
				Name:    ProfileFile,
				Aliases: []string{"pf"},
				Default: "",
				Usage:   "Profile file name to be used",
			},
		},
	}

	if err := app.Execute(os.Args); err != nil {
		log.Fatal(err)
	}
}
