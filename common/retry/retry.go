package retry

import (
	"fmt"
)

type Retryable interface {
	// Required Retry handler interface decides whether a retry is required for the given error.
	Required(int, error) bool

	// DoActionBeforeRetry Do something before retry again. The first time is not executed.
	DoActionBeforeRetry(int, error)
	// DoAction Do action
	DoAction() (interface{}, error)
	// Log Print log
	Log(log string)
}

// Invoke invokes the given function and performs retries according to the retry options.
func Invoke(retryable Retryable) (interface{}, error) {
	attempts := 1
	for {
		result, e := retryable.DoAction()
		if e == nil {
			if attempts > 1 {
				retryable.Log(fmt.Sprintf("success on attempt #%d", attempts))
			}
			return result, nil
		}

		retryable.Log(fmt.Sprintf("failed with error [%s] on attempt #%d", e, attempts))
		if !retryable.Required(attempts, e) {
			if attempts > 1 {
				retryable.Log(fmt.Sprintf("retry for error [%s] is not warranted after %d attempt(s)", e, attempts))
			}
			return result, e
		}

		retryable.Log(fmt.Sprintf("retry for error [%s] is warranted after %d attempt(s)", e, attempts))
		retryable.DoActionBeforeRetry(attempts, e)
		attempts++
	}
}
