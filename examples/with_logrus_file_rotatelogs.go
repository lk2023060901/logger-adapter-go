package main

import (
	"github.com/lk2023060901/logger-adapter-go/logger"
	"github.com/lk2023060901/logger-adapter-go/rotation"
	"github.com/sirupsen/logrus"
	"time"
)

func main() {
	writer, err := rotation.NewTimeRotationWriter(rotation.TimeRotationConfig{
		Filename: "logs/logrus",      // 实际文件是 slog.当前年月日
		LinkName: "logs/logrus.log",  // 始终指向最新日志
		MaxAge:   7 * 24 * time.Hour, // 保留 7 天
		Rotation: time.Second,        // 每天轮转
		Stdout:   true,
		Pattern:  ".%Y-%m-%d.log",
	})
	if err != nil {
		panic(err)
	}

	log := logrus.New()
	log.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
	log.SetLevel(logrus.InfoLevel)
	log.SetOutput(writer)

	adapter := logger.NewLogrusAdapter(log)
	logger.SetLogger(adapter)

	logger.Info("Logrus example", "version", "v1.0.0", "mode", "demo")
	logger.Debug("Debug info", "module", "logrus")
}
