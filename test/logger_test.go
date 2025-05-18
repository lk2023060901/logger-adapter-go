package test

import (
	"bytes"
	"github.com/lk2023060901/logger-adapter-go/logger"
	"log/slog"
	"strings"
	"testing"
)

func TestLoggerOutput(t *testing.T) {
	var buf bytes.Buffer
	handler := slog.NewTextHandler(&buf, nil)
	l := logger.NewSLogAdapter(slog.New(handler))

	logger.SetLogger(l)
	logger.Info("hello world", "key", "value")

	out := buf.String()
	if !strings.Contains(out, "hello world") || !strings.Contains(out, "key=value") {
		t.Errorf("unexpected log output: %s", out)
	}
}

func TestSetLoggerIsThreadSafe(t *testing.T) {
	done := make(chan struct{})
	for i := 0; i < 100; i++ {
		go func() {
			l := logger.NewSLogAdapter(slog.New(slog.NewTextHandler(&bytes.Buffer{}, nil)))
			logger.SetLogger(l)
			_ = logger.GetLogger()
			done <- struct{}{}
		}()
	}
	for i := 0; i < 100; i++ {
		<-done
	}
}
