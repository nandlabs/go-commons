package errutils

import (
	"errors"
	"fmt"
	"testing"
)

func TestFmtError(t *testing.T) {
	t.Run("Testing execution of error utils", func(t *testing.T) {
		want := errors.New("testing error utils")
		got := FmtError("testing error utils")
		if got.Error() != want.Error() {
			t.Errorf("invalid error string generated")
		}
	})

	t.Run("Testing execution of error utils with formatting", func(t *testing.T) {
		want := errors.New(fmt.Sprintf("expecting errors %d", 0))
		got := FmtError("expecting errors %d", 0)
		if got.Error() != want.Error() {
			t.Errorf("invalid error string generated")
		}
	})
}
