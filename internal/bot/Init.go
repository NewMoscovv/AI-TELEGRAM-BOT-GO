package bot

import (
	"DeepSee_MAI/internal/config"
	tele "gopkg.in/telebot.v3"
	"time"
)

func InitBot(cfg *config.Config) (*tele.Bot, error) {

	// настройка характеристик бота
	pref := tele.Settings{
		Token:     cfg.TelegramToken,
		Poller:    &tele.LongPoller{Timeout: 10 * time.Second},
		ParseMode: tele.ModeMarkdown,
	}

	// создание экземлпяра бота
	bot, err := tele.NewBot(pref)
	if err != nil {
		return nil, err
	}

	return bot, nil
}
