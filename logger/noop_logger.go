package logger

import (
	"context"
)

// noopLogger 是一个不做任何事情的 Logger 实现，适合作为默认值。
type noopLogger struct{}

// NewNoopLogger 返回一个空实现的日志器
func NewNoopLogger() Logger {
	return &noopLogger{}
}

func (n *noopLogger) Trace(format string, args ...any) {}
func (n *noopLogger) Debug(msg string, args ...any)    {}
func (n *noopLogger) Info(msg string, args ...any)     {}
func (n *noopLogger) Warn(msg string, args ...any)     {}
func (n *noopLogger) Error(msg string, args ...any)    {}
func (n *noopLogger) Fatal(msg string, args ...any)    {}
func (n *noopLogger) Panic(msg string, args ...any)    {}

func (n *noopLogger) WithContext(ctx context.Context) Logger {
	return n
}
