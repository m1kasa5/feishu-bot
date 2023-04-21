package openai

import (
	"strings"

	"github.com/pandodao/tokenizer-go"
)

type AIMode float64

const (
	Fresh      AIMode = 0.1
	Warmth     AIMode = 0.4
	Balance    AIMode = 0.7
	Creativity AIMode = 1.0
)

var AIModeMap = map[string]AIMode{
	"清新": Fresh,
	"温暖": Warmth,
	"平衡": Balance,
	"创意": Creativity,
}

var AIModeStrs = []string{
	"清新",
	"温暖",
	"平衡",
	"创意",
}

const (
	maxTokens = 2000
	engine    = "gpt-3.5-turbo"
)

type Messages struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

// ChatGPTResponseBody 请求体
type ChatGPTResponseBody struct {
	ID      string                 `json:"id"`
	Object  string                 `json:"object"`
	Created int                    `json:"created"`
	Model   string                 `json:"model"`
	Choices []ChatGPTChoiceItem    `json:"choices"`
	Usage   map[string]interface{} `json:"usage"`
}

type ChatGPTChoiceItem struct {
	Message      Messages `json:"message"`
	Index        int      `json:"index"`
	FinishReason string   `json:"finish_reason"`
}

// ChatGPTRequestBody 响应体
type ChatGPTRequestBody struct {
	Model            string     `json:"model"`
	Messages         []Messages `json:"messages"`
	MaxTokens        int        `json:"max_tokens"`
	Temperature      AIMode     `json:"temperature"`
	TopP             int        `json:"top_p"`
	FrequencyPenalty int        `json:"frequency_penalty"`
	PresencePenalty  int        `json:"presence_penalty"`
}

func (msg *Messages) CalculateTokenLength() int {
	text := strings.TrimSpace(msg.Content)
	return tokenizer.MustCalToken(text)
}

func (msg *Messages) Completions() {

}
