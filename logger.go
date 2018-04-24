package log

import (
	"fmt"
	"context"
)

type Logger interface {
	Name() string
	Debugf(c context.Context, format string, args ...interface{})
	Infof(c context.Context, format string, args ...interface{})
	Warningf(c context.Context, format string, args ...interface{})
	Errorf(c context.Context, format string, args ...interface{})
	Criticalf(c context.Context, format string, args ...interface{})
}

var _loggers []Logger

func NumberOfLoggers() int {
	return len(_loggers)
}

func AddLogger(logger Logger) {
	name := logger.Name()
	for _, l := range _loggers {
		if l.Name() == name {
			panic(fmt.Sprintf("Duplicate logger name: [%v], len(_loggers): %d", name, len(_loggers)))
		}
	}
	_loggers = append(_loggers, logger)
}

var loggingDisabled = "loggingDisabled" // TODO: Check core libs on how to do it proepry

func NewContextWithLoggingDisabled(c context.Context) context.Context {
	return context.WithValue(c, &loggingDisabled, true)
}

func NewContextWithLoggingEnabled(c context.Context) context.Context {
	return context.WithValue(c, &loggingDisabled, false)
}

func isDisabled(c context.Context) bool {
	if disabled := c.Value(&loggingDisabled); disabled != nil {
		return disabled.(bool)
	}
	return false
}

type LogFunc func(c context.Context, format string, args ...interface{})

func Debugf(c context.Context, format string, args ...interface{}) {
	if isDisabled(c) {
		return
	}
	for _, l := range _loggers {
		l.Debugf(c, format, args...)
	}
}

func Infof(c context.Context, format string, args ...interface{}) {
	if isDisabled(c) {
		return
	}
	for _, l := range _loggers {
		l.Infof(c, format, args...)
	}
}

func Warningf(c context.Context, format string, args ...interface{}) {
	if isDisabled(c) {
		return
	}
	for _, l := range _loggers {
		l.Warningf(c, format, args...)
	}
}

func Errorf(c context.Context, format string, args ...interface{}) {
	if isDisabled(c) {
		return
	}
	for _, l := range _loggers {
		l.Errorf(c, format, args...)
	}
}

func Criticalf(c context.Context, format string, args ...interface{}) {
	if isDisabled(c) {
		return
	}
	for _, l := range _loggers {
		l.Criticalf(c, format, args...)
	}
}
