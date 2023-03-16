package retry

import (
	"errors"
	"fmt"
	"time"
)

// Action Do action.
type Action[T any] func() (T, error)

type Retryable interface {
	// Required Retry handler interface decides whether a retry is required for the given error.
	Required(attempt int, e error) bool
	// DoActionBeforeRetry Do something before retry again. The first time is not executed.
	DoActionBeforeRetry(attempt int, e error)
	// RetryInterval retry interval time, param is the number of last retries
	RetryInterval(attempt int) time.Duration
	// Logout output log
	Logout(log string)
}

type RetryableCustom struct {
	RequiredFunc            func(attempt int, e error) bool
	DoActionBeforeRetryFunc func(attempt int, e error)
	RetryIntervalFunc       func(attempt int) time.Duration
	LogoutFunc              func(log string)
}

func (p *RetryableCustom) Required(attempt int, e error) bool {
	if p.RequiredFunc == nil {
		return false
	}
	if errors.Is(e, &ErrorActionIsNil{}) {
		return false
	}
	return p.RequiredFunc(attempt, e)
}

func (p *RetryableCustom) DoActionBeforeRetry(attempt int, e error) {
	if p.DoActionBeforeRetryFunc == nil {
		return
	}
	p.DoActionBeforeRetryFunc(attempt, e)
}

func (p *RetryableCustom) RetryInterval(attempt int) time.Duration {
	if p.RetryIntervalFunc == nil {
		return 100 * time.Millisecond
	}
	return p.RetryIntervalFunc(attempt)
}

func (p *RetryableCustom) Logout(log string) {
	if p.LogoutFunc == nil {
		return
	}
	p.LogoutFunc(log)
}

// Retry do retry the given function and performs retries according to the retry options.
func Retry[T any](retryable Retryable, action Action[T]) (result T, e error) {
	if action == nil {
		e = &ErrorActionIsNil{}
		return
	}
	if retryable == nil {
		result, e = action()
		return
	}
	attempts := 0
	for {
		attempts++
		if attempts > 1 {
			retryable.DoActionBeforeRetry(attempts, e)
		}
		result, e = action()
		if e == nil {
			if attempts > 1 {
				retryable.Logout(fmt.Sprintf("success on attempt #%d", attempts))
			}
			return
		}
		retryable.Logout(fmt.Sprintf("failed with error [%s] on attempt #%d", e, attempts))
		if !retryable.Required(attempts, e) {
			retryable.Logout(fmt.Sprintf("retry for error [%s] is not warranted after %d attempt(s)", e, attempts))
			return
		}
		interval := retryable.RetryInterval(attempts)
		retryable.Logout(fmt.Sprintf("retry for error [%s] is warranted after %d attempt(s). the retry will begin after %s", e, attempts, interval))
		time.Sleep(interval)
	}
}
