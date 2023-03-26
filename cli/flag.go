package cli

import (
	"flag"
	"strings"
)

var (
	mappedFlags = make(map[string]interface{})
	flagMap     = make(map[string]*FlagBase)
)

// FlagBase a flag will always be prefixed with -- (name) or - (alias)
type FlagBase struct {
	Name    string
	Usage   string
	Aliases []string
	// default value of the flag
	Default interface{}
	Value   interface{}
	//Var     any // pointer
}

// HelpFlag built-in flag present in the system
var HelpFlag = &FlagBase{
	Name:    "help",
	Usage:   "show help",
	Aliases: []string{"-h", "--help"},
	Default: "",
}

func hasFlag(flags []*FlagBase, flag *FlagBase) bool {
	for _, exist := range flags {
		if flag == exist {
			return true
		}
	}
	return false
}

// improve based on the type of flags
func setFlags(commandFlags []*FlagBase, inputFlags []string) {
	parsedFlags := parseFlags(commandFlags, inputFlags)
	for _, f := range parsedFlags {
		if f.Name == "help" {
			f.AddHelpFlag()
		} else {
			f.AddFlagToSet()
		}
	}
}

func (f *FlagBase) AddFlagToSet() {
	flag.String(f.Name, f.Value.(string), f.Usage)
}

func (f *FlagBase) AddHelpFlag() {
	flag.Bool(f.Name, true, f.Usage)
}

func parseFlags(commandFlags []*FlagBase, inputFlags []string) []*FlagBase {
	createFlagMap(commandFlags)
	var result []*FlagBase
	for _, item := range inputFlags {
		itemArr := strings.Split(item, "=")
		if len(itemArr) > 1 {
			key := itemArr[0]
			val := itemArr[1]
			mappedFlag := flagMap[key]
			result = append(result, &FlagBase{
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

func createFlagMap(commandFlags []*FlagBase) {
	for _, item := range commandFlags {
		for _, alias := range item.Aliases {
			flagMap[alias] = &FlagBase{
				Name:    item.Name,
				Usage:   item.Usage,
				Aliases: nil,
				Default: item.Default,
				Value:   nil,
			}
		}
	}
}
