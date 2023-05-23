package logs

import (
	"context"
	"testing"
)

func TestLogger(t *testing.T) {
	l := NewLogger()
	l.Trace(context.Background(), "hello world")
	l.Debug(context.Background(), "hello world")
	l.Infof(context.Background(), "hello world, %s", "1")
	l.Infof(context.Background(), "hello world")
	l.Infof(context.Background(), "hello world")
	l.Warn(context.Background(), "hello world")
	l.Error(context.Background(), "hello world")
	l.Fatal(context.Background(), "hello world")
}
