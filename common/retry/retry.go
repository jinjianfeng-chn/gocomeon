package retry

import (
	"fmt"
)

type LogOutput interface {
	Debugln(args ...interface{})
}

type ConsoleLogOutput struct {
}

func (p *ConsoleLogOutput) Debugln(args ...interface{}) {
	fmt.Println(args)
}

type Retryable interface {
	// Required Retry handler interface decides whether a retry is required for the given error.
	Required(int, error) bool
	// DoActionBeforeRetry Do something before retry again. The first time is not executed.
	DoActionBeforeRetry(int, error)
	// DoAction Do action
	DoAction() (interface{}, error)
	// GetLogOutput Get log output
	GetLogOutput() LogOutput
}

// Invoke invokes the given function and performs retries according to the retry options.
func Invoke(retryable Retryable) (interface{}, error) {
	logOutput := retryable.GetLogOutput()
	if logOutput == nil {
		logOutput = &ConsoleLogOutput{}
	}

	attempts := 1
	for {
		result, e := retryable.DoAction()
		if e == nil {
			if attempts > 1 {
				logOutput.Debugln(fmt.Sprintf("success on attempt #%d", attempts))
			}
			return result, nil
		}

		logOutput.Debugln(fmt.Sprintf("failed with error [%s] on attempt #%d", e, attempts))
		if !retryable.Required(attempts, e) {
			if attempts > 1 {
				logOutput.Debugln(fmt.Sprintf("retry for error [%s] is not warranted after %d attempt(s)", e, attempts))
			}
			return result, e
		}

		logOutput.Debugln(fmt.Sprintf("retry for error [%s] is warranted after %d attempt(s)", e, attempts))
		retryable.DoActionBeforeRetry(attempts, e)
		attempts++
	}
}
