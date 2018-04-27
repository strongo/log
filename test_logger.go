package log

import (
	"context"
)

// Level specifies log record type
type Level int

const (
	// LevelDebug for debug records
	LevelDebug Level = iota
	// LevelInfo for info records
	LevelInfo
	// LevelWarning for warnings
	LevelWarning
	// LevelError for errors
	LevelError
	// LevelCritical for critical records
	LevelCritical
)

// Message holds log record data
type Message struct {
	Level  Level
	Format string
	Args   []interface{}
}

// TestLogger is a test logger
type TestLogger struct {
	name     string
	Messages []Message
}

// Name returns name of the logger
func (logger *TestLogger) Name() string {
	return logger.name
}
func (logger *TestLogger) add(level Level, format string, args ...interface{}) {
	logger.Messages = append(logger.Messages, Message{Level: level, Format: format, Args: args})
}

// Debugf logs debug record
func (logger *TestLogger) Debugf(ctx context.Context, format string, args ...interface{}) {
	logger.add(LevelDebug, format, args...)
}

// Infof is like Debugf, but at Info level.
func (logger *TestLogger) Infof(ctx context.Context, format string, args ...interface{}) {
	logger.add(LevelInfo, format, args...)
}

// Warningf is like Debugf, but at Warning level.
func (logger *TestLogger) Warningf(ctx context.Context, format string, args ...interface{}) {
	logger.add(LevelWarning, format, args...)
}

// Errorf is like Debugf, but at Error level.
func (logger *TestLogger) Errorf(ctx context.Context, format string, args ...interface{}) {
	logger.add(LevelError, format, args...)
}

// Criticalf is like Debugf, but at Critical level.
func (logger *TestLogger) Criticalf(ctx context.Context, format string, args ...interface{}) {
	logger.add(LevelCritical, format, args...)
}
