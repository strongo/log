package log

import (
	"context"
	"testing"
)

func TestStrongoLogger(t *testing.T) {
	oldLoggers := _loggers
	defer func() { _loggers = oldLoggers }()

	tl := &TestLogger{name: "strongo-test-logger"}
	_loggers = []Logger{tl}

	sl := StrongoLogger{}
	ctx := context.Background()

	sl.Debugf(ctx, "debug %d", 1)
	sl.Infof(ctx, "info %d", 2)
	sl.Warningf(ctx, "warning %d", 3)
	sl.Errorf(ctx, "error %d", 4)
	sl.Criticalf(ctx, "critical %d", 5)

	if len(tl.Messages) != 5 {
		t.Fatalf("expected 5 messages, got %d", len(tl.Messages))
	}
}
