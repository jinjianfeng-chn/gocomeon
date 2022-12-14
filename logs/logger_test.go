package logs

import (
	"testing"
)

func TestLogger(t *testing.T) {
	logs.Trace("hello world")
	logs.Debug("hello world")
	logs.Infof("hello world")
	logs.Infof("hello world")
	logs.Infof("hello world")
	logs.Warn("hello world")
	logs.Error("hello world")
	logs.Fatal("hello world")
}
