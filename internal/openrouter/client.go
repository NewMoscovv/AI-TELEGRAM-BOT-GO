package openrouter

import (
	"DeepSee_MAI/pkg/consts"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type ClientResponse interface {
	GetResponse(prompt string) (string, error)
}

type Client struct {
	APIKey      string
	APIUrl      string
	Model       string
	Prompt      string
	ChatHistory *ChatHistory
}

func NewClient(APIKey, APIUrl, Model string, prompt string, history *ChatHistory) *Client {
	return &Client{
		APIKey:      APIKey,
		APIUrl:      APIUrl,
		Model:       Model,
		Prompt:      prompt,
		ChatHistory: history,
	}
}

type ResponseBody struct {
	Choices []struct {
		Message Message `json:"message"`
	} `json:"choices"`
}

func (c *Client) GetResponse(prompt string, chatID int64) (string, error) {

	c.ChatHistory.AddMessage(chatID, "user", prompt)

	messages := c.ChatHistory.GetHistory(chatID)

	messages = append([]Message{
		{
			Role:    "system",
			Content: c.Prompt,
		},
	}, messages...)

	requestBody := struct {
		Model    string    `json:"model"`
		Messages []Message `json:"messages"`
	}{
		Model:    c.Model,
		Messages: messages,
	}

	body, err := json.Marshal(requestBody)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", c.APIUrl, bytes.NewBuffer(body))

	req.Header.Set("Authorization", "Bearer "+c.APIKey)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("%s: %v", consts.ResponseBodyError, err)
	}

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf(
			"%s: %s\nТело ответа: %s",
			consts.ApiRouterError,
			resp.Status,
			string(responseBody))
	}

	// Парсим JSON
	var response ResponseBody
	if err := json.Unmarshal(responseBody, &response); err != nil {
		return "", fmt.Errorf("%s: %v\nТело ответа: %s",
			consts.JSONParsingError,
			err,
			string(responseBody))
	}

	// Проверяем, что ответ содержит choices
	if len(response.Choices) == 0 {
		return "", fmt.Errorf("%s", consts.EmptyAnswerByAIError)
	}

	// Возвращаем ответ
	return response.Choices[0].Message.Content, nil
}
