package logger

import "log/slog"

func GetLogger() *slog.Logger {
	return slog.New(slog.Default().Handler())
}
