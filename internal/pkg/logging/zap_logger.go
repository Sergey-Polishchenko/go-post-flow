package logging

import "go.uber.org/zap"

type ZapLogger struct {
	logger *zap.Logger
}

func NewZapLogger() Logger {
	logger, _ := zap.NewProduction()
	return &ZapLogger{logger: logger}
}

func (z *ZapLogger) Info(msg string, fields ...interface{}) {
	z.logger.Info(msg, zap.Any("fields", fields))
}

func (z *ZapLogger) Debug(msg string, fields ...interface{}) {
	z.logger.Debug(msg, zap.Any("fields", fields))
}

func (z *ZapLogger) Error(msg string, fields ...interface{}) {
	z.logger.Error(msg, zap.Any("fields", fields))
}
