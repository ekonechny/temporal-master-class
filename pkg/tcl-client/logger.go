package tcl_client

import (
	"go.temporal.io/sdk/log"
	"go.uber.org/zap"
)

type zapLogger struct {
	*zap.SugaredLogger
}

func (l *zapLogger) Debug(msg string, keyvals ...interface{}) {
	l.Debugw(msg, keyvals...)
}

func (l *zapLogger) Info(msg string, keyvals ...interface{}) {
	l.Infow(msg, keyvals...)
}

func (l *zapLogger) Warn(msg string, keyvals ...interface{}) {
	l.Warnw(msg, keyvals...)
}

func (l *zapLogger) Error(msg string, keyvals ...interface{}) {
	l.Errorw(msg, keyvals...)
}

func NewLogger() (log.Logger, error) {
	logger, err := zap.NewProduction()
	if err != nil {
		return nil, err
	}

	return &zapLogger{logger.Sugar()}, nil
}
