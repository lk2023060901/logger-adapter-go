package main

import (
	"github.com/lk2023060901/logger-adapter-go/logger"
	"github.com/lk2023060901/logger-adapter-go/rotation"
	"log/slog"
	"time"
)

func main() {
	writer, err := rotation.NewTimeRotationWriter(rotation.TimeRotationConfig{
		Filename: "logs/slog",        // 实际文件是 slog.当前年月日
		LinkName: "logs/slog.log",    // 始终指向最新日志
		MaxAge:   7 * 24 * time.Hour, // 保留 7 天
		Rotation: time.Hour,          // 每天轮转
		Stdout:   true,
		Pattern:  "%Y-%m-%d",
	})
	if err != nil {
		panic(err)
	}

	handler := slog.NewTextHandler(writer, nil)
	slogLogger := slog.New(handler)
	adapter := logger.NewSLogAdapter(slogLogger)
	logger.SetLogger(adapter)

	logger.Info("rotatelogs 测试", "daily", "yes", "demo", "rotatelogs")
}
