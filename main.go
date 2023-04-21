package main

import "github.com/spf13/pflag"

var (
	_ = pflag.StringP("config", "c", "./config.yaml", "apiserver config path")
)

func main() {
	pflag.Parse()
}

// git remote set-url origin https://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/@github.com
// git remote set-url origin https://mymikasa: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@github.com/m1kasa5/feishu-bot.git
