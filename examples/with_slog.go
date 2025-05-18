package main

import (
	"github.com/lk2023060901/logger-adapter-go/logger"
	"log/slog"
	"os"
)

func main() {
	slogLogger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	adapter := logger.NewSLogAdapter(slogLogger)
	logger.SetLogger(adapter)

	logger.Info("Hello, slog logger", "user", "liu", "role", "admin")
	logger.Debug("Debug message", "module", "init")
}
