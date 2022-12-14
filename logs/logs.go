package logs

import "sync"

var l Logger
var LogLevel = INFO
var LogPrefix = ""
var LogCalldepth = 2
var LogFlag = 0
var mutex sync.Mutex

func SetLogLevel(level Level) {
	LogLevel = level
}

func SetLogPrefix(prefix string) {
	LogPrefix = prefix
}

func SetLogCalldepth(calldepth int) {
	LogCalldepth = calldepth
}

func SetLogFlag(flag int) {
	LogFlag = flag
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
	if l == nil {
		mutex.Lock()
		defer mutex.Unlock()
		if l == nil {
			l = New(WithLogLevel(LogLevel), WithLogPrefix(LogPrefix), WithLogCalldepth(LogCalldepth))
		}
	}
	return l
}

func Trace(args ...interface{}) {
	GetLogger().Trace(args...)
}

func Tracef(format string, args ...interface{}) {
	GetLogger().Tracef(format, args...)
}

func Debug(args ...interface{}) {
	GetLogger().Debug(args...)
}

func Debugf(format string, args ...interface{}) {
	GetLogger().Debugf(format, args...)
}

func Info(args ...interface{}) {
	GetLogger().Info(args...)
}

func Infof(format string, args ...interface{}) {
	GetLogger().Infof(format, args...)
}

func Warn(args ...interface{}) {
	GetLogger().Warn(args...)
}

func Warnf(format string, args ...interface{}) {
	GetLogger().Warnf(format, args...)
}

func Error(args ...interface{}) {
	GetLogger().Error(args...)
}

func Errorf(format string, args ...interface{}) {
	GetLogger().Errorf(format, args...)
}

func Fatal(args ...interface{}) {
	GetLogger().Fatal(args...)
}

func Fatalf(format string, args ...interface{}) {
	GetLogger().Fatalf(format, args...)
}

func Panic(args ...interface{}) {
	GetLogger().Panic(args...)
}

func Panicf(format string, args ...interface{}) {
	GetLogger().Panicf(format, args...)
}
