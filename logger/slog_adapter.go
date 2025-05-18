package logger

import "log/slog"

type SLogAdapter struct {
	logger *slog.Logger
}

func NewSLogAdapter(l *slog.Logger) *SLogAdapter {
	return &SLogAdapter{logger: l}
}

func (s *SLogAdapter) Trace(msg string, args ...any) {

}

func (s *SLogAdapter) Debug(msg string, args ...any) {
	s.logger.Debug(msg, args...)
}

func (s *SLogAdapter) Info(msg string, args ...any) {
	s.logger.Info(msg, args...)
}

func (s *SLogAdapter) Warn(msg string, args ...any) {
	s.logger.Warn(msg, args...)
}

func (s *SLogAdapter) Error(msg string, args ...any) {
	s.logger.Error(msg, args...)
}

func (s *SLogAdapter) Fatal(msg string, args ...any) {
}

func (s *SLogAdapter) Panic(msg string, args ...any) {

}
