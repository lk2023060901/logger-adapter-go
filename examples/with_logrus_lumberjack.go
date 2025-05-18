package main

import (
	"github.com/lk2023060901/logger-adapter-go/logger"
	"github.com/lk2023060901/logger-adapter-go/rotation"
	"github.com/sirupsen/logrus"
)

func main() {
	writer := rotation.NewWriter(rotation.Config{
		Filename:   "logs/logrus.log",
		MaxSize:    10, // 10MB
		MaxBackups: 5,
		MaxAge:     7, // 7天
		Compress:   true,
		Stdout:     true,
	})

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
