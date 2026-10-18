package log

import (
	"context"
	"testing"
)

func TestLoggerPackage(t *testing.T) {
	oldLoggers := _loggers
	defer func() { _loggers = oldLoggers }()
	_loggers = nil

	if n := NumberOfLoggers(); n != 0 {
		t.Fatalf("expected 0 loggers, got %d", n)
	}

	tl := &TestLogger{name: "logger-pkg-test"}
	AddLogger(tl)

	if n := NumberOfLoggers(); n != 1 {
		t.Fatalf("expected 1 logger, got %d", n)
	}

	// Test panic on duplicate logger
	func() {
		defer func() {
			if r := recover(); r == nil {
				t.Error("expected panic on duplicate logger name")
			}
		}()
		AddLogger(&TestLogger{name: "logger-pkg-test"})
	}()

	ctx := context.Background()

	// Test context functions
	if isDisabled(ctx) {
		t.Error("expected false for background context")
	}

	ctxDisabled := NewContextWithLoggingDisabled(ctx)
	if !isDisabled(ctxDisabled) {
		t.Error("expected true for disabled context")
	}

	ctxEnabled := NewContextWithLoggingEnabled(ctxDisabled)
	if isDisabled(ctxEnabled) {
		t.Error("expected false for re-enabled context")
	}

	// Logging with disabled context
	Debugf(ctxDisabled, "debug")
	Infof(ctxDisabled, "info")
	Warningf(ctxDisabled, "warn")
	Errorf(ctxDisabled, "error")
	Criticalf(ctxDisabled, "critical")

	if len(tl.Messages) != 0 {
		t.Fatalf("expected 0 messages logged when disabled, got %d", len(tl.Messages))
	}

	// Logging with enabled context
	Debugf(ctx, "debug %s", "1")
	Infof(ctx, "info %s", "2")
	Warningf(ctx, "warn %s", "3")
	Errorf(ctx, "error %s", "4")
	Criticalf(ctx, "critical %s", "5")

	if len(tl.Messages) != 5 {
		t.Fatalf("expected 5 messages logged when enabled, got %d", len(tl.Messages))
	}
}
