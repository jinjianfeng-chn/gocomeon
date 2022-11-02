package retry

import (
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/jinjianfeng-chn/gocomeon/common/backoff"
)

func TestRetryableTimes(t *testing.T) {

	retryableTimes := RetryableTimes{
		Attempts: 3,
		Interval: 1000,
	}

	action := func() (string, error) {
		return "", errors.New("do action error")
	}

	retryableInvoke := &RetryableInvoke[string]{
		Retryable: &retryableTimes,
	}

	t.Run("normal", func(t *testing.T) {
		r, e := retryableInvoke.Invoke(action)
		fmt.Println(r)

		msg := "do action error"
		if e.Error() != msg {
			t.Fatal(fmt.Sprintf("error message expected [%s], but [%s] got", msg, e.Error()))
		}
	})

	t.Run("action is nil", func(t *testing.T) {
		_, e := retryableInvoke.Invoke(nil)
		if !errors.Is(e, &ErrorActionIsNil{}) {
			t.Fatal(fmt.Sprintf("error message expected [%s], but [%s] got", &ErrorActionIsNil{}, e.Error()))
		}
	})
}

func TestRetryableTimesBackoff(t *testing.T) {

	action := func() (string, error) {
		return "", errors.New("do action error")
	}

	retryableTimes := RetryableTimesBackoff{
		RetryableTimes: RetryableTimes{
			Attempts: 3,
		},
		Backoff: backoff.Backoff{
			InitialBackoff: time.Second,
			MaxBackoff:     10 * time.Second,
			BackoffFactor:  2,
		},
	}

	retryableInvoke := &RetryableInvoke[string]{
		Retryable: &retryableTimes,
	}

	t.Run("normal", func(t *testing.T) {
		r, e := retryableInvoke.Invoke(action)
		fmt.Println(r)

		msg := "do action error"
		if e.Error() != msg {
			t.Fatal(fmt.Sprintf("error message expected [%s], but [%s] got", msg, e.Error()))
		}
	})

	t.Run("action is nil", func(t *testing.T) {
		_, e := retryableInvoke.Invoke(nil)
		if !errors.Is(e, &ErrorActionIsNil{}) {
			t.Fatal(fmt.Sprintf("error message expected [%s], but [%s] got", &ErrorActionIsNil{}, e.Error()))
		}
	})
}
