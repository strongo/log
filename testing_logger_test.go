package log

import (
	"testing"
)

func TestNewTestingLogger(t *testing.T) {
	logger := NewTestingLogger(t, "")
	if name := logger.Name(); name != "*testingT" {
		t.Errorf("logger.Name() = %s, want *testingT", name)
	}

	const loggerName = "test-logger-1"
	logger = NewTestingLogger(t, loggerName)
	if name := logger.Name(); name != loggerName {
		t.Errorf("logger.Name() = %s, want %s", name, loggerName)
	}
	logger.Debugf(nil, "Hello, %s!", "World")
	logger.Infof(nil, "Hello, %s!", "World")
	logger.Warningf(nil, "Hello, %s!", "World")
	logger.Errorf(nil, "Hello, %s!", "World")
	logger.Criticalf(nil, "Hello, %s!", "World")

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	_ = NewTestingLogger(nil, loggerName)
}
