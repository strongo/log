package log

import (
	"context"
	"fmt"
	"strings"
	"testing"
)

func TestNewPrinter(t *testing.T) {
	var output string
	printf := func(format string, a ...any) (n int, err error) {
		msg := fmt.Sprintf(format, a...)
		output += msg
		return len(msg), nil
	}

	l := NewPrinter("test-printer", printf)
	if name := l.Name(); name != "test-printer" {
		t.Errorf("expected name 'test-printer', got %s", name)
	}

	ctx := context.Background()

	l.Debugf(ctx, "debug %d", 1)
	if !strings.Contains(output, "DEBUG: debug 1\n") {
		t.Errorf("unexpected output: %q", output)
	}

	l.Infof(ctx, "info %d", 2)
	if !strings.Contains(output, "INFO: info 2\n") {
		t.Errorf("unexpected output: %q", output)
	}

	l.Warningf(ctx, "warning %d", 3)
	if !strings.Contains(output, "WARNING: warning 3\n") {
		t.Errorf("unexpected output: %q", output)
	}

	l.Errorf(ctx, "error %d", 4)
	if !strings.Contains(output, "ERROR: error 4\n") {
		t.Errorf("unexpected output: %q", output)
	}

	l.Criticalf(ctx, "critical %d", 5)
	if !strings.Contains(output, "CRITICAL: critical 5\n") {
		t.Errorf("unexpected output: %q", output)
	}
}
