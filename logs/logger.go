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

type logger struct {
	level     Level
	prefix    string
	calldepth int
	stdout    *log.Logger
	stderr    *log.Logger
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
	s := fmt.Sprintf(format, args...)
	_ = p.stdout.Output(p.calldepth, s)
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
