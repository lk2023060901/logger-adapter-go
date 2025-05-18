package main

import (
	"github.com/lk2023060901/logger-adapter-go/logger"
	"github.com/lk2023060901/logger-adapter-go/rotation"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"time"
)

func main() {
	writer, err := rotation.NewTimeRotationWriter(rotation.TimeRotationConfig{
		Filename: "logs/zap",         // 实际文件是 slog.当前年月日
		LinkName: "logs/zap.log",     // 始终指向最新日志
		MaxAge:   7 * 24 * time.Hour, // 保留 7 天
		Rotation: time.Hour,          // 每天轮转
		Stdout:   true,
	})
	if err != nil {
		panic(err)
	}

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
