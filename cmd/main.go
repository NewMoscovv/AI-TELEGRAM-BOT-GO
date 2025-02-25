package main

import (
	b "DeepSee_MAI/internal/bot"
	"DeepSee_MAI/internal/config"
	"DeepSee_MAI/internal/handlers/message"
	"DeepSee_MAI/pkg/logger"
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

	lgr.Info.Printf("Бот запущен с именем @%s", bot.Me.Username)

	message.SetupHandlers(bot, lgr)

	bot.Start()
}
