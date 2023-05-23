package logs

import (
	"context"
	"sync"
)

var l Logger
var LogLevel = INFO
var LogPrefix = "[GO]"
var LogCalldepth = 3
var mutex sync.Mutex

func SetLogLevel(level Level) {
	LogLevel = level
	if l != nil {
		ResetLogger()
	}
}

func SetLogPrefix(prefix string) {
	LogPrefix = prefix
	if l != nil {
		ResetLogger()
	}
}

func SetLogCalldepth(calldepth int) {
	LogCalldepth = calldepth
	if l != nil {
		ResetLogger()
	}
}

func SetLogger(logger Logger) {
	mutex.Lock()
	defer mutex.Unlock()
	l = logger
}

func ResetLogger() {
	mutex.Lock()
	defer mutex.Unlock()
	l = nil
}

func GetLogger() Logger {
	if l != nil {
		return l
	}
	mutex.Lock()
	defer mutex.Unlock()
	if l == nil {
		l = NewLogger(
			WithLogLevel(LogLevel),
			WithLogFormatter(&logFormatter{LogPrefix, LogCalldepth}),
		)
	}
	return l
}

func Trace(ctx context.Context, args ...interface{}) {
	GetLogger().Trace(ctx, args...)
}

func Tracef(ctx context.Context, format string, args ...interface{}) {
	GetLogger().Tracef(ctx, format, args...)
}

func Debug(ctx context.Context, args ...interface{}) {
	GetLogger().Debug(ctx, args...)
}

func Debugf(ctx context.Context, format string, args ...interface{}) {
	GetLogger().Debugf(ctx, format, args...)
}

func Info(ctx context.Context, args ...interface{}) {
	GetLogger().Info(ctx, args...)
}

func Infof(ctx context.Context, format string, args ...interface{}) {
	GetLogger().Infof(ctx, format, args...)
}

func Warn(ctx context.Context, args ...interface{}) {
	GetLogger().Warn(ctx, args...)
}

func Warnf(ctx context.Context, format string, args ...interface{}) {
	GetLogger().Warnf(ctx, format, args...)
}

func Error(ctx context.Context, args ...interface{}) {
	GetLogger().Error(ctx, args...)
}

func Errorf(ctx context.Context, format string, args ...interface{}) {
	GetLogger().Errorf(ctx, format, args...)
}

func Fatal(ctx context.Context, args ...interface{}) {
	GetLogger().Fatal(ctx, args...)
}

func Fatalf(ctx context.Context, format string, args ...interface{}) {
	GetLogger().Fatalf(ctx, format, args...)
}

func Panic(ctx context.Context, args ...interface{}) {
	GetLogger().Panic(ctx, args...)
}

func Panicf(ctx context.Context, format string, args ...interface{}) {
	GetLogger().Panicf(ctx, format, args...)
}
