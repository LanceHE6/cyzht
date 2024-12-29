package logger

import (
	"errors"
	"testing"
)

func TestLogger(t *testing.T) {
	Logger.Info("This is an info message")
	Logger.Infof("This is an info message with %s", "args")
	Logger.Error("This is an error message")
	Logger.Errorf("This is an error message with %s", "args")
	err := errors.New("this is a mock error")
	Logger.ErrorWithErr(err)
	Logger.Debug("This is a debug message")
	Logger.Debugf("This is a debug message with %s", "args")
	Logger.Warn("This is a warning message")
	Logger.Warnf("This is a warning message with %s", "args")
}
