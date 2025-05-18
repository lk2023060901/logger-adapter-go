package main

import (
	"github.com/lk2023060901/logger-adapter-go/logger"
	"github.com/lk2023060901/logger-adapter-go/rotation"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	writer := rotation.NewWriter(rotation.Config{
		Filename:   "logs/zap.log",
		MaxSize:    10, // 10MB
		MaxBackups: 5,
		MaxAge:     7, // 7天
		Compress:   true,
		Stdout:     true,
	})

	cfg := zap.NewProductionEncoderConfig()
	core := zapcore.NewCore(
		zapcore.NewConsoleEncoder(cfg),
		zapcore.AddSync(writer),
		zapcore.DebugLevel,
	)
	rawLogger := zap.New(core)
	sugar := rawLogger.Sugar()
	adapter := logger.NewZapAdapter(sugar)

	logger.SetLogger(adapter)
	logger.Info("Hello from zap", "user", "liu", "env", "dev")
	logger.Debug("Debug details", "module", "zap-demo")
}
