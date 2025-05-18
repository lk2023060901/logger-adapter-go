package logger

import "go.uber.org/zap"

type ZapAdapter struct {
	logger *zap.SugaredLogger
}

func NewZapAdapter(l *zap.SugaredLogger) *ZapAdapter {
	return &ZapAdapter{logger: l}
}

func (s *ZapAdapter) Trace(msg string, args ...any) {
}

func (s *ZapAdapter) Debug(msg string, args ...any) {
	s.logger.Debugw(msg, args...)
}

func (s *ZapAdapter) Info(msg string, args ...any) {
	s.logger.Infow(msg, args...)
}

func (s *ZapAdapter) Warn(msg string, args ...any) {
	s.logger.Warnw(msg, args...)
}

func (s *ZapAdapter) Error(msg string, args ...any) {
	s.logger.Errorw(msg, args...)
}

func (s *ZapAdapter) Fatal(msg string, args ...any) {
	s.logger.Fatalw(msg, args...)
}

func (s *ZapAdapter) Panic(msg string, args ...any) {
	s.logger.Panicw(msg, args...)
}
