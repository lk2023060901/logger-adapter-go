package main

import (
	"context"
	"github.com/lk2023060901/logger-adapter-go/logger"
	"log/slog"
	"os"
)

func main() {
	// 创建一个包含请求 id 的 slog handler
	handler := slog.NewTextHandler(os.Stdout, nil)
	slogLogger := slog.New(handler)

	// 创建上下文日志器（绑定 request_id）
	requestLogger := slogLogger.With("request_id", "abc-123").With("Name", "Luke")
	ctx := logger.WithLogger(context.Background(), logger.NewSLogAdapter(requestLogger))

	// 日志输出将包含 request_id
	logger.InfoContext(ctx, "这是一个带 request_id 的日志", "step", "start")
}
