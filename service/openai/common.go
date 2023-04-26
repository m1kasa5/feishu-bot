package openai

import (
	"feishu-bot/initialization"
	"feishu-bot/service/loadbalancer"
)

type PlatForm string

const AzureApiUrlV1 = "openai.azure.com/openai/deployments/"

const (
	OpenAI PlatForm = "openai"
	Azure  PlatForm = "azure"
)

type AzureConfig struct {
	BaseURL        string
	ResourceName   string
	DeploymentName string
	ApiVersion     string
	ApiToken       string
}

type ChatGPT struct {
	Lb        *loadbalancer.LoadBalancer
	ApiKey    []string
	ApiUrl    string
	HttpProxy string
	Platform  PlatForm

	AzureConfig AzureConfig
}

type requestBodyType int

const (
	jsonBody requestBodyType = iota
	formVoiceDataBody
	formPictureDataBody

	nilBody
)

func NewChatGPT(config initialization.Config) *ChatGPT {
	var lb *loadbalancer.LoadBalancer
	if config.AzureOn {
		keys := []string{config.AzureOpenaiToken}
		lb = loadbalancer.NewLoadBalancer(keys)
	} else {
		lb = loadbalancer.NewLoadBalancer(config.OpenaiApiKeys)
	}

	platform := OpenAI
	if config.AzureOn {
		platform = Azure
	}
	return &ChatGPT{
		Lb:        lb,
		ApiKey:    config.OpenaiApiKeys,
		ApiUrl:    config.OpenaiApiUrl,
		HttpProxy: config.HttpProxy,
		Platform:  platform,
		AzureConfig: AzureConfig{
			BaseURL:        AzureApiUrlV1,
			ResourceName:   config.AzureResourceName,
			DeploymentName: config.AzureDeploymentName,
			ApiVersion:     config.AzureApiVersion,
			ApiToken:       config.AzureOpenaiToken,
		},
	}
}
