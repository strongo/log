package log

import (
	"context"
)

// NewPrinter creates new logger that prints to provided function
// Example: NewPrinter("stdout", fmt.Printf)
func NewPrinter(name string, printf func(format string, a ...any) (n int, err error)) Logger {
	return printer{name, printf}
}

type printer struct {
	name   string
	printf func(format string, a ...any) (n int, err error)
}

func (p printer) Name() string {
	return p.name
}

func (p printer) Debugf(c context.Context, format string, args ...interface{}) {
	p.log("DEBUG", format, args...)
}

func (p printer) Infof(c context.Context, format string, args ...interface{}) {
	p.log("INFO", format, args...)
}

func (p printer) Warningf(c context.Context, format string, args ...interface{}) {
	p.log("WARNING", format, args...)
}

func (p printer) Errorf(c context.Context, format string, args ...interface{}) {
	p.log("ERROR", format, args...)
}

func (p printer) Criticalf(c context.Context, format string, args ...interface{}) {
	p.log("CRITICAL", format, args...)
}

func (p printer) log(level, format string, args ...interface{}) {
	_, _ = p.printf(level+": "+format+"\n", args...)
}
