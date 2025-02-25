package message

import (
	"DeepSee_MAI/internal/config"
	"DeepSee_MAI/internal/openrouter"
	"DeepSee_MAI/pkg/logger"
	tele "gopkg.in/telebot.v3"
	"strings"
)

type MsgHandler interface {
	HandleStart(c tele.Context) error
	HandleText(c tele.Context) error
}

type Handler struct {
	Bot    *tele.Bot
	OpnRtr openrouter.ClientResponse
	lgr    *logger.Logger
}

func NewHandler(bot *tele.Bot, logger *logger.Logger, cfg *config.Config) *Handler {
	return &Handler{
		Bot:    bot,
		OpnRtr: openrouter.NewClient(cfg.OpnRtrToken, cfg.APIUrl, cfg.Model),
		lgr:    logger,
	}
}

func SetupHandlers(bot *tele.Bot, logger *logger.Logger, cfg *config.Config) {
	var msgHandler MsgHandler

	msgHandler = NewHandler(bot, logger, cfg)

	bot.Handle("/start", msgHandler.HandleStart)
	bot.Handle(tele.OnText, msgHandler.HandleText)

}

func (h *Handler) HandleStart(c tele.Context) error {
	h.lgr.Info.Printf("%s | %s", c.Sender().Username, c.Text())

	return c.Send("Привет!")
}

func (h *Handler) HandleText(c tele.Context) error {
	h.lgr.Info.Printf("%s | %s", c.Sender().Username, c.Text())

	for {
		// Печатает...
		c.Notify(tele.Typing)

		response, err := h.OpnRtr.GetResponse(c.Text())
		if err != nil {
			h.lgr.Err.Printf("%s", err.Error())
			return c.Send("Ой, что-то пошло не так. Обратитесь в поддержку - <b>@new_moscovv</b>")
		}
		if response == "" {
			h.lgr.Err.Printf("пустой ответ от ИИ, выполнение повторного запроса...")
		} else {
			h.lgr.Info.Printf("%s | \"%s\"", h.Bot.Me.Username, strings.Replace(response, "\n\n", "\n", -1))
			return c.Send(response)
		}
	}
}
