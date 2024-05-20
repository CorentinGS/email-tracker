package logger

import (
	"log/slog"
	"os"
)

type Logger struct {
	logLevel slog.Level
	l        *slog.Logger
}

type LoggerOption func(*Logger)

func WithLogLevel(level string) LoggerOption {
	return func(l *Logger) {
		switch level {
		case "debug":
			l.logLevel = slog.LevelDebug
		case "info":
			l.logLevel = slog.LevelInfo
		case "warn":
			l.logLevel = slog.LevelWarn
		case "error":
			l.logLevel = slog.LevelError
		default:
			l.logLevel = slog.LevelInfo
		}

		l.l = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: l.logLevel}))
	}
}

func NewLogger(options ...LoggerOption) *Logger {
	logger := &Logger{
		logLevel: slog.LevelInfo,
		l:        slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo})),
	}

	for _, option := range options {
		option(logger)
	}

	return logger
}

func (l *Logger) GetLogLevel() slog.Level {
	return l.logLevel
}

func (l *Logger) SetAsGlobalHandler() *Logger {

	slog.SetDefault(l.l)

	return l
}
