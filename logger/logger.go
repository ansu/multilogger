package logger

import (
	"context"
	"os"

	"github.com/sirupsen/logrus"
	"go.uber.org/zap"
)

type Logger interface {
	Debug(msg string, fields map[string]interface{})
	Info(msg string, fields map[string]interface{})
	Warn(msg string, fields map[string]interface{})
	Error(msg string, fields map[string]interface{})
	Fatal(msg string, fields map[string]interface{})
}

type LogrusLogger struct {
	logger *logrus.Logger
	ctx    context.Context
}

func NewLogrusLogger(loggerType string, ctx context.Context) *LogrusLogger {
	logger := logrus.New()
	logger.Out = os.Stdout

	return &LogrusLogger{logger: logger, ctx: ctx}
}

func (l *LogrusLogger) Debug(msg string, fields map[string]interface{}) {

	l.logger.WithFields(fields).Debug(msg)
}

func (l *LogrusLogger) Info(msg string, fields map[string]interface{}) {
	l.addContextCommonFields(fields)

	l.logger.WithFields(fields).Info(msg)
}

func (l *LogrusLogger) Warn(msg string, fields map[string]interface{}) {
	l.logger.WithFields(fields).Warn(msg)
}

func (l *LogrusLogger) Error(msg string, fields map[string]interface{}) {
	l.logger.WithFields(fields).Error(msg)
}

func (l *LogrusLogger) Fatal(msg string, fields map[string]interface{}) {
	l.logger.WithFields(fields).Fatal(msg)
}

func (l *LogrusLogger) addContextCommonFields(fields map[string]interface{}) {
	if l.ctx != nil {
		for k, v := range l.ctx.Value("commonFields").(map[string]interface{}) {
			if _, ok := fields[k]; !ok {
				fields[k] = v
			}
		}
	}
}

type ZapLogger struct {
	logger *zap.Logger
	ctx    context.Context
}

func NewZapLogger(loggerType string, ctx context.Context) *ZapLogger {
	logger, _ := zap.NewProduction()

	return &ZapLogger{logger: logger, ctx: ctx}
}

func (l *ZapLogger) Debug(msg string, fields map[string]interface{}) {
	l.addContextCommonFields(fields)

	l.logger.Debug("", zap.Any("args", fields))
}

func (l *ZapLogger) Info(msg string, fields map[string]interface{}) {
	l.addContextCommonFields(fields)

	l.logger.Info("", zap.Any("args", fields))
}

func (l *ZapLogger) Warn(msg string, fields map[string]interface{}) {
	l.addContextCommonFields(fields)

	l.logger.Warn("", zap.Any("args", fields))
}

func (l *ZapLogger) Error(msg string, fields map[string]interface{}) {
	l.addContextCommonFields(fields)

	l.logger.Error("", zap.Any("args", fields))
}

func (l *ZapLogger) Fatal(msg string, fields map[string]interface{}) {
	l.addContextCommonFields(fields)

	l.logger.Fatal("", zap.Any("args", fields))
}

func (l *ZapLogger) addContextCommonFields(fields map[string]interface{}) {
	if l.ctx != nil {
		for k, v := range l.ctx.Value("commonFields").(map[string]interface{}) {
			if _, ok := fields[k]; !ok {
				fields[k] = v
			}
		}
	}
}

type LoggerWrapper struct {
	logger Logger
}

func NewLoggerWrapper(loggerType string, ctx context.Context) *LoggerWrapper {
	var logger Logger

	switch loggerType {
	case "logrus":
		logger = NewLogrusLogger(loggerType, ctx)
	case "zap":
		logger = NewZapLogger(loggerType, ctx)
	default:
		logger = NewLogrusLogger(loggerType, ctx)
	}

	return &LoggerWrapper{logger: logger}
}

func (lw *LoggerWrapper) Debug(msg string, fields map[string]interface{}) {

	lw.logger.Debug(msg, fields)
}

func (lw *LoggerWrapper) Info(msg string, fields map[string]interface{}) {
	lw.logger.Info(msg, fields)
}

func (lw *LoggerWrapper) Warn(msg string, fields map[string]interface{}) {
	lw.logger.Warn(msg, fields)
}

func (lw *LoggerWrapper) Error(msg string, fields map[string]interface{}) {
	lw.logger.Error(msg, fields)
}

func (lw *LoggerWrapper) Fatal(msg string, fields map[string]interface{}) {
	lw.logger.Fatal(msg, fields)
}
