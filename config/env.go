package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Env        string
	Port       string
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
}

var AppConfig *Config

func Init() {
	if os.Getenv("ENV") == "" {
		err := godotenv.Load()
		if err != nil {
			panic("Error loading .env")
		}
	}

	AppConfig = &Config{
		Env:        os.Getenv("ENV"),
		Port:       os.Getenv("PORT"),
		DBHost:     os.Getenv("DB_HOST"),
		DBPort:     os.Getenv("DB_PORT"),
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBName:     os.Getenv("DB_NAME"),
	}
	validateConfig()
}

func validateConfig() {
	if AppConfig.Port == "" {
		panic("PORT is required")
	}

	if AppConfig.DBHost == "" {
		panic("DB_HOST is required")
	}

	if AppConfig.DBPort == "" {
		panic("DB_PORT is required")
	}

	if AppConfig.DBUser == "" {
		panic("DB_USER is required")
	}

	if AppConfig.DBPassword == "" {
		panic("DB_PASSWORD is required")
	}

	if AppConfig.DBName == "" {
		panic("DB_NAME is required")
	}

	if AppConfig.Env != "local" && AppConfig.Env != "production" {
		panic("ENV must be local or production")
	}

}
