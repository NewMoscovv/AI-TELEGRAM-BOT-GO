package main

import (
	b "DeepSee_MAI/internal/bot"
	"DeepSee_MAI/internal/config"
	"DeepSee_MAI/pkg/consts"
	"DeepSee_MAI/pkg/logger"
)

func main() {

	lgr := logger.InitLogger()

	cfg, err := config.InitConfig()
	if err != nil {
		lgr.Err.Fatalf("%s\n%v", consts.ConfigInitialisationError, err)
	}

	// инициализация бота
	app, err := b.InitApp(cfg, lgr)
	if err != nil {
		lgr.Err.Fatalf("%s\n%v", consts.BotStartPollingError, err)
	}

	app.Start()

}
