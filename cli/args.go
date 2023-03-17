package cli

import (
	"strings"
)

type Args interface {
	Get(n int) string
	First() string
	FetchArgs() []string
}

type args []string

func (a *args) Get(n int) string {
	if len(*a) > n {
		return (*a)[n]
	}
	return ""
}

func (a *args) First() string {
	return a.Get(0)
}

func (a *args) FetchArgs() []string {
	var outputArgs []string
	var tail []string
	if len(*a) >= 2 {
		tail = (*a)[1:]
	}
	for _, item := range tail {
		if !isFlag(item) {
			outputArgs = append(outputArgs, item)
		}
	}
	return outputArgs
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
