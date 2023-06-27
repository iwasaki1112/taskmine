package config

import (
	"errors"
	"log"
	"os"
	"taskmine/helper"

	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type Config struct {
	DBUser          string
	DBPassword      string
	DBHost          string
	DBName          string
	DBPort          string
	SlackWebHookURL string
}

func LoadEnvironmentVariables() error {
	path, err := helper.FindRoot()
	if err != nil {
		log.Printf("Not get root path")
	}

	err = godotenv.Load(path + "/.env")
	if err != nil {
		log.Printf("No .env file found: %s", path)
	}

	return err
}

func LoadGoogleOAuthConfig() (*oauth2.Config, error) {
	config := &oauth2.Config{
		ClientID:     os.Getenv("GOOGLE_OAUTH_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_OAUTH_CLIENT_SECRET"),
		Endpoint:     google.Endpoint,
		RedirectURL:  "http://localhost:8080/callback",
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email", "https://www.googleapis.com/auth/userinfo.profile"},
	}

	if config.ClientID == "" || config.ClientSecret == "" {
		return nil, errors.New("missing google clientID and ClientSecret")
	}

	return config, nil
}

func LoadDBConfig() (*Config, error) {
	path, err := helper.FindRoot()
	if err != nil {
		log.Printf("Not get root path")
	}

	err = godotenv.Load(path + "/.env")
	if err != nil {
		log.Printf("No .env file found: %s", path)
	}

	config := &Config{
		DBUser:          os.Getenv("DB_USER"),
		DBPassword:      os.Getenv("DB_PASSWORD"),
		DBHost:          os.Getenv("DB_HOST"),
		DBName:          os.Getenv("DB_NAME"),
		DBPort:          os.Getenv("DB_PORT"),
		SlackWebHookURL: os.Getenv("SLACK_WEBHOOK_URL"),
	}

	if config.SlackWebHookURL == "" || config.DBUser == "" || config.DBPassword == "" || config.DBHost == "" || config.DBName == "" || config.DBPort == "" {
		return nil, errors.New("missing required environment variable")
	}

	return config, nil

}
