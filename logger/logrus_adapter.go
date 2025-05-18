package logger

import "github.com/sirupsen/logrus"

type LogrusAdapter struct {
	logger *logrus.Logger
}

func NewLogrusAdapter(l *logrus.Logger) *LogrusAdapter {
	return &LogrusAdapter{logger: l}
}

func (l *LogrusAdapter) Trace(msg string, args ...any) {

}

func (l *LogrusAdapter) Debug(msg string, args ...any) {
	l.logger.WithFields(convertFields(args...)).Debug(msg)
}

func (l *LogrusAdapter) Info(msg string, args ...any) {
	l.logger.WithFields(convertFields(args...)).Info(msg)
}

func (l *LogrusAdapter) Warn(msg string, args ...any) {
	l.logger.WithFields(convertFields(args...)).Warn(msg)
}

func (l *LogrusAdapter) Error(msg string, args ...any) {
	l.logger.WithFields(convertFields(args...)).Error(msg)
}

func (l *LogrusAdapter) Fatal(msg string, args ...any) {
	l.logger.WithFields(convertFields(args...)).Fatal(msg)
}

func (l *LogrusAdapter) Panic(msg string, args ...any) {
	l.logger.WithFields(convertFields(args...)).Panic(msg)
}

// 将 args 转换成 logrus.Fields
func convertFields(args ...any) logrus.Fields {
	fields := logrus.Fields{}
	for i := 0; i < len(args)-1; i += 2 {
		key, ok := args[i].(string)
		if !ok {
			continue
		}
		fields[key] = args[i+1]
	}
	return fields
}
