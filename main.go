package main

import (
	"feishu-bot/initialization"
	"fmt"
	"github.com/spf13/pflag"
)

var (
	cfg = pflag.StringP("config", "c", "./config.yaml", "apiserver config path")
)

func main() {
	pflag.Parse()
	config := initialization.LoadConfig(*cfg)
	fmt.Println(config.AzureApiVersion)
	initialization.LoadLarkClient(*config)
	//gpt := openai.NewChatGPT(*config)
	//handlers.
}

// git remote set-url origin https://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/@github.com
// git remote set-url origin https://mymikasa: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@github.com/m1kasa5/feishu-bot.git
