package logger

import (
	"log/slog"
	"os"
)

type NullWriter struct{}

func (NullWriter) Write([]byte) (int, error) { return 0, nil }

func NullLogger() *slog.Logger {
	return slog.New(slog.NewTextHandler(NullWriter{}, nil))
}

func GetLogger() *slog.Logger {
	if os.Getenv("DEBUG") == "1" {
		return slog.New(slog.Default().Handler())
	}

	return NullLogger()
}
