package utils

import (
	"log/slog"
	"os"
	"path/filepath"
)

var Logger *slog.Logger

// InitLogger initializes the structured logger to write to logs/app.log
func InitLogger() error {
	configDir, err := os.UserConfigDir()
	if err != nil {
		configDir = "."
	}
	logDir := filepath.Join(configDir, "pos-mini", "logs")
	if err := os.MkdirAll(logDir, 0755); err != nil {
		return err
	}

	logFile := filepath.Join(logDir, "app.log")
	file, err := os.OpenFile(logFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return err
	}

	// Use JSON handler for structured logging
	handler := slog.NewJSONHandler(file, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	})

	Logger = slog.New(handler)
	slog.SetDefault(Logger)

	return nil
}
