## cli

This project is a Go library for building powerful and user-friendly command-line interfaces (CLIs). The library makes it easy to create and manage complex command structures, parse command-line arguments, and provide helpful error messages and usage information to the user.

---
- [Installation](#installation)
- [Features](#features)
- [Usage](#usage)
---

### Installation

```bash
go get go.nandlabs.io/commons/cli
```

### Features

* Easy to use API for building complex command structures 
* Argument parsing and validation 
* Automatically generates usage and help information 
* Written in Go and follows best practices for Go programming

### Usage

```go
package main

import (
	"fmt"
	"os"

	cli "go.nandlabs.io/commons/cli"
)

func main() {
	app := cli.NewApp("my-cli")
	app.Usage = "A simple CLI for demonstration purposes"

	app.Commands = []cli.Command{
		{
			Name:    "greet",
			Aliases: []string{"g"},
			Usage:   "Greet the user",
			Action: func(c *cli.Context) error {
				fmt.Println("Hello, world!")
				return nil
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Println("Error running CLI:", err)
		os.Exit(1)
	}
}
```