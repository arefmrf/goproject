package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func LoadDBConfig() DBConfig {
	err := godotenv.Load() // Load .env file
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return DBConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Username: os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
		Database: os.Getenv("DB_DATABASE"),
	}
}
