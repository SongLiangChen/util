package turn

import (
	"errors"
	"sync"
	"time"
)

var (
	ErrTurnWaitTimeOut = errors.New("turn: wait time out")
)

// Turn a concurrent controller
type Turn struct {
	timers sync.Pool
	queue  chan struct{}
}

// NewTurn returns a Turn, the param 'cap' is maximum number of concurrency allowed
func NewTurn(cap int) *Turn {
	return &Turn{
		queue: make(chan struct{}, cap),
		timers: sync.Pool{
			New: func() interface{} {
				t := time.NewTimer(time.Hour)
				t.Stop()
				return t
			},
		},
	}
}

// Get blocking to get a concurrent permission
func (t *Turn) Get() {
	t.queue <- struct{}{}
}

// Free release a concurrent permission
func (t *Turn) Free() {
	<-t.queue
}

// Wait try to get a concurrent permission in 'timeout' period.
// it returns ErrTurnWaitTimeOut if NOT success in 'timeout' period
func (t *Turn) Wait(timeout time.Duration) error {
	select {
	case t.queue <- struct{}{}:
		return nil

	default:
		timer := t.timers.Get().(*time.Timer)
		timer.Reset(timeout)

		select {
		case t.queue <- struct{}{}:
			if !timer.Stop() {
				<-timer.C
			}

			t.timers.Put(timer)
			return nil

		case <-timer.C:
			t.timers.Put(timer)
			return ErrTurnWaitTimeOut
		}
	}
}
