package bot

import (
	"DeepSee_MAI/internal/config"
	"DeepSee_MAI/internal/handlers/message"
	"DeepSee_MAI/internal/openrouter"
	"DeepSee_MAI/pkg/logger"
	tele "gopkg.in/telebot.v3"
	"time"
)

type App struct {
	Bot            *tele.Bot
	OpnRtr         *openrouter.Client
	msgHandler     message.MsgHandler
	Lgr            *logger.Logger
	SystemMessages SystemMessages
}

type SystemMessages struct {
	BotMessages config.BotMessages
}

func InitApp(cfg *config.Config, lgr *logger.Logger) (*App, error) {

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

	// подключение к openRouter
	opnRtr := openrouter.NewClient(cfg.OpnRtrToken, cfg.APIUrl, cfg.Model, cfg.Prompt)

	return &App{
		Bot:            bot,
		OpnRtr:         opnRtr,
		Lgr:            lgr,
		SystemMessages: SystemMessages{cfg.BotMessages},
	}, nil
}

func (app *App) Start() {

	app.Lgr.Info.Printf("Бот запущен с именем @%s", app.Bot.Me.Username)

	app.setupHandlers()

	app.Bot.Start()
}

func (app *App) setupHandlers() {

	app.msgHandler = app.newHandler()

	app.Bot.Handle("/start", app.msgHandler.HandleStart)
	app.Bot.Handle(tele.OnText, app.msgHandler.HandleText)

}

func (app *App) newHandler() *message.Handler {
	return &message.Handler{
		Bot:         app.Bot,
		OpnRtr:      app.OpnRtr,
		BotMessages: app.SystemMessages.BotMessages,
		Lgr:         app.Lgr,
	}
}
