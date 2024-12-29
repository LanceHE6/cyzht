package logger

import (
	gologger "github.com/phachon/go-logger"
)

// Logger 使用go-logger封装的日志打印
var Logger = &logger{
	Logger: gologger.NewLogger(),
}

type logger struct {
	Logger *gologger.Logger
}

func (l *logger) Info(msg string) {
	l.Logger.Info(msg)
}
func (l *logger) Infof(format string, args ...interface{}) {
	l.Logger.Infof(format, args...)
}

func (l *logger) Error(errMsg string) {
	l.Logger.Error(errMsg)
}
func (l *logger) Errorf(format string, args ...interface{}) {
	l.Logger.Errorf(format, args...)
}
func (l *logger) ErrorWithErr(err error) {
	l.Logger.Error(err.Error())
}

func (l *logger) Debug(msg string) {
	l.Logger.Debug(msg)
}
func (l *logger) Debugf(format string, args ...interface{}) {
	l.Logger.Debugf(format, args...)
}

func (l *logger) Warn(msg string) {
	l.Logger.Warning(msg)
}
func (l *logger) Warnf(format string, args ...interface{}) {
	l.Logger.Warningf(format, args...)
}
