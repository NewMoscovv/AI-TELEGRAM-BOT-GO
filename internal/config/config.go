package config

import (
	"DeepSee_MAI/pkg/consts"
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

	Prompt      string `mapstructure:"prompt"`
	BotMessages BotMessages
}

type BotMessages struct {
	Errors Errors
}

type Errors struct {
	SmthGoneWrong string `mapstructure:"smth_gone_wrong"`
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

	if err := viper.UnmarshalKey("messages.errors", &config.BotMessages.Errors); err != nil {
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
		return errors.New(consts.NoTelegramToken)
	}

	openRouterToken := os.Getenv("OPENROUTER_TOKEN")
	if openRouterToken == "" {
		return errors.New(consts.NoOpenRouterToken)
	}

	apiUrl := os.Getenv("API_URL")
	if apiUrl == "" {
		return errors.New(consts.NoUrl)
	}

	model := os.Getenv("MODEL")
	if model == "" {
		return errors.New(consts.NoModel)
	}

	cfg.TelegramToken = telegramToken
	cfg.OpnRtrToken = openRouterToken
	cfg.APIUrl = apiUrl
	cfg.Model = model

	return nil

}
