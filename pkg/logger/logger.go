package logger

import (
	"context"
	"fmt"
	"log/slog"
	"os"

	"github.com/TDanilV16/merkle-tree-coursework/pkg/config"
)

var logger *slog.Logger

type Logger struct {
	*slog.Logger
}

func Default() *Logger {
	logger = slog.New(slog.NewTextHandler(os.Stdout, nil))
	slog.SetDefault(logger)
	return &Logger{}
}

func SetupFromConfig(cfg *config.Config) *Logger {
	if cfg == nil {
		return Default()
	}

	var level slog.Level

	switch cfg.Logging.Level {
	case "debug":
		level = slog.LevelDebug
	case "info":
		level = slog.LevelInfo
	case "warn":
		level = slog.LevelWarn
	case "error":
		level = slog.LevelError
	default:
		level = slog.LevelInfo
	}

	opts := &slog.HandlerOptions{
		Level:     level,
		AddSource: cfg.Logging.AddSource,
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			if a.Key == slog.SourceKey {
				source := a.Value.Any().(*slog.Source)

				shortFile := shortenPath(source.File)
				source.File = shortFile
				source.Function = shortenPath(source.Function)
			}

			return a
		},
	}

	var handler slog.Handler
	if cfg.Logging.Format == "json" {
		handler = slog.NewJSONHandler(os.Stdout, opts)
	} else {
		handler = slog.NewTextHandler(os.Stdout, opts)
	}

	logger = slog.New(handler)
	slog.SetDefault(logger)

	return &Logger{}
}

func (l *Logger) Info(ctx context.Context, msg string, args ...any) {
	logger.Log(ctx, slog.LevelInfo, msg, args...)
}

func (l *Logger) Infof(ctx context.Context, format string, args ...any) {
	l.Info(ctx, fmt.Sprintf(format, args...))
}

func (l *Logger) Debug(ctx context.Context, msg string, args ...any) {
	logger.Log(ctx, slog.LevelDebug, msg, args...)
}

func (l *Logger) Debugf(ctx context.Context, format string, args ...any) {
	l.Debug(ctx, fmt.Sprintf(format, args...))
}

func (l *Logger) Warn(ctx context.Context, msg string, args ...any) {
	logger.Log(ctx, slog.LevelWarn, msg, args...)
}

func (l *Logger) Warnf(ctx context.Context, format string, args ...any) {
	l.Warn(ctx, fmt.Sprintf(format, args...))
}

func (l *Logger) Error(ctx context.Context, msg string, args ...any) {
	logger.Log(ctx, slog.LevelError, msg, args...)
}

func (l *Logger) Errorf(ctx context.Context, format string, args ...any) {
	l.Error(ctx, fmt.Sprintf(format, args...))
}

func shortenPath(path string) string {
	const keepDirs = 2

	parts := []rune(path)
	count := 0
	for i := len(parts) - 1; i >= 0; i-- {
		if parts[i] == '/' {
			count++
			if count == keepDirs+1 {
				return string(parts[i+1:])
			}
		}
	}
	return path
}
