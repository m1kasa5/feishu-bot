package handlers

import (
	"feishu-bot/initialization"
	"feishu-bot/service/openai"
)

type MessageHandler struct {
	//sessionCache ser
}

func NewMessageHandler(gpt *openai.ChatGPT, config initialization.Config) MessageHandlerInterface {
	panic("implement me")
}
