package app

import (
	"bytes"
	"encoding/json"
	"net/http"
	"strings"
)

const Warning = 1
const Critical = 2
const Info = 3

func LogMessage(message string, messageType int, additionalData []string) {
	chatId := ""
	switch messageType {
	case Info:
	case Critical:
	case Warning:
		chatId = conf.Chat
		break
	}
	values := []string{message}
	for i := range additionalData {
		values = append(values, additionalData[i])
	}
	TelegramMessage(chatId, strings.Join(values, "\n"), "html")
}

func TelegramMessage(chatId string, message string, parseMode string) {
	requestBody, _ := json.Marshal(map[string]string{
		"chat_id":    chatId,
		"text":       message,
		"parse_mode": parseMode,
	})
	url := "https://api.telegram.org/bot" + conf.BotToken + "/sendMessage"
	_, _ = http.Post(url, "application/json", bytes.NewBuffer(requestBody))
}
