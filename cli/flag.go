package cli

import (
	"flag"
	"strings"
)

var (
	mappedFlags = make(map[string]interface{})
	flagMap     = make(map[string]*Flag)
)

// Flag a flag will always be prefixed with -- (name) or - (alias)
type Flag struct {
	Name    string
	Usage   string
	Aliases []string
	// default value of the flag
	Default interface{}
	Value   interface{}
	//Var     any // pointer
}

// HelpFlag built-in flag present in the system
var HelpFlag = &Flag{
	Name:    "help",
	Usage:   "show help",
	Aliases: []string{"-h", "--help"},
	Default: "",
}

func hasFlag(flags []*Flag, flag *Flag) bool {
	for _, exist := range flags {
		if flag == exist {
			return true
		}
	}
	return false
}

// improve based on the type of flags
func setFlags(commandFlags []*Flag, inputFlags []string) {
	parsedFlags := parseFlags(commandFlags, inputFlags)
	for _, f := range parsedFlags {
		if f.Name == "help" {
			f.AddHelpFlag()
		} else {
			f.AddFlagToSet()
		}
	}
}

func (f *Flag) AddFlagToSet() {
	flag.String(f.Name, f.Value.(string), f.Usage)
}

func (f *Flag) AddHelpFlag() {
	flag.Bool(f.Name, true, f.Usage)
}

func parseFlags(commandFlags []*Flag, inputFlags []string) []*Flag {
	createFlagMap(commandFlags)
	var result []*Flag
	for _, item := range inputFlags {
		itemArr := strings.Split(item, "=")
		if len(itemArr) > 1 {
			key := itemArr[0]
			val := itemArr[1]
			mappedFlag := flagMap[key]
			result = append(result, &Flag{
				Name:    mappedFlag.Name,
				Usage:   mappedFlag.Usage,
				Aliases: nil,
				Default: mappedFlag.Default,
				Value:   val,
			})
		}
	}
	return result
}

func createFlagMap(commandFlags []*Flag) {
	for _, item := range commandFlags {
		for _, alias := range item.Aliases {
			flagMap[alias] = &Flag{
				Name:    item.Name,
				Usage:   item.Usage,
				Aliases: nil,
				Default: item.Default,
				Value:   nil,
			}
		}
	}
}
