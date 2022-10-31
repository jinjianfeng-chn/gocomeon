package retry

import (
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/jinjianfeng-chn/gocomeon/common/backoff"
)

func TestRetryableTimes(t *testing.T) {

	retryableTimes := RetryableTimes[string]{
		Attempts: 3,
		Interval: 1000,
		Action: func() (string, error) {
			return "", errors.New("do action error")
		},
	}

	t.Run("normal", func(t *testing.T) {
		r, e := Invoke[string](&retryableTimes)
		fmt.Println(r)

		msg := "do action error"
		if e.Error() != msg {
			t.Fatal(fmt.Sprintf("error message expected [%s], but [%s] got", msg, e.Error()))
		}
	})

	t.Run("action is nil", func(t *testing.T) {
		action := retryableTimes.Action
		retryableTimes.Action = nil
		_, e := Invoke[string](&retryableTimes)

		if !errors.Is(e, &ErrorActionIsNil{}) {
			t.Fatal(fmt.Sprintf("error message expected [%s], but [%s] got", &ErrorActionIsNil{}, e.Error()))
		}
		retryableTimes.Action = action
	})
}

func TestRetryableTimesBackoff(t *testing.T) {

	retryableTimes := RetryableTimesBackoff[string]{
		RetryableTimes: RetryableTimes[string]{
			Attempts: 3,
			Action: func() (string, error) {
				return "", errors.New("do action error")
			},
		},
		Backoff: backoff.Backoff{
			InitialBackoff: time.Second,
			MaxBackoff:     10 * time.Second,
			BackoffFactor:  2,
		},
	}

	t.Run("normal", func(t *testing.T) {
		r, e := Invoke[string](&retryableTimes)
		fmt.Println(r)

		msg := "do action error"
		if e.Error() != msg {
			t.Fatal(fmt.Sprintf("error message expected [%s], but [%s] got", msg, e.Error()))
		}
	})

	t.Run("action is nil", func(t *testing.T) {
		action := retryableTimes.Action
		retryableTimes.Action = nil
		_, e := Invoke[string](&retryableTimes)

		if !errors.Is(e, &ErrorActionIsNil{}) {
			t.Fatal(fmt.Sprintf("error message expected [%s], but [%s] got", &ErrorActionIsNil{}, e.Error()))
		}
		retryableTimes.Action = action
	})
}
