package config

import (
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	TelegramToken string
}

func NewConfig() *Config {

	err := godotenv.Load(".env")
	if err != nil {
		// TODO: log info via logger...
	}

	telegramToken := os.Getenv("TELEGRAM_TOKEN")
	if telegramToken == "" {
		// TODO: log info via logger...
	}

	return &Config{
		TelegramToken: telegramToken,
	}

}
