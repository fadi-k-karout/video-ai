package logger

import (
	"log/slog"
	"os"
	"strings"

	"videoai/internal/config"
)

func Init(cfg *config.Config) *slog.Logger {
	var level slog.Level
	switch strings.ToLower(cfg.LogLevel) {
	case config.LogLevelDebug:
		level = slog.LevelDebug
	case config.LogLevelInfo:
		level = slog.LevelInfo
	case config.LogLevelWarn:
		level = slog.LevelWarn
	case config.LogLevelError:
		level = slog.LevelError
	default:
		level = slog.LevelInfo
	}

	var handler slog.Handler
	if strings.EqualFold(cfg.LogFormat, config.LogFormatText) {
		handler = slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: level})
	} else {
		handler = slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: level})
	}

	logger := slog.New(handler)
	slog.SetDefault(logger)
	return logger
}