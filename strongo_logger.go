package log

import "context"

type StrongoLogger struct {
}

func (logger StrongoLogger) Infof(c context.Context, format string, args ...interface{}) {
	Infof(c, format, args...)
}

func (logger StrongoLogger) Errorf(c context.Context, format string, args ...interface{}) {
	Errorf(c, format, args...)
}

func (logger StrongoLogger) Debugf(c context.Context, format string, args ...interface{}) {
	Debugf(c, format, args...)
}

func (logger StrongoLogger) Warningf(c context.Context, format string, args ...interface{}) {
	Warningf(c, format, args...)
}

func (logger StrongoLogger) Criticalf(c context.Context, format string, args ...interface{}) {
	Criticalf(c, format, args...)
}
