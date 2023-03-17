package retry

import (
	"context"
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
	ctx, cancel := context.WithCancel(context.Background())
	second := 3
	go func() {
		time.Sleep(time.Duration(second) * time.Second)
		cancel()
	}()
	result, e := Retry(ctx, &TestRetryable{}, func() (string, error) {
		attempts++
		return "", errors.New(fmt.Sprintf("do action error %d", attempts))
	})
	fmt.Println(fmt.Sprintf("result=%v, error=%v", result, e))
	msg := fmt.Sprintf("do action error %d", second)
	if e.Error() != msg {
		t.Fatalf("error message expected [%s], but [%s] got", msg, e.Error())
	}
}
