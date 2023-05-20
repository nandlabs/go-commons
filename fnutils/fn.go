package fnutils

import (
	"time"
)

// ExecuteAfterSecs executes the given function after the specified timeout duration,
// expressed in seconds.
//
// It converts the timeout value from seconds to a duration in seconds and then calls
// the ExecuteAfter function, passing the converted duration and the provided function.
//
// Example:
//
//	ExecuteAfterSecs(func() {
//	    fmt.Println("Hello, World!")
//	}, 10)
//
// @param fn The function to be executed.
//
// @param timeout The timeout duration in seconds.
func ExecuteAfterSecs(fn func(), timeout int) {
	ExecuteAfter(fn, time.Second*time.Duration(timeout))
}

// ExecuteAfterMs executes the given function after the specified timeout duration,
// expressed in milliseconds.
//
// It converts the timeout value from milliseconds to a duration in seconds and then calls
// the ExecuteAfter function, passing the converted duration and the provided function.
//
// Example:
//
//	ExecuteAfterMs(func() {
//	    fmt.Println("Hello, World!")
//	}, 1000)
//
// @param fn The function to be executed.
//
// @param timeout The timeout duration in milliseconds.
func ExecuteAfterMs(fn func(), timeout int64) {
	ExecuteAfter(fn, time.Millisecond*time.Duration(timeout))
}

// ExecuteAfterMin executes the given function after the specified timeout duration,
// expressed in minutes.
//
// It converts the timeout value from minutes to a duration in seconds and then calls
// the ExecuteAfter function, passing the converted duration and the provided function.
//
// Example:
//
//	ExecuteAfterMin(func() {
//	    fmt.Println("Hello, World!")
//	}, 5)
//
// @param fn The function to be executed.
//
// @param timeout The timeout duration in minutes.
func ExecuteAfterMin(fn func(), timeout int) {
	ExecuteAfter(fn, time.Minute*time.Duration(timeout))
}

// ExecuteAfter executes the given function after the specified timeout duration.
//
// It waits for the specified duration and then calls the provided function.
//
// Example:
//
//	ExecuteAfter(func() {
//	    fmt.Println("Hello, World!")
//	}, time.Second)
//
// @param fn The function to be executed.
//
// @param timeout The duration to wait before executing the function.
func ExecuteAfter(fn func(), timeout time.Duration) {
	select {
	case <-time.After(timeout):
		{
			fn()
		}
	}
}
