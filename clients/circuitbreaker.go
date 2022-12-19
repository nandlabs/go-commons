package clients

import (
	"errors"
	"sync/atomic"
	"time"
)

const (
	circuitClosed uint32 = iota
	circuitHalfOpen
	circuitOpen
	defaultTimeout          = 300
	defaultMaxHalfOpen      = 5
	defaultSuccessThreshold = 3
	defaultFailureThreshold = 3
)

var CBOpenErr = errors.New("the Circuit breaker is open and unable to process request")

type BreakerInfo struct {
	FailureThreshold uint64
	SuccessThreshold uint64
	MaxHalfOpen      uint32
	Timeout          uint32
}

// CircuitBreaker struct
type CircuitBreaker struct {
	*BreakerInfo
	currentState    uint32
	successCounter  uint64
	failureCounter  uint64
	halfOpenCounter uint32
}

func NewCB(info *BreakerInfo) (cb *CircuitBreaker) {

	if info == nil {
		info = &BreakerInfo{}
	}

	if info.SuccessThreshold == 0 {
		info.SuccessThreshold = defaultSuccessThreshold
	}
	if info.FailureThreshold == 0 {
		info.FailureThreshold = defaultFailureThreshold
	}
	if info.MaxHalfOpen == 0 {
		info.MaxHalfOpen = defaultMaxHalfOpen
	}

	if info.Timeout == 0 {
		info.Timeout = defaultTimeout
	}

	return &CircuitBreaker{
		BreakerInfo:     info,
		successCounter:  0,
		failureCounter:  0,
		halfOpenCounter: 0,
		currentState:    circuitClosed,
	}

}

func (cb *CircuitBreaker) CanExecute() (err error) {
	state := cb.getState()
	if state == circuitOpen {
		err = CBOpenErr
	} else if state == circuitHalfOpen {
		val := atomic.AddUint32(&cb.halfOpenCounter, 1)
		if val > cb.MaxHalfOpen {
			cb.updateState(circuitHalfOpen, circuitOpen)
			err = CBOpenErr
		}
	}
	return
}

func (cb *CircuitBreaker) OnExecution(success bool) {
	var val uint64
	state := cb.getState()
	if success {
		val = atomic.AddUint64(&cb.successCounter, 1)
		if state == circuitHalfOpen {
			if val >= cb.SuccessThreshold {
				cb.updateState(circuitHalfOpen, circuitClosed)
			}
		}
	} else {
		val = atomic.AddUint64(&cb.failureCounter, 1)
		// Check if the failure threshold is reached
		if state == circuitClosed {
			if val >= cb.FailureThreshold {
				cb.updateState(circuitClosed, circuitOpen)
			}
		}
	}
}

func (cb *CircuitBreaker) Reset() {
	atomic.StoreUint32(&cb.currentState, circuitClosed)
	atomic.StoreUint64(&cb.failureCounter, 0)
	atomic.StoreUint64(&cb.successCounter, 0)
	atomic.StoreUint32(&cb.halfOpenCounter, 0)

}

func (cb *CircuitBreaker) updateState(oldState, newState uint32) {

	if atomic.CompareAndSwapUint32(&cb.currentState, oldState, newState) {
		atomic.StoreUint64(&cb.successCounter, 0)
		atomic.StoreUint64(&cb.failureCounter, 0)
		atomic.StoreUint32(&cb.halfOpenCounter, 0)
		//Check if moving to circuitOpen state
		if newState == circuitOpen {
			//Start Timer for HalfOpen
			go func() {
				select {
				case <-time.After(time.Second * time.Duration(cb.Timeout)):
					{
						cb.updateState(circuitOpen, circuitHalfOpen)
					}
				}
			}()
		}
	}

}

func (cb *CircuitBreaker) getState() (s uint32) {
	atomic.LoadUint32(&s)
	return
}
