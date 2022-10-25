package retry

import (
	"errors"
	"time"
)

type ErrorActionIsNil struct {
}

func (p *ErrorActionIsNil) Error() string {
	return "Action can not be nil"
}

// Action Do action
type Action func() (interface{}, error)

// ActionBeforeRetry Do something before retry again.
type ActionBeforeRetry func(int, error)

// RetryableInterval Retry interval function
type RetryableInterval func(int) time.Duration

type RetryableTimes struct {
	// Attempts the number retry attempts
	Attempts int
	// RetryableInterval retry time interval, unit is millisecond
	RetryableInterval RetryableInterval
	// Action do action
	Action Action
	// ActionBeforeRetry do something before retry again
	ActionBeforeRetry ActionBeforeRetry
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
	return p.RetryableInterval(attempt)
}

func (p *RetryableTimes) DoActionBeforeRetry(attempt int, e error) {
	if p.ActionBeforeRetry != nil {
		p.ActionBeforeRetry(attempt, e)
	}
}

func (p *RetryableTimes) DoAction() (interface{}, error) {
	if p.Action == nil {
		return nil, &ErrorActionIsNil{}
	}
	return p.Action()
}

func (p *RetryableTimes) GetLogOutput() LogOutput {
	return p.LogOutput
}
