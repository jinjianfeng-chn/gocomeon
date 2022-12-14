package logs

import (
	"testing"
)

func TestLogger(t *testing.T) {
	l := New()
	l.Trace("hello world")
	l.Debug("hello world")
	l.Infof("hello world, %s", "1")
	l.Infof("hello world")
	l.Infof("hello world")
	l.Warn("hello world")
	l.Error("hello world")
	l.Fatal("hello world")
}
