package fnutils

import (
	"testing"
	"time"
)

func TestExecuteAfter(t *testing.T) {
	t.Run("Execution occurs after the specified timeout", func(t *testing.T) {
		executed := false
		fn := func() {
			executed = true
		}
		timeout := 5 * time.Second

		ExecuteAfter(fn, timeout)

		if !executed {
			t.Error("Expected execution to occur after the specified timeout")
		}
	})

	t.Run("Execution occurs immediately when timeout is zero", func(t *testing.T) {
		executed := false
		fn := func() {
			executed = true
		}
		timeout := 0 * time.Second

		ExecuteAfter(fn, timeout)

		if !executed {
			t.Error("Expected execution to occur immediately")
		}
	})

	t.Run("Execution occurs only once", func(t *testing.T) {
		counter := 0
		fn := func() {
			counter++
		}
		timeout := 1 * time.Second

		ExecuteAfter(fn, timeout)
		time.Sleep(2 * time.Second)

		if counter != 1 {
			t.Errorf("Expected execution to occur once, got %d times", counter)
		}
	})
	//
	//t.Run("Multiple executions occur if the function takes longer than the timeout", func(t *testing.T) {
	//	executionCount := 0
	//	fn := func() {
	//		executionCount++
	//		time.Sleep(2 * time.Second)
	//	}
	//	timeout := 1 * time.Second
	//
	//	go ExecuteAfter(fn, timeout)
	//	time.Sleep(3 * time.Second)
	//
	//	if executionCount != 2 {
	//		t.Errorf("Expected execution to occur twice, got %d times", executionCount)
	//	}
	//})
	//
	//t.Run("No execution occurs if the function is nil", func(t *testing.T) {
	//	_ = func() {
	//		// This function should not be executed
	//		t.Error("Function should not be executed")
	//	}
	//	timeout := 1 * time.Second
	//
	//	ExecuteAfter(nil, timeout)
	//})
	//
	//t.Run("No execution occurs if the timeout is set to 0 and the function is nil", func(t *testing.T) {
	//	_ = func() {
	//		// This function should not be executed
	//		t.Error("Function should not be executed")
	//	}
	//	timeout := 0 * time.Second
	//
	//	ExecuteAfter(nil, timeout)
	//})
	//
	//t.Run("Execution does not occur when timeout is negative", func(t *testing.T) {
	//	executed := false
	//	fn := func() {
	//		executed = true
	//	}
	//	timeout := -1 * time.Second
	//
	//	ExecuteAfter(fn, timeout)
	//
	//	if executed {
	//		t.Error("Expected execution to not occur with negative timeout")
	//	}
	//})
}

func TestExecuteAfterSecs(t *testing.T) {
	t.Run("Execution after specific timeout (seconds)", func(t *testing.T) {
		executed := false
		fn := func() {
			executed = true
		}
		ExecuteAfterSecs(fn, 5)
		if !executed {
			t.Error("Expected execution to occur after the specified timeout")
		}
	})
}

func TestExecuteAfterMs(t *testing.T) {
	t.Run("Execution after specific timeout (milliseconds)", func(t *testing.T) {
		executed := false
		fn := func() {
			executed = true
		}
		ExecuteAfterMs(fn, 50)
		if !executed {
			t.Error("Expected execution to occur after the specified timeout")
		}
	})
}

func TestExecuteAfterMin(t *testing.T) {
	t.Run("Execution after specific timeout (minutes)", func(t *testing.T) {
		executed := false
		fn := func() {
			executed = true
		}
		ExecuteAfterMin(fn, 1)
		if !executed {
			t.Error("Expected execution to occur after the specified timeout")
		}
	})
}
