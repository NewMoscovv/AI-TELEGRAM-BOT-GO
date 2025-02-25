package main

import (
	b "DeepSee_MAI/internal/bot"
	"DeepSee_MAI/internal/config"
	"DeepSee_MAI/pkg/logger"
	tele "gopkg.in/telebot.v3"
)

func main() {

	lgr := logger.InitLogger()

	cfg, err := config.InitConfig()
	if err != nil {
		lgr.Err.Fatalf("Ошибка инициализации конфигурации\n%v", err)
	}

	// инициализация бота
	bot, err := b.InitBot(cfg)
	if err != nil {
		lgr.Err.Fatalf("Ошибка запуска бота\n%v", err)
	}

	bot.Handle("/start", func(c tele.Context) error {
		return c.Send("Base implementation")
	})

	bot.Start()
}
