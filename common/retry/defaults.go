package retry

import (
	"errors"
	"time"

	"github.com/jinjianfeng-chn/gocomeon/common/backoff"
)

type ErrorActionIsNil struct {
}

func (p *ErrorActionIsNil) Error() string {
	return "Action can not be nil"
}

type LogOutput interface {
	Debugln(args ...any)
}

// Action Do action.
type Action[T any] func() (T, error)

// ActionBeforeRetry Do something before retry again.
type ActionBeforeRetry func(int, error)

// RetryableTimes Retry until the maximum number of retry times is reached
type RetryableTimes[T any] struct {
	// Attempts the number retry attempts
	Attempts int
	// Interval retry time interval, unit is millisecond
	Interval int64
	// Action do action
	Action Action[T]
	// ActionBeforeRetry do something before retry again
	ActionBeforeRetry ActionBeforeRetry
	// LogOutput print log
	LogOutput LogOutput
}

func (p *RetryableTimes[T]) Required(attempt int, e error) bool {
	if errors.Is(e, &ErrorActionIsNil{}) {
		return false
	}
	return attempt < p.Attempts
}

func (p *RetryableTimes[T]) RetryInterval(attempt int) time.Duration {
	return time.Duration(p.Interval * int64(time.Millisecond))
}

func (p *RetryableTimes[T]) DoActionBeforeRetry(attempt int, e error) {
	if p.ActionBeforeRetry != nil {
		p.ActionBeforeRetry(attempt, e)
	}
}

func (p *RetryableTimes[T]) DoAction() (T, error) {
	if p.Action == nil {
		var result T
		return result, &ErrorActionIsNil{}
	}
	return p.Action()
}

func (p *RetryableTimes[T]) Logout(log string) {
	if p.LogOutput != nil {
		p.LogOutput.Debugln(log)
	}
}

// RetryableTimesBackoff Retry until the maximum number of retry times is reached.
// The backoff algorithm is used for the retry interval.
type RetryableTimesBackoff[T any] struct {
	RetryableTimes[T]
	backoff.Backoff
}

func (p *RetryableTimesBackoff[T]) RetryInterval(attempt int) time.Duration {
	if p.InitialBackoff == 0 {
		p.InitialBackoff = time.Duration(p.Interval * int64(time.Millisecond))
	}
	return p.Next()
}
