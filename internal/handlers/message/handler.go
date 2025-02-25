package message

import (
	"DeepSee_MAI/pkg/logger"
	tele "gopkg.in/telebot.v3"
)

type MsgHandler interface {
	HandleStart(c tele.Context) error
	HandleText(c tele.Context) error
}

type Handler struct {
	Bot *tele.Bot
	lgr *logger.Logger
}

func NewHandler(bot *tele.Bot, logger *logger.Logger) *Handler {
	return &Handler{
		Bot: bot,
		lgr: logger,
	}
}

func SetupHandlers(bot *tele.Bot, logger *logger.Logger) {
	var msgHandler MsgHandler

	msgHandler = NewHandler(bot, logger)

	bot.Handle("/start", msgHandler.HandleStart)
	bot.Handle(tele.OnText, msgHandler.HandleText)

}

func (h *Handler) HandleStart(c tele.Context) error {
	h.lgr.Info.Printf("%s | %s", c.Sender().Username, c.Text())

	return c.Send("Привет!")
}

func (h *Handler) HandleText(c tele.Context) error {
	h.lgr.Info.Printf("%s | %s", c.Sender().Username, c.Text())

	return c.Send("Привет?")
}
