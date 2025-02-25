package config

import (
	"errors"
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	TelegramToken string
}

func NewConfig() (*Config, error) {

	err := godotenv.Load(".env")
	if err != nil {
		return nil, err
	}

	telegramToken := os.Getenv("TELEGRAM_TOKEN")
	if telegramToken == "" {
		return nil, errors.New("отсутствует TELEGRAM_TOKEN")
	}

	return &Config{
		TelegramToken: telegramToken,
	}, nil

}
