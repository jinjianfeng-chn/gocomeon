package logs

import (
	"context"
	"testing"
)

func TestLog(t *testing.T) {
	Trace(context.Background(), "hello world")
	Debug(context.Background(), "hello world")
	Infof(context.Background(), "hello world, %s", "1")
	Infof(context.Background(), "hello world")
	Infof(context.Background(), "hello world")
	Warn(context.Background(), "hello world")
	Error(context.Background(), "hello world")
	Fatal(context.Background(), "hello world")
}
