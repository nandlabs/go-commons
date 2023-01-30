package cli

import (
	"flag"
	"fmt"
)

type Flag interface {
	fmt.Stringer

	Apply(*flag.FlagSet) error

	Names() []string

	IsSet() bool
}

func flagSet(name string, flags []Flag) (*flag.FlagSet, error) {
	set := flag.NewFlagSet(name, flag.ContinueOnError)
	for _, f := range flags {
		if err := f.Apply(set); err != nil {
			return nil, err
		}
	}
	return set, nil
}
