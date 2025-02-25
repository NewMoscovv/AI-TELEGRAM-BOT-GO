package config

import (
	"errors"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"os"
)

type Config struct {
	TelegramToken string
	OpnRtrToken   string
	APIUrl        string
	Model         string

	Prompt string `mapstructure:"prompt"`
}

func InitConfig() (*Config, error) {
	viper.AddConfigPath("./")
	viper.SetConfigName("main")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	err := parseEnv(&config)
	if err != nil {
		return nil, err
	}

	return &config, nil

}

func parseEnv(cfg *Config) error {
	err := godotenv.Load(".env")
	if err != nil {
		return err
	}

	telegramToken := os.Getenv("TELEGRAM_TOKEN")
	if telegramToken == "" {
		return errors.New("отсутствует TELEGRAM_TOKEN")
	}

	openRouterToken := os.Getenv("OPENROUTER_TOKEN")
	if openRouterToken == "" {
		return errors.New("отсутствует токен OpenRouter")
	}

	apiUrl := os.Getenv("API_URL")
	if apiUrl == "" {
		return errors.New("отсутствует API_URL")
	}

	model := os.Getenv("MODEL")
	if model == "" {
		return errors.New("отсутствует название Модели")
	}

	cfg.TelegramToken = telegramToken
	cfg.OpnRtrToken = openRouterToken
	cfg.APIUrl = apiUrl
	cfg.Model = model

	return nil

}
