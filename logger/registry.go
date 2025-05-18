package logger

import "sync"

var (
	mu            sync.RWMutex
	defaultLogger Logger = &noopLogger{}
)

func SetLogger(logger Logger) {
	mu.Lock()
	defer mu.Unlock()
	defaultLogger = logger
}

// GetLogger 获取当前默认 logger（线程安全）
func GetLogger() Logger {
	mu.RLock()
	defer mu.RUnlock()
	return defaultLogger
}

func Trace(msg string, args ...any) {
	mu.Lock()
	defer mu.Unlock()
	if defaultLogger != nil {
		defaultLogger.Trace(msg, args...)
	}
}

func Debug(msg string, args ...any) {
	mu.Lock()
	defer mu.Unlock()
	if defaultLogger != nil {
		defaultLogger.Debug(msg, args...)
	}
}

func Info(msg string, args ...any) {
	mu.Lock()
	defer mu.Unlock()
	if defaultLogger != nil {
		defaultLogger.Info(msg, args...)
	}
}

func Warn(msg string, args ...any) {
	mu.Lock()
	defer mu.Unlock()
	if defaultLogger != nil {
		defaultLogger.Warn(msg, args...)
	}
}

func Error(msg string, args ...any) {
	mu.Lock()
	defer mu.Unlock()
	if defaultLogger != nil {
		defaultLogger.Error(msg, args...)
	}
}

func Fatal(msg string, args ...any) {
	mu.Lock()
	defer mu.Unlock()
	if defaultLogger != nil {
		defaultLogger.Fatal(msg, args...)
	}
}

func Panic(msg string, args ...any) {
	mu.Lock()
	defer mu.Unlock()
	if defaultLogger != nil {
		defaultLogger.Panic(msg, args...)
	}
}
