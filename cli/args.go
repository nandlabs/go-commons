package cli

import (
	"strings"
)

type Args interface {
	Get(n int) string
	First() string
	FetchArgs() *ArgsCli
}

type args []string

type ArgsCli struct {
	inputCommands []string
	inputFlags    []string
}

func (a *args) Get(n int) string {
	if len(*a) > n {
		return (*a)[n]
	}
	return ""
}

func (a *args) First() string {
	return a.Get(0)
}

func (a *args) FetchArgs() *ArgsCli {
	var outputCommands []string
	var outputFlags []string
	var tail []string
	if len(*a) >= 2 {
		tail = (*a)[1:]
	}
	for _, item := range tail {
		if isFlag(item) {
			trimmedItem := strings.TrimPrefix(strings.TrimPrefix(item, "-"), "--")
			outputFlags = append(outputFlags, trimmedItem)
		} else {
			outputCommands = append(outputCommands, item)
		}
	}
	return &ArgsCli{
		inputCommands: outputCommands,
		inputFlags:    outputFlags,
	}
}

func isFlag(item string) bool {
	if strings.HasPrefix(item, "-") || strings.HasPrefix(item, "--") {
		return true
	}
	return false
}

func (a *args) checkForHelp() (isPresent bool) {
	programArgs := (*a)[1:]
	if len(programArgs) > 0 {
		lastItem := programArgs[len(programArgs)-1]
		if lastItem == "-help" || lastItem == "-h" {
			isPresent = true
		}
	}

	return
}
