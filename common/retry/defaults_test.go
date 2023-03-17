package retry

import (
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/jinjianfeng-chn/gocomeon/common/backoff"
	"github.com/jinjianfeng-chn/gocomeon/logs"
)

func TestRetryableTimes(t *testing.T) {
	logs.SetLogLevel(logs.DEBUG)

	retryableTimes := &RetryableTimes{
		Attempts: 3,
		Interval: 1000,
	}

	action := func() (string, error) {
		return "", errors.New("do action error")
	}

	t.Run("normal", func(t *testing.T) {
		r, e := Retry[string](retryableTimes, action)
		fmt.Println(r)

		msg := "do action error"
		if e.Error() != msg {
			t.Fatal(fmt.Sprintf("error message expected [%s], but [%s] got", msg, e.Error()))
		}
	})

	t.Run("action is nil", func(t *testing.T) {
		_, e := Retry[string](retryableTimes, nil)
		if !errors.Is(e, &ErrorActionIsNil{}) {
			t.Fatal(fmt.Sprintf("error message expected [%s], but [%s] got", &ErrorActionIsNil{}, e.Error()))
		}
	})
}

func TestRetryableTimesBackoff(t *testing.T) {
	logs.SetLogLevel(logs.DEBUG)

	action := func() (string, error) {
		return "", errors.New("do action error")
	}

	retryableTimes := &RetryableTimesBackoff{
		RetryableTimes: RetryableTimes{
			Attempts:  10,
			LogOutput: logs.GetLogger(),
		},
		Backoff: backoff.Backoff{
			InitialBackoff: time.Second,
			MaxBackoff:     10 * time.Second,
			BackoffFactor:  2,
		},
	}

	t.Run("normal", func(t *testing.T) {
		r, e := Retry[string](retryableTimes, action)
		fmt.Println(r)

		msg := "do action error"
		if e.Error() != msg {
			t.Fatal(fmt.Sprintf("error message expected [%s], but [%s] got", msg, e.Error()))
		}
	})

	t.Run("action is nil", func(t *testing.T) {
		_, e := Retry[string](retryableTimes, nil)
		if !errors.Is(e, &ErrorActionIsNil{}) {
			t.Fatal(fmt.Sprintf("error message expected [%s], but [%s] got", &ErrorActionIsNil{}, e.Error()))
		}
	})
}
