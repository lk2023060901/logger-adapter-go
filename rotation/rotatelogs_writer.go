package rotation

import (
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"io"
	"os"
	"time"
)

// TimeRotationConfig 配置项
type TimeRotationConfig struct {
	Filename string        // 日志文件路径（不包含扩展名）
	LinkName string        // 当前日志的软链接名
	MaxAge   time.Duration // 日志最大保留时间
	Rotation time.Duration // 日志切割间隔
	Stdout   bool          // 是否同时输出到终端
	Pattern  string        // 文件后缀
}

type TimeRotationWriter struct {
	io.Writer
	Raw *rotatelogs.RotateLogs
}

func NewDefaultTimeRotationConfig(filename string) TimeRotationConfig {
	return TimeRotationConfig{
		Filename: filename,
		LinkName: filename,
		MaxAge:   90,
		Rotation: time.Hour,
		Stdout:   true,
		Pattern:  ".%Y-%m-%d_%H.log",
	}
}

func NewTimeRotationWriter(cfg TimeRotationConfig) (*TimeRotationWriter, error) {
	raw, err := rotatelogs.New(
		cfg.Filename+cfg.Pattern,
		rotatelogs.WithLinkName(cfg.LinkName),
		rotatelogs.WithRotationTime(cfg.Rotation),
		rotatelogs.WithMaxAge(cfg.MaxAge),
	)
	if err != nil {
		return nil, err
	}

	var w io.Writer = raw
	if cfg.Stdout {
		w = io.MultiWriter(os.Stdout, raw)
	}

	return &TimeRotationWriter{Writer: w, Raw: raw}, nil
}
