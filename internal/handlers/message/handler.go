package message

import (
	tele "gopkg.in/telebot.v3"
)

type MsgHandler interface {
	HandleStart(c tele.Context) error
	HandleText(c tele.Context) error
}

type Handler struct {
	Bot *tele.Bot
}

func NewHandler(bot *tele.Bot) *Handler {
	return &Handler{
		Bot: bot,
	}
}

func SetupHandlers(bot *tele.Bot) {
	// Todo: implement me
}

func (h *Handler) HandleStart(c tele.Context) error {
	return c.Send("Привет!")
}

func (h *Handler) HandleText(c tele.Context) error {
	return c.Send("Привет?")
}
