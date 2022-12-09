package errutils

import (
	"errors"
	"fmt"
)

func FmtError(f string, v ...any) error {
	return errors.New(fmt.Sprintf(f, v...))
}
