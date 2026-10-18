package log

import (
	"context"
	"testing"
)

func TestTestLogger(t *testing.T) {
	tl := &TestLogger{name: "my-test-logger"}
	if name := tl.Name(); name != "my-test-logger" {
		t.Errorf("expected 'my-test-logger', got %q", name)
	}

	ctx := context.Background()

	tl.Debugf(ctx, "debug %s", "arg")
	tl.Infof(ctx, "info %s", "arg")
	tl.Warningf(ctx, "warn %s", "arg")
	tl.Errorf(ctx, "err %s", "arg")
	tl.Criticalf(ctx, "crit %s", "arg")

	if len(tl.Messages) != 5 {
		t.Fatalf("expected 5 messages, got %d", len(tl.Messages))
	}

	expectedLevels := []Level{LevelDebug, LevelInfo, LevelWarning, LevelError, LevelCritical}
	for i, expected := range expectedLevels {
		if tl.Messages[i].Level != expected {
			t.Errorf("msg %d: expected level %v, got %v", i, expected, tl.Messages[i].Level)
		}
	}
}
