package messagerepo

import (
	"context"

	"github.com/jackc/pgx/v4"
)

type logger struct {
	l Logger
}

func (l logger) Log(ctx context.Context, level pgx.LogLevel, msg string, data map[string]interface{}) {
	switch level {
	case pgx.LogLevelTrace:
		l.l.Debug(msg, data)
	case pgx.LogLevelDebug:
		l.l.Debug(msg, data)
	case pgx.LogLevelInfo:
		l.l.Info(msg, data)
	case pgx.LogLevelWarn:
		l.l.Warn(msg, data)
	case pgx.LogLevelError:
		l.l.Error(msg, data)
	case pgx.LogLevelNone:
		l.l.Debug(msg, data)
	}
}
