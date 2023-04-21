package openai

import "feishu-bot/service/loadbalancer"

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
	Lb          *loadbalancer.LoadBalancer
	ApiKey      []string
	ApiUrl      string
	HttpProxy   string
	Platform    PlatForm
	AzureConfig AzureConfig
}

type requestBodyType int

const (
	jsonBody requestBodyType = iota
	formVoiceDataBody
	formPictureDataBody

	nilBody
)
