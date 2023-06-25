package config

import (
	"errors"
	"log"
	"os"
	"taskmine/helper"

	"github.com/joho/godotenv"
)

type Config struct {
	DBUser     string
	DBPassword string
	DBHost     string
	DBName     string
	DBPort     string
}

func LoadConfig() (*Config, error) {
	path, err := helper.FindRoot()
	if err != nil {
		log.Printf("Not get root path")
	}

	err = godotenv.Load(path + "/.env")
	if err != nil {
		log.Printf("No .env file found: %s", path)
	}

	config := &Config{
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBHost:     os.Getenv("DB_HOST"),
		DBName:     os.Getenv("DB_NAME"),
		DBPort:     os.Getenv("DB_PORT"),
	}

	if config.DBUser == "" || config.DBPassword == "" || config.DBHost == "" || config.DBName == "" || config.DBPort == "" {
		return nil, errors.New("missing required environment variable")
	}

	return config, nil

}
