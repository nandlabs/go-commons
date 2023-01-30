package main

import (
	"fmt"
	"go.nandlabs.io/commons/cli"
	"log"
	"os"
)

func main() {
	app := &cli.App{
		Name:  "dummy",
		Usage: "dummy app for cli",
		Action: func(ctx *cli.Context) error {
			fmt.Println("dummy world")
			return nil
		},
	}

	if err := app.Execute(os.Args); err != nil {
		log.Fatal(err)
	}
}
