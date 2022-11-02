package retry

import (
	"errors"
	"fmt"
	"testing"
	"time"
)

type TestRetryable struct {
	attempts int
}

func (p *TestRetryable) Required(attempt int, e error) bool {
	return attempt < 5
}

func (p *TestRetryable) DoActionBeforeRetry(attempt int, e error) {
	fmt.Println(fmt.Sprintf("attempt=%d, error=%v", attempt, e))
}

func (p *TestRetryable) RetryInterval(attempt int) time.Duration {
	return time.Second
}

func (p *TestRetryable) Logout(log string) {

}

func TestInvoke(t *testing.T) {
	var attempts int
	result, e := Retry(&TestRetryable{}, func() (string, error) {
		attempts++
		return "", errors.New(fmt.Sprintf("do action error %d", attempts))
	})
	fmt.Println(fmt.Sprintf("result=%v, error=%v", result, e))
	msg := "do action error 5"
	if e.Error() != msg {
		t.Fatalf("error message expected [%s], but [%s] got", msg, e.Error())
	}
}
