package log

import (
	"context"
	"testing"
)

// NewTestingLogger returns a logger that uses `testing.T` logger
func NewTestingLogger(t *testing.T) Logger {
	return testingLogger{t: t}
}

// testingLogger is tests
type testingLogger struct {
	t *testing.T
}

// Name returns the name of the logger.
func (testingLogger) Name() string {
	return "testingLogger"
}

// Debugf logs a debug message.
func (logger testingLogger) Debugf(ctx context.Context, format string, args ...interface{}) {
	logger.t.Logf("Debug: "+format, args...)
}

// Infof is like Debugf, but at Info level.
func (logger testingLogger) Infof(ctx context.Context, format string, args ...interface{}) {
	logger.t.Logf("Info: "+format, args...)
}

// Warningf is like Debugf, but at Warning level.
func (logger testingLogger) Warningf(ctx context.Context, format string, args ...interface{}) {
	logger.t.Logf("Warning: "+format, args...)
}

// Errorf is like Debugf, but at Error level.
func (logger testingLogger) Errorf(ctx context.Context, format string, args ...interface{}) {
	logger.t.Logf("Error: "+format, args...)
}

// Criticalf is like Debugf, but at Critical level.
func (logger testingLogger) Criticalf(ctx context.Context, format string, args ...interface{}) {
	logger.t.Logf("Critical: "+format, args...)
}
