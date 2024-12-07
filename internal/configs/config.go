package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	DatabaseURL string
	JWTSecret   string
	SMTPConfig  struct {
		Host     string
		Port     int
		Username string
		Password string
	}
}

func LoadConfig(filePath string) (*Config, error) {
	if err := godotenv.Load(filePath); err != nil {
		return nil, err
	}

	var config Config

	config.DatabaseURL = os.Getenv("DATABASE_URL")
	config.JWTSecret = os.Getenv("JWT_SECRET")
	config.SMTPConfig.Host = os.Getenv("SMTP_HOST")
	config.SMTPConfig.Port, _ = strconv.Atoi(os.Getenv("SMTP_PORT"))
	config.SMTPConfig.Username = os.Getenv("SMTP_USERNAME")
	config.SMTPConfig.Password = os.Getenv("SMTP_PASSWORD")

	return &config, nil
}
