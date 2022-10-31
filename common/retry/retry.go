package retry

import (
	"fmt"
	"time"
)

type Retryable[T any] interface {
	// Required Retry handler interface decides whether a retry is required for the given error.
	Required(int, error) bool
	// DoActionBeforeRetry Do something before retry again. The first time is not executed.
	DoActionBeforeRetry(int, error)
	// DoAction Do action
	DoAction() (T, error)
	// RetryInterval retry interval time, param is the number of last retries
	RetryInterval(int) time.Duration
	// Logout output log
	Logout(log string)
}

// Invoke invokes the given function and performs retries according to the retry options.
func Invoke[T any](retryable Retryable[T]) (T, error) {
	attempts := 0
	var e error
	var result T

	for {
		attempts++
		if attempts > 1 {
			retryable.DoActionBeforeRetry(attempts, e)
		}
		result, e = retryable.DoAction()
		if e == nil {
			if attempts > 1 {
				retryable.Logout(fmt.Sprintf("success on attempt #%d", attempts))
			}
			return result, nil
		}
		retryable.Logout(fmt.Sprintf("failed with error [%s] on attempt #%d", e, attempts))
		if !retryable.Required(attempts, e) {
			retryable.Logout(fmt.Sprintf("retry for error [%s] is not warranted after %d attempt(s)", e, attempts))
			return result, e
		}
		interval := retryable.RetryInterval(attempts)
		retryable.Logout(fmt.Sprintf("retry for error [%s] is warranted after %d attempt(s). the retry will begin after %s", e, attempts, interval))
		time.Sleep(interval)
	}
}
