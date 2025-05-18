package benchmark

import (
	"github.com/lk2023060901/logger-adapter-go/logger"
	"github.com/sirupsen/logrus"
	"go.uber.org/zap"
	"log/slog"
	"os"
	"testing"
)

func BenchmarkSlog(b *testing.B) {
	handler := slog.NewTextHandler(os.Stdout, nil)
	s := logger.NewSLogAdapter(slog.New(handler))

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s.Info("hello slog", "k", i)
	}
}

func BenchmarkZap(b *testing.B) {
	z, _ := zap.NewProduction()
	s := logger.NewZapAdapter(z.Sugar())

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s.Info("hello zap", "k", i)
	}
}

func BenchmarkLogrus(b *testing.B) {
	l := logrus.New()
	s := logger.NewLogrusAdapter(l)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s.Info("hello logrus", "k", i)
	}
}

func BenchmarkLoggerThroughAdapter(b *testing.B) {
	handler := slog.NewTextHandler(os.Stdout, nil)
	slogAdapter := logger.NewSLogAdapter(slog.New(handler))
	logger.SetLogger(slogAdapter)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		logger.Info("hello world", "k", i)
	}
}
