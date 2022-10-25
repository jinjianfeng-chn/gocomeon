package retry

import "time"

// Action Do action
type Action func() (interface{}, error)

// ActionBeforeRetry Do something before retry again.
type ActionBeforeRetry func(int, error)

type RetryableTimes struct {
	// Attempts the number retry attempts
	Attempts int
	// Interval retry time interval, unit is millisecond
	Interval time.Duration
	// Action do action
	Action Action
	// ActionBeforeRetry do something before retry again
	ActionBeforeRetry ActionBeforeRetry
	// LogOutput print log
	LogOutput LogOutput
}

func (p *RetryableTimes) Required(attempt int, e error) bool {
	return attempt < p.Attempts
}

func (p *RetryableTimes) DoActionBeforeRetry(attempt int, e error) {
	time.Sleep(p.Interval * time.Microsecond)
	p.ActionBeforeRetry(attempt, e)
}

func (p *RetryableTimes) DoAction() (interface{}, error) {
	return p.Action()
}

func (p *RetryableTimes) GetLogOutput() LogOutput {
	return p.LogOutput
}
