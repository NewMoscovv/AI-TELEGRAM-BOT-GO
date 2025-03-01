package openrouter

import (
	"sync"
)

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type ChatHistory struct {
	mu       sync.RWMutex
	sessions map[int64][]Message
	maxLen   int
}

func NewChatHistory(maxLen int) *ChatHistory {
	return &ChatHistory{
		sessions: make(map[int64][]Message),
		maxLen:   maxLen,
	}
}

func (h *ChatHistory) AddMessage(chatID int64, role string, content string) {
	h.mu.Lock()
	defer h.mu.Unlock()

	messages := append(h.sessions[chatID], Message{
		Role:    role,
		Content: content,
	})

	// Сохраняем только последние maxLen сообщений
	if len(messages) > h.maxLen {
		messages = messages[len(messages)-h.maxLen:]
	}

	h.sessions[chatID] = messages
}

// GetHistory возвращает историю сообщений для чата
func (h *ChatHistory) GetHistory(chatID int64) []Message {
	h.mu.RLock()
	defer h.mu.RUnlock()

	return h.sessions[chatID]
}

// ClearHistory очищает историю для чата
func (h *ChatHistory) ClearHistory(chatID int64) {
	h.mu.Lock()
	defer h.mu.Unlock()

	delete(h.sessions, chatID)
}
