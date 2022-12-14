package logs

type LoggerOption func(*logger)

func WithLogLevel(level Level) LoggerOption {
	return func(l *logger) {
		l.level = level
	}
}

func WithLogPrefix(prefix string) LoggerOption {
	return func(l *logger) {
		l.prefix = prefix
	}
}

func WithLogCalldepth(calldepth int) LoggerOption {
	return func(l *logger) {
		l.calldepth = calldepth
	}
}

func WithLogFlat(flag int) LoggerOption {
	return func(l *logger) {
		l.flag = flag
	}
}
