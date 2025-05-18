package rotation

import (
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"os"
	"path/filepath"
)

// Config 配置项
type Config struct {
	Filename   string // 日志文件路径
	MaxSize    int    // 单位：MB
	MaxAge     int    // 保留天数
	MaxBackups int    // 保留旧文件个数
	Compress   bool   // 是否压缩
	Stdout     bool   // 同时输出到终端
}

func NewDefaultLumberjackConfig(filename string) Config {
	return Config{
		Filename:   filename,
		MaxSize:    512,
		MaxAge:     90,
		MaxBackups: 2200,
		Compress:   true,
		Stdout:     true,
	}
}

// NewLumberjackWriter 创建符合 io.Writer 接口的输出
func NewLumberjackWriter(cfg Config) io.Writer {
	fileWriter := &lumberjack.Logger{
		Filename:   filepath.Clean(cfg.Filename),
		MaxSize:    cfg.MaxSize,
		MaxAge:     cfg.MaxAge,
		MaxBackups: cfg.MaxBackups,
		Compress:   cfg.Compress,
	}

	if cfg.Stdout {
		return io.MultiWriter(os.Stdout, fileWriter)
	}
	return fileWriter
}
