package bot

import "DeepSee_MAI/internal/openrouter"

type User struct {
	ChatHistory *ChatHistory
}

type ChatHistory struct {
	Messages []openrouter.Message
}

func NewUser() *User {
	messages := make([]openrouter.Message, 0)
	return &User{
		ChatHistory: &ChatHistory{messages},
	}
}
