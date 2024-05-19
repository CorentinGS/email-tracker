package logger

import (
	"log/slog"
	"os"
	"sync"
)

var (
	logger *Logger   //nolint:gochecknoglobals //Singleton
	once   sync.Once //nolint:gochecknoglobals //Singleton
)

type Logger struct {
	logLevel slog.Level
}

func GetLogger() *Logger {
	once.Do(func() {
		logger = &Logger{
			logLevel: slog.LevelInfo,
		}
	})
	return logger
}

func (l *Logger) SetLogLevel(level string) *Logger {
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
	return l
}

func (l *Logger) GetLogLevel() slog.Level {
	return l.logLevel
}

func (l *Logger) CreateGlobalHandler() {
	handler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level:     l.logLevel,
		AddSource: false,
	})

	logger := slog.New(handler)

	// slog.SetLogLoggerLevel(l.logLevel)

	slog.SetDefault(logger)
}
