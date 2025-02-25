package main

import (
	"DeepSee_MAI/config"
	"DeepSee_MAI/pkg/logger"
	tele "gopkg.in/telebot.v3"
	"time"
)

func main() {

	lgr := logger.InitLogger()

	cfg, err := config.InitConfig()
	if err != nil {
		lgr.Err.Fatalf("Ошибка при инициализации конфигурации\n%v", err)
	}

	// настройка характеристик бота
	pref := tele.Settings{
		Token:     cfg.TelegramToken,
		Poller:    &tele.LongPoller{Timeout: 10 * time.Second},
		ParseMode: tele.ModeHTML,
	}

	// создание экземлпяра бота
	bot, err := tele.NewBot(pref)
	if err != nil {
		lgr.Err.Fatalf("Ошибка при запуске бота\n%v", err)
	}

	bot.Handle("/start", func(c tele.Context) error {
		return c.Send("Base implementation")
	})

	bot.Start()
}
