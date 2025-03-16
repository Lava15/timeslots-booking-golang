package config

import (
	"log/slog"

	"github.com/joho/godotenv"
)

func Init() {
	err := godotenv.Load()
	if err != nil {
		slog.Info("Error loading .env", err)
	}
}
