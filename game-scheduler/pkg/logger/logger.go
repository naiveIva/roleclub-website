package logger

import (
	"log/slog"
	"os"
)

func Init() *slog.Logger {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	return logger
}

func Error(err error) slog.Attr {
	return slog.String("error", err.Error())
}
