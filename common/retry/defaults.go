package retry

import "time"

// Action Do action
type Action func() (interface{}, error)

// ActionBeforeRetry Do something before retry again.
type ActionBeforeRetry func(int, error)

// Logger Print log
type Logger func(log string)

type RetryableTimes struct {
	// Attempts the number retry attempts
	Attempts int
	// Interval retry time interval, unit is millisecond
	Interval time.Duration
	// Action do action
	Action Action
	// ActionBeforeRetry do something before retry again
	ActionBeforeRetry ActionBeforeRetry
	// Logger print log
	Logger Logger
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

func (p *RetryableTimes) Log(log string) {
	p.Logger(log)
}
