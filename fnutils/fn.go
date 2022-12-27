package fnutils

import (
	"time"
)

func ExecuteAfterSecs(fn func(), timeout int) {
	ExecuteAfter(fn, time.Second*time.Duration(timeout))
}

func ExecuteAfterMs(fn func(), timeout int64) {
	ExecuteAfter(fn, time.Millisecond*time.Duration(timeout))
}
func ExecuteAfterMin(fn func(), timeout int) {
	ExecuteAfter(fn, time.Minute*time.Duration(timeout))
}

func ExecuteAfter(fn func(), timeout time.Duration) {
	select {
	case <-time.After(timeout):
		{
			fn()
		}
	}
}
