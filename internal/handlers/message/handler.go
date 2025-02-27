package message

import (
	"DeepSee_MAI/internal/config"
	"DeepSee_MAI/internal/openrouter"
	"DeepSee_MAI/pkg/consts"
	"DeepSee_MAI/pkg/logger"
	tele "gopkg.in/telebot.v3"
	"strings"
	"time"
)

type MsgHandler interface {
	HandleStart(c tele.Context) error
	HandleText(c tele.Context) error
}

type Handler struct {
	Bot      *tele.Bot
	OpnRtr   *openrouter.Client
	Lgr      *logger.Logger
	Messages config.Messages
}

func (h *Handler) HandleStart(c tele.Context) error {
	h.Lgr.Info.Printf("%s | %s", c.Sender().Username, c.Text())
	h.Lgr.Info.Printf("%s | %s", h.Bot.Me.Username, c.Text())

	return c.Send(h.Messages.Errors.SmthGoneWrong,
		&tele.SendOptions{
			ParseMode: tele.ModeHTML,
		})
}

func (h *Handler) HandleText(c tele.Context) error {
	h.Lgr.Info.Printf("%s | %s", c.Sender().Username, c.Text())

	for i := 0; i < consts.MaxAmountResponses; i++ {
		// Печатает...
		err := c.Notify(tele.Typing)
		if err != nil {
			h.Lgr.Err.Printf("%s\n%s", consts.TypingAnimationError, err.Error())
		}

		response, err := h.OpnRtr.GetResponse(c.Text())
		if err != nil {
			h.Lgr.Err.Printf("%s", err.Error())
			return c.Send(h.Messages.Errors.SmthGoneWrong)
		}
		if response == "" {
			h.Lgr.Err.Printf("пустой ответ от ИИ, выполнение повторного запроса...")
		} else {
			h.Lgr.Info.Printf("%s | \"%s\"", h.Bot.Me.Username, strings.Replace(response, "\n\n", "\n", -1))
			return c.Send(response)
		}
		time.Sleep(1 * time.Second)
	}

	return c.Send("Произошла ошибка при получении ответа от ИИ. Пожалуйста, повторите запрос или обратитесь в поддержку")

}
