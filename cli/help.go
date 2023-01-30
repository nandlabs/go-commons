package cli

import (
	"fmt"
	"io"
	"strings"
	"text/tabwriter"
	"text/template"
)

var (
	HelpFlags = [2]string{"help", "h"}
)

var PrintHelp helpPrinter = printHelp

var PrintCustomHelp helpCustomPrinter = printCustomHelp

type helpPrinter func(w io.Writer, template string, data interface{})

type helpCustomPrinter func(w io.Writer, template string, data interface{}, customFunc map[string]interface{})

var helpCommand = &Command{
	Name:      "help",
	Aliases:   []string{"h"},
	Usage:     "Shows a list of commands or help for one command",
	ArgsUsage: "[command]",
	Action: func(conTxt *Context) error {
		fmt.Println("invoking help action")
		args := conTxt.Args()
		argsPresent := args.First() != ""
		firstArg := args.First()

		if conTxt.Command.Name == "help" || conTxt.Command.Name == "h" {
			conTxt = conTxt.parentContext
		}

		if argsPresent {
			return ShowCommandHelp(conTxt, firstArg)
		}

		if conTxt.parentContext.App == nil {
			_ = ShowAppHelp(conTxt)
			return nil
		}

		return nil
	},
}

func isHelp(arg string) bool {
	found := false
	for _, name := range HelpFlags {
		if arg == name {
			found = true
		}
	}
	return found
}

func ShowCommandHelp(conTxt *Context, command string) error {
	commands := conTxt.App.Commands
	if conTxt.Command.Commands != nil {
		commands = conTxt.Command.Commands
	}
	for _, c := range commands {
		if c.HasName(command) {
			template := CommandHelpTemplate

			PrintHelp(conTxt.App.writer(), template, c)

			return nil
		}
	}

	// add check for the command not found

	//conTxt.App.CommandNotFound(conTxt, command)

	return nil
}

func ShowAppHelp(conTxt *Context) error {
	tpl := AppHelpTemplate
	printHelp(conTxt.App.writer(), tpl, conTxt.App)
	return nil
}

func printHelp(out io.Writer, template string, data interface{}) {
	PrintCustomHelp(out, template, data, nil)
}

func printCustomHelp(out io.Writer, templ string, data interface{}, customFuncs map[string]interface{}) {
	const maxLineLength = 1000

	funcMap := template.FuncMap{
		"join":           strings.Join,
		"subtract":       subtract,
		"indent":         indent,
		"nindent":        nindent,
		"trim":           strings.TrimSpace,
		"wrap":           func(input string, offset int) string { return wrap(input, offset, maxLineLength) },
		"offset":         offset,
		"offsetCommands": offsetCommands,
	}
	w := tabwriter.NewWriter(out, 1, 8, 2, ' ', 0)
	t := template.Must(template.New("help").Funcs(funcMap).Parse(templ))
	t.New("helpNameTemplate").Parse(helpNameTemplate)
	t.New("usageTemplate").Parse(usageTemplate)
	t.New("descriptionTemplate").Parse(descriptionTemplate)

	err := t.Execute(w, data)
	if err != nil {
		fmt.Println(err)
	}
	_ = w.Flush()
}

func subtract(a, b int) int {
	return a - b
}

func indent(spaces int, v string) string {
	pad := strings.Repeat(" ", spaces)
	return pad + strings.Replace(v, "\n", "\n"+pad, -1)
}

func nindent(spaces int, v string) string {
	return "\n" + indent(spaces, v)
}

func wrap(input string, offset int, wrapAt int) string {
	var ss []string

	lines := strings.Split(input, "\n")
	padding := strings.Repeat(" ", offset)

	for i, line := range lines {
		if line == "" {
			ss = append(ss, line)
		} else {
			wrapped := wrapLine(line, offset, wrapAt, padding)
			if i == 0 {
				ss = append(ss, wrapped)
			} else {
				ss = append(ss, padding+wrapped)
			}
		}
	}
	return strings.Join(ss, "\n")
}
func wrapLine(input string, offset int, wrapAt int, padding string) string {
	if wrapAt <= offset || len(input) <= wrapAt-offset {
		return input
	}

	lineWidth := wrapAt - offset
	words := strings.Fields(input)
	if len(words) == 0 {
		return input
	}

	wrapped := words[0]
	spaceLeft := lineWidth - len(wrapped)
	for _, word := range words[1:] {
		if len(word)+1 > spaceLeft {
			wrapped += "\n" + padding + word
			spaceLeft = lineWidth - len(word)
		} else {
			wrapped += " " + word
			spaceLeft -= 1 + len(word)
		}
	}
	return wrapped
}

func offset(input string, fixed int) int {
	return len(input) + fixed
}

func offsetCommands(cmds []*Command, fixed int) int {
	var max int = 0
	for _, cmd := range cmds {
		s := strings.Join(cmd.Names(), ", ")
		if len(s) > max {
			max = len(s)
		}
	}
	return max + fixed
}
