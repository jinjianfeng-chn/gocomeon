package logs

import "io"

type LoggerOption func(*logger)

func WithLogLevel(level Level) LoggerOption {
	return func(l *logger) {
		l.level = level
	}
}

func WithLogFormatter(logFormatter LogFormatter) LoggerOption {
	return func(l *logger) {
		l.Formatter = logFormatter
	}
}

func WithLogWriter(w io.Writer) LoggerOption {
	return func(l *logger) {
		l.Writer = w
	}
}

func WithLogWriterError(w io.Writer) LoggerOption {
	return func(l *logger) {
		l.WriterError = w
	}
}
