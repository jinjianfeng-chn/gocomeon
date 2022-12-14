package logs

import (
	"fmt"
	"log"
	"os"
)

type Level int

// 基本从低到高
const (
	ALL Level = iota + 1
	TRACE
	DEBUG
	INFO
	WARN
	ERROR
	FATAL
	OFF
)

type Logger interface {
	Trace(args ...interface{})

	Tracef(format string, args ...interface{})

	Debug(args ...interface{})

	Debugf(format string, args ...interface{})

	Info(args ...interface{})

	Infof(format string, args ...interface{})

	Warn(args ...interface{})

	Warnf(format string, args ...interface{})

	Error(args ...interface{})

	Errorf(format string, args ...interface{})

	Fatal(args ...interface{})

	Fatalf(format string, args ...interface{})

	Panic(args ...interface{})

	Panicf(format string, args ...interface{})
}

var logs = NewLogger(INFO, "")

func New(options ...LoggerOption) Logger {
	l := &logger{}
	for _, option := range options {
		option(l)
	}
	if l.level == 0 {
		l.level = INFO
	}
	if l.calldepth == 0 {
		l.calldepth = 2
	}
	l.stdout = log.New(os.Stdout, l.prefix, log.Ldate|log.Ltime|log.Lmicroseconds|log.Llongfile)
	l.stderr = log.New(os.Stderr, l.prefix, log.Ldate|log.Ltime|log.Lmicroseconds|log.Llongfile)
	return l
}

func NewLogger(level Level, prefix string) Logger {
	return &logger{
		level:     level,
		prefix:    prefix,
		calldepth: 2,
		stdout:    log.New(os.Stdout, prefix, log.Ldate|log.Ltime|log.Lmicroseconds|log.Llongfile),
		stderr:    log.New(os.Stderr, prefix, log.Ldate|log.Ltime|log.Lmicroseconds|log.Llongfile),
	}
}

type logger struct {
	level     Level
	prefix    string
	stdout    *log.Logger
	stderr    *log.Logger
	calldepth int
}

func (p *logger) Trace(args ...interface{}) {
	if TRACE < p.level {
		return
	}
	_ = p.stdout.Output(p.calldepth, fmt.Sprintln(args...))
}

func (p *logger) Tracef(format string, args ...interface{}) {
	if TRACE < p.level {
		return
	}
	_ = p.stdout.Output(p.calldepth, fmt.Sprintf(format, args...))
}

func (p *logger) Debug(args ...interface{}) {
	if DEBUG < p.level {
		return
	}
	_ = p.stdout.Output(p.calldepth, fmt.Sprintln(args...))
}

func (p *logger) Debugf(format string, args ...interface{}) {
	if DEBUG < p.level {
		return
	}
	_ = p.stdout.Output(p.calldepth, fmt.Sprintf(format, args...))
}

func (p *logger) Info(args ...interface{}) {
	if INFO < p.level {
		return
	}
	_ = p.stdout.Output(p.calldepth, fmt.Sprintln(args...))
}

func (p *logger) Infof(format string, args ...interface{}) {
	if INFO < p.level {
		return
	}
	_ = p.stdout.Output(p.calldepth, fmt.Sprintf(format, args...))
}

func (p *logger) Warn(args ...interface{}) {
	if WARN < p.level {
		return
	}
	_ = p.stdout.Output(p.calldepth, fmt.Sprintln(args...))
}

func (p *logger) Warnf(format string, args ...interface{}) {
	if WARN < p.level {
		return
	}
	_ = p.stdout.Output(p.calldepth, fmt.Sprintf(format, args...))
}

func (p *logger) Error(args ...interface{}) {
	if ERROR < p.level {
		return
	}
	_ = p.stderr.Output(p.calldepth, fmt.Sprintln(args...))
}

func (p *logger) Errorf(format string, args ...interface{}) {
	if ERROR < p.level {
		return
	}
	_ = p.stderr.Output(p.calldepth, fmt.Sprintf(format, args...))
}

func (p *logger) Fatal(args ...interface{}) {
	if FATAL < p.level {
		return
	}
	_ = p.stderr.Output(p.calldepth, fmt.Sprintln(args...))
	os.Exit(1)
}

func (p *logger) Fatalf(format string, args ...interface{}) {
	if FATAL < p.level {
		return
	}
	_ = p.stderr.Output(p.calldepth, fmt.Sprintf(format, args...))
	os.Exit(1)
}

func (p *logger) Panic(args ...interface{}) {
	if FATAL < p.level {
		return
	}
	s := fmt.Sprintln(args...)
	_ = p.stderr.Output(p.calldepth, s)
	panic(s)
}

func (p *logger) Panicf(format string, args ...interface{}) {
	if FATAL < p.level {
		return
	}
	s := fmt.Sprintf(format, args...)
	_ = p.stderr.Output(p.calldepth, s)
	panic(s)
}

func Trace(args ...interface{}) {
	logs.Trace(args)
}

func Tracef(format string, args ...interface{}) {
	logs.Tracef(format, args)
}

func Debug(args ...interface{}) {
	logs.Debug(args)
}

func Debugf(format string, args ...interface{}) {
	logs.Debugf(format, args)
}

func Info(args ...interface{}) {
	logs.Info(args)
}

func Infof(format string, args ...interface{}) {
	logs.Infof(format, args)
}

func Warn(args ...interface{}) {
	logs.Warn(args)
}

func Warnf(format string, args ...interface{}) {
	logs.Warnf(format, args)
}

func Error(args ...interface{}) {
	logs.Error(args)
}

func Errorf(format string, args ...interface{}) {
	logs.Errorf(format, args)
}

func Fatal(args ...interface{}) {
	logs.Fatal(args)
}

func Fatalf(format string, args ...interface{}) {
	logs.Fatalf(format, args)
}

func Panic(args ...interface{}) {
	logs.Panic(args)
}

func Panicf(format string, args ...interface{}) {
	logs.Panicf(format, args)
}
