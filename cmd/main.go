package main

import (
	"DeepSee_MAI/config"
	"DeepSee_MAI/pkg/logger"
)

func main() {

	lgr := logger.InitLogger()

	cfg, err := config.InitConfig()
	if err != nil {
		lgr.Err.Fatalf("Ошибка при инициализации конфигурации\n%v", err)
	}

	lgr.Info.Println(cfg)

}
