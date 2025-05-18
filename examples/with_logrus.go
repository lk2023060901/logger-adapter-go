package main

import (
	"github.com/lk2023060901/logger-adapter-go/logger"
	"github.com/sirupsen/logrus"
)

func main() {
	log := logrus.New()
	log.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
	log.SetLevel(logrus.InfoLevel)

	adapter := logger.NewLogrusAdapter(log)
	logger.SetLogger(adapter)

	logger.Info("Logrus example", "version", "v1.0.0", "mode", "demo")
	logger.Debug("Debug info", "module", "logrus")
}
