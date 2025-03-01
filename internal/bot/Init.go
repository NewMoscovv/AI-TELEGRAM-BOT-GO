package bot

import (
	"DeepSee_MAI/internal/config"
	"DeepSee_MAI/internal/handlers/message"
	"DeepSee_MAI/internal/openrouter"
	"DeepSee_MAI/pkg/consts"
	"DeepSee_MAI/pkg/logger"
	tele "gopkg.in/telebot.v3"
	"time"
)

type App struct {
	BotConfig *StructBot
	OpnRtr    *openrouter.Client
	Lgr       *logger.Logger
}

type StructBot struct {
	Bot            *tele.Bot
	SystemMessages SystemMessages
	msgHandler     message.MsgHandler
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

	botConfig := StructBot{
		Bot:            bot,
		SystemMessages: SystemMessages{cfg.BotMessages},
	}

	chatHistory := openrouter.NewChatHistory(consts.MaxFreeDialogLen)

	// подключение к openRouter
	opnRtr := openrouter.NewClient(cfg.OpnRtrToken, cfg.APIUrl, cfg.Model, cfg.Prompt, chatHistory)

	return &App{
		BotConfig: &botConfig,
		OpnRtr:    opnRtr,
		Lgr:       lgr,
	}, nil
}

func (app *App) Start() {

	app.Lgr.Info.Printf("Бот запущен с именем @%s", app.BotConfig.Bot.Me.Username)

	app.setupHandlers()

	app.BotConfig.Bot.Start()
}

func (app *App) setupHandlers() {

	app.BotConfig.msgHandler = app.newHandler()

	app.BotConfig.Bot.Handle("/start", app.BotConfig.msgHandler.HandleStart)
	app.BotConfig.Bot.Handle(tele.OnText, app.BotConfig.msgHandler.HandleText)

}

func (app *App) newHandler() *message.Handler {
	return &message.Handler{
		Bot:         app.BotConfig.Bot,
		OpnRtr:      app.OpnRtr,
		BotMessages: app.BotConfig.SystemMessages.BotMessages,
		Lgr:         app.Lgr,
	}
}
