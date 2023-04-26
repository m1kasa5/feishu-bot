package handlers

import (
	"feishu-bot/initialization"
	"feishu-bot/service/openai"
)

type MessageHandlerInterface interface {
}

var handlers MessageHandlerInterface

func InitHandlers(gpt *openai.ChatGPT, config initialization.Config) {
	handlers = NewMessageHandler
}
