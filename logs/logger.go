package logs

import (
	"context"
	"fmt"
	"io"
	"os"
	"runtime"
	"strings"
	"time"
)

var wd string

func GetWd() string {
	if wd != "" {
		return wd
	}
	s, e := os.Getwd()
	if e != nil {
		wd = "/"
	}
	wd = s
	return wd
}

func GetRelativePath(s string) string {
	return strings.TrimPrefix(s, GetWd()+"/")
}

func GetCaller(calldepth int) (string, int) {
	_, f, l, _ := runtime.Caller(calldepth)
	rf := GetRelativePath(f)
	return rf, l
}

type Level int

// 级别从低到高
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
	Trace(ctx context.Context, args ...interface{})

	Tracef(ctx context.Context, format string, args ...interface{})

	Debug(ctx context.Context, args ...interface{})

	Debugf(ctx context.Context, format string, args ...interface{})

	Info(ctx context.Context, args ...interface{})

	Infof(ctx context.Context, format string, args ...interface{})

	Warn(ctx context.Context, args ...interface{})

	Warnf(ctx context.Context, format string, args ...interface{})

	Error(ctx context.Context, args ...interface{})

	Errorf(ctx context.Context, format string, args ...interface{})

	Fatal(ctx context.Context, args ...interface{})

	Fatalf(ctx context.Context, format string, args ...interface{})

	Panic(ctx context.Context, args ...interface{})

	Panicf(ctx context.Context, format string, args ...interface{})
}

type LogFormatter interface {
	Format(ctx context.Context, s string) string
}

type logFormatter struct {
	prefix    string
	calldepth int
}

func (p *logFormatter) Format(ctx context.Context, s string) string {
	f, l := GetCaller(p.calldepth)
	return fmt.Sprintf("%s  %s  %s[%d] - %s\n", p.prefix, time.Now().Format("2006-01-02 15:04:05.000"), f, l, s)
}

func NewLogger(options ...LoggerOption) Logger {
	l := &logger{}
	for i, _ := range options {
		options[i](l)
	}
	if l.Formatter == nil {
		l.Formatter = &logFormatter{
			prefix:    "[GO]",
			calldepth: 2,
		}
	}
	if l.Writer == nil {
		l.Writer = os.Stdout
	}
	if l.WriterError == nil {
		l.WriterError = os.Stderr
	}
	return l
}

type logger struct {
	level       Level
	Formatter   LogFormatter
	Writer      io.Writer
	WriterError io.Writer
}

func (p *logger) Trace(ctx context.Context, args ...interface{}) {
	if TRACE < p.level {
		return
	}
	s := p.Formatter.Format(ctx, fmt.Sprint(args...))
	_, _ = p.Writer.Write([]byte(s))
}

func (p *logger) Tracef(ctx context.Context, format string, args ...interface{}) {
	if TRACE < p.level {
		return
	}
	s := p.Formatter.Format(ctx, fmt.Sprintf(format, args...))
	_, _ = p.Writer.Write([]byte(s))
}

func (p *logger) Debug(ctx context.Context, args ...interface{}) {
	if DEBUG < p.level {
		return
	}
	s := p.Formatter.Format(ctx, fmt.Sprint(args...))
	_, _ = p.Writer.Write([]byte(s))
}

func (p *logger) Debugf(ctx context.Context, format string, args ...interface{}) {
	if DEBUG < p.level {
		return
	}
	s := p.Formatter.Format(ctx, fmt.Sprintf(format, args...))
	_, _ = p.Writer.Write([]byte(s))
}

func (p *logger) Info(ctx context.Context, args ...interface{}) {
	if INFO < p.level {
		return
	}
	s := p.Formatter.Format(ctx, fmt.Sprint(args...))
	_, _ = p.Writer.Write([]byte(s))
}

func (p *logger) Infof(ctx context.Context, format string, args ...interface{}) {
	if INFO < p.level {
		return
	}
	s := p.Formatter.Format(ctx, fmt.Sprintf(format, args...))
	_, _ = p.Writer.Write([]byte(s))
}

func (p *logger) Warn(ctx context.Context, args ...interface{}) {
	if WARN < p.level {
		return
	}
	s := p.Formatter.Format(ctx, fmt.Sprint(args...))
	_, _ = p.Writer.Write([]byte(s))
}

func (p *logger) Warnf(ctx context.Context, format string, args ...interface{}) {
	if WARN < p.level {
		return
	}
	s := p.Formatter.Format(ctx, fmt.Sprintf(format, args...))
	_, _ = p.Writer.Write([]byte(s))
}

func (p *logger) Error(ctx context.Context, args ...interface{}) {
	if ERROR < p.level {
		return
	}
	s := p.Formatter.Format(ctx, fmt.Sprint(args...))
	_, _ = p.Writer.Write([]byte(s))
}

func (p *logger) Errorf(ctx context.Context, format string, args ...interface{}) {
	if ERROR < p.level {
		return
	}
	s := p.Formatter.Format(ctx, fmt.Sprintf(format, args...))
	_, _ = p.Writer.Write([]byte(s))
}

func (p *logger) Fatal(ctx context.Context, args ...interface{}) {
	if FATAL < p.level {
		return
	}
	s := p.Formatter.Format(ctx, fmt.Sprint(args...))
	_, _ = p.Writer.Write([]byte(s))
	os.Exit(1)
}

func (p *logger) Fatalf(ctx context.Context, format string, args ...interface{}) {
	if FATAL < p.level {
		return
	}
	s := p.Formatter.Format(ctx, fmt.Sprintf(format, args...))
	_, _ = p.Writer.Write([]byte(s))
	os.Exit(1)
}

func (p *logger) Panic(ctx context.Context, args ...interface{}) {
	if FATAL < p.level {
		return
	}
	s := p.Formatter.Format(ctx, fmt.Sprint(args...))
	_, _ = p.Writer.Write([]byte(s))
	panic(s)
}

func (p *logger) Panicf(ctx context.Context, format string, args ...interface{}) {
	if FATAL < p.level {
		return
	}
	s := p.Formatter.Format(ctx, fmt.Sprintf(format, args...))
	_, _ = p.Writer.Write([]byte(s))
	panic(s)
}
