package errutils

import (
	"errors"
	"fmt"
)

// FmtError formats an error message using a format string and optional values,
// and returns it as an error object.
//
// It takes a format string f and variadic arguments v, and uses them to construct
// the error message by calling fmt.Sprintf function. The resulting error message
// is then wrapped into an error object using errors.New.
//
// Example:
//
//	err := FmtError("Invalid input: %s", userInput)
//	if err != nil {
//		log.Println(err)
//	}
//
// @param f The format string specifying the error message.
// @param v Optional values to be inserted into the format string.
//
// @returns An error object containing the formatted error message.
func FmtError(f string, v ...any) error {
	return errors.New(fmt.Sprintf(f, v...))
}
