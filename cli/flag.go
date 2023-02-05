package cli

import (
	"flag"
)

// a flag will always be rpefixed with -- (name) or - (alias)
type FlagBase struct {
	Name    string
	Usage   string
	Aliases []string
	// default value of the flag
	Default interface{}
	Var     any // pointer
}

var HelpFlag = &FlagBase{
	Name:    "help",
	Usage:   "show help",
	Aliases: []string{"h"},
}

//type Flag interface {
//	fmt.Stringer
//
//	Apply(*flag.FlagSet) error
//
//	Names() []string
//
//	IsSet() bool
//}

func flagSet(name string, flags []*FlagBase) (*flag.FlagSet, error) {
	set := flag.NewFlagSet(name, flag.ContinueOnError)
	for _, f := range flags {
		if err := f.Apply(set); err != nil {
			return nil, err
		}
	}
	return set, nil
}

func hasFlag(flags []*FlagBase, flag *FlagBase) bool {
	for _, exist := range flags {
		if flag == exist {
			return true
		}
	}
	return false
}

func (flag *FlagBase) Apply(flagSet *flag.FlagSet) error {
	flagSet.String(flag.Name, "", flag.Usage)
	return nil
}
