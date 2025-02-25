package config

import (
	"errors"
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	TelegramToken string
	OpnRtrToken   string
	APIUrl        string
	Model         string
}

func InitConfig() (*Config, error) {

	err := godotenv.Load(".env")
	if err != nil {
		return nil, err
	}

	telegramToken := os.Getenv("TELEGRAM_TOKEN")
	if telegramToken == "" {
		return nil, errors.New("отсутствует TELEGRAM_TOKEN")
	}

	openRouterToken := os.Getenv("OPENROUTER_TOKEN")
	if openRouterToken == "" {
		return nil, errors.New("отсутствует токен OpenRouter")
	}

	apiUrl := os.Getenv("API_URL")
	if apiUrl == "" {
		return nil, errors.New("отсутствует API_URL")
	}

	model := os.Getenv("MODEL")
	if model == "" {
		return nil, errors.New("отсутствует название Модели")
	}

	return &Config{
		TelegramToken: telegramToken,
		OpnRtrToken:   openRouterToken,
		APIUrl:        apiUrl,
		Model:         model,
	}, nil

}
