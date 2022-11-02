package retry

import (
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

type RetryableInvoke[T any] struct {
	Retryable Retryable
}

// Invoke invokes the given function and performs retries according to the retry options.
func (p *RetryableInvoke[T]) Invoke(action Action[T]) (result T, e error) {
	if action == nil {
		e = &ErrorActionIsNil{}
		return
	}
	attempts := 0
	for {
		attempts++
		if attempts > 1 {
			p.Retryable.DoActionBeforeRetry(attempts, e)
		}
		result, e = action()
		if e == nil {
			if attempts > 1 {
				p.Retryable.Logout(fmt.Sprintf("success on attempt #%d", attempts))
			}
			return
		}
		p.Retryable.Logout(fmt.Sprintf("failed with error [%s] on attempt #%d", e, attempts))
		if !p.Retryable.Required(attempts, e) {
			p.Retryable.Logout(fmt.Sprintf("retry for error [%s] is not warranted after %d attempt(s)", e, attempts))
			return
		}
		interval := p.Retryable.RetryInterval(attempts)
		p.Retryable.Logout(fmt.Sprintf("retry for error [%s] is warranted after %d attempt(s). the retry will begin after %s", e, attempts, interval))
		time.Sleep(interval)
	}
}

func Retry[T any](retryable Retryable, action Action[T]) (result T, e error) {
	return (&RetryableInvoke[T]{retryable}).Invoke(action)
}
