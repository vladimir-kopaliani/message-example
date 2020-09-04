package logger

import (
	"context"

	"go.uber.org/zap"
)

// Logger ...
type Logger struct {
	logger *zap.SugaredLogger
}

// NewLogger ...
func NewLogger(ctx context.Context, isDebugMode bool) (*Logger, error) {
	var l *zap.Logger
	var err error

	if isDebugMode {
		l, err = zap.NewDevelopment()
	} else {
		l, err = zap.NewProduction()
	}
	if err != nil {
		return nil, err
	}

	lo := Logger{
		logger: l.Sugar(),
	}
	lo.Debug("Logger created")

	return &lo, nil
}

// Debug for fatall output
func (l Logger) Debug(msg ...interface{}) {
	l.logger.Debug(msg...)
}

// Info for default logging
func (l Logger) Info(msg ...interface{}) {
	l.logger.Info(msg...)
}

// Infof for formatted default logging
func (l Logger) Infof(message string, args ...interface{}) {
	l.logger.Infof(message, args...)
}

// Warn for important messages
func (l Logger) Warn(msg ...interface{}) {
	l.logger.Warn(msg...)
}

// Error for showing erros
func (l Logger) Error(msg ...interface{}) {
	l.logger.Error(msg...)
}

// Fatal for fatal error. Calls os.Exit(1).
func (l Logger) Fatal(msg ...interface{}) {
	l.logger.Fatal(msg...)
}

// Close flush buffer
func (l Logger) Close() error {
	l.Debug("Logger is closing")
	defer l.Debug("Logger closed")

	return l.logger.Sync()
}
