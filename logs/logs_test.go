package logs

import "testing"

func TestLog(t *testing.T) {
	Trace("hello world")
	Debug("hello world")
	Infof("hello world, %s", "1")
	Infof("hello world")
	Infof("hello world")
	Warn("hello world")
	Error("hello world")
	Fatal("hello world")
}
