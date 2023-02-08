package cli

import (
	"flag"
	"fmt"
)

var (
	mappedFlags = make(map[string]interface{})
)

// FlagBase a flag will always be prefixed with -- (name) or - (alias)
type FlagBase struct {
	Name    string
	Usage   string
	Aliases []string
	// default value of the flag
	Default interface{}
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
func setFlags(name string, inputFlags []*FlagBase) {
	//set := flag.NewFlagSet(name, flag.ContinueOnError)
	fmt.Println(inputFlags)
	for _, f := range inputFlags {
		fmt.Println(f.Name)
		if f.Name == "help" {
			f.AddHelpFlag()
		} else {
			f.AddFlagToSet()
		}
	}
}

func (f *FlagBase) AddFlagToSet() {
	flag.String(f.Name, f.Default.(string), f.Usage)
}

func (f *FlagBase) AddHelpFlag() {
	flag.Bool(f.Name, true, f.Usage)
}
