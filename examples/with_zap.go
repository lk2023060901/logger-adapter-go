package main

import (
	"github.com/lk2023060901/logger-adapter-go/logger"
	"go.uber.org/zap"
)

func main() {
	rawLogger, _ := zap.NewProduction()
	sugar := rawLogger.Sugar()
	adapter := logger.NewZapAdapter(sugar)

	logger.SetLogger(adapter)
	logger.Info("Hello from zap", "user", "liu", "env", "dev")
	logger.Debug("Debug details", "module", "zap-demo")
}
