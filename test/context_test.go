package test

import (
	"bytes"
	"context"
	"github.com/lk2023060901/logger-adapter-go/logger"
	"log/slog"
	"strings"
	"testing"
)

func TestContextLogger(t *testing.T) {
	var buf bytes.Buffer
	base := slog.NewTextHandler(&buf, nil)
	s := slog.New(base).With("request_id", "abc123")
	ctx := logger.WithLogger(context.Background(), logger.NewSLogAdapter(s))

	logger.InfoContext(ctx, "with ctx logger")

	out := buf.String()
	if !strings.Contains(out, "abc123") || !strings.Contains(out, "with ctx logger") {
		t.Errorf("context logger output unexpected: %s", out)
	}
}
