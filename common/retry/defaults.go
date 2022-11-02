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

// RetryableTimes Retry until the maximum number of retry times is reached
type RetryableTimes struct {
	// Attempts the number retry attempts
	Attempts int
	// Interval retry time interval, unit is millisecond
	Interval int64
	// ActionBeforeRetry do something before retry again
	ActionBeforeRetry func(int, error)
	// LogOutput print log
	LogOutput LogOutput
}

func (p *RetryableTimes) Required(attempt int, e error) bool {
	if errors.Is(e, &ErrorActionIsNil{}) {
		return false
	}
	return attempt < p.Attempts
}

func (p *RetryableTimes) RetryInterval(attempt int) time.Duration {
	return time.Duration(p.Interval * int64(time.Millisecond))
}

func (p *RetryableTimes) DoActionBeforeRetry(attempt int, e error) {
	if p.ActionBeforeRetry != nil {
		p.ActionBeforeRetry(attempt, e)
	}
}

func (p *RetryableTimes) Logout(log string) {
	if p.LogOutput != nil {
		p.LogOutput.Debugln(log)
	}
}

// RetryableTimesBackoff Retry until the maximum number of retry times is reached.
// The backoff algorithm is used for the retry interval.
type RetryableTimesBackoff struct {
	RetryableTimes
	backoff.Backoff
}

func (p *RetryableTimesBackoff) RetryInterval(attempt int) time.Duration {
	if p.InitialBackoff == 0 {
		p.InitialBackoff = time.Duration(p.Interval * int64(time.Millisecond))
	}
	return p.Next()
}
