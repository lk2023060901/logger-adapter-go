package main

import (
	"github.com/lk2023060901/logger-adapter-go/logger"
	"github.com/lk2023060901/logger-adapter-go/rotation"
	"log/slog"
)

func main() {
	writer := rotation.NewWriter(rotation.Config{
		Filename:   "logs/slog.log",
		MaxSize:    10, // 10MB
		MaxBackups: 5,
		MaxAge:     7, // 7天
		Compress:   true,
		Stdout:     true,
	})

	handler := slog.NewTextHandler(writer, nil)
	slogLogger := slog.New(handler)
	adapter := logger.NewSLogAdapter(slogLogger)
	logger.SetLogger(adapter)

	logger.Info("slog with file rotation", "user", "liu")
	logger.Debug("rotation works", "file", "slog.log")
}
