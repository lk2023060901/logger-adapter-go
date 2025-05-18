package logger

import "context"

type ctxLoggerKey struct{}

func WithLogger(ctx context.Context, logger Logger) context.Context {
	return context.WithValue(ctx, ctxLoggerKey{}, logger)
}

func FromContext(ctx context.Context) Logger {
	if ctx == nil {
		return GetLogger()
	}
	if l, ok := ctx.Value(ctxLoggerKey{}).(Logger); ok {
		return l
	}
	return GetLogger()
}

func TraceContext(ctx context.Context, msg string, args ...any) {
	FromContext(ctx).Trace(msg, args...)
}

func DebugContext(ctx context.Context, msg string, args ...any) {
	FromContext(ctx).Debug(msg, args...)
}

func InfoContext(ctx context.Context, msg string, args ...any) {
	FromContext(ctx).Info(msg, args...)
}

func WarnContext(ctx context.Context, msg string, args ...any) {
	FromContext(ctx).Warn(msg, args...)
}

func ErrorContext(ctx context.Context, msg string, args ...any) {
	FromContext(ctx).Error(msg, args...)
}

func FatalContext(ctx context.Context, msg string, args ...any) {
	FromContext(ctx).Fatal(msg, args...)
}

func PanicContext(ctx context.Context, msg string, args ...any) {
	FromContext(ctx).Panic(msg, args...)
}
