package initialization

import (
	"fmt"
	"github.com/spf13/viper"
	"strconv"
)

type Config struct {
	FeishuAppId                string
	FeishuAppSecret            string
	FeishuAppEncryptKey        string
	FeishuAppVerificationToken string
	FeishuBotName              string
	OpenaiApiKeys              []string
	HttpPort                   int
	HttpsPort                  int
	UseHttps                   bool
	CertFile                   string
	KeyFile                    string
	OpenaiApiUrl               string
	HttpProxy                  string
	AzureOn                    bool
	AzureApiVersion            string
	AzureDeploymentName        string
	AzureResourceName          string
	AzureOpenaiToken           string
}

func LoadConfig(cfg string) *Config {
	viper.SetConfigFile(cfg)
	err := viper.ReadInConfig()
	if err != nil { // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
	viper.AutomaticEnv()
	config := &Config{
		FeishuAppId:                getViperStringValue("APP_ID", ""),
		FeishuAppSecret:            getViperStringValue("APP_SECRET", ""),
		FeishuAppEncryptKey:        getViperStringValue("APP_ENCRYPT_KEY", ""),
		FeishuAppVerificationToken: getViperStringValue("APP_VERIFICATION_TOKEN", ""),
		FeishuBotName:              getViperStringValue("BOT_NAME", ""),
		OpenaiApiKeys:              getViperStringArray("OPENAI_KEY", nil),
		HttpPort:                   getViperIntValue("HTTP_PORT", 9000),
		HttpsPort:                  getViperIntValue("HTTPS_PORT", 9001),
		UseHttps:                   getViperBoolValue("USE_HTTPS", false),
		CertFile:                   getViperStringValue("CERT_FILE", "cert.pem"),
		KeyFile:                    getViperStringValue("KEY_FILE", "key.pem"),
		OpenaiApiUrl:               getViperStringValue("API_URL", "https://api.openai.com"),
		HttpProxy:                  getViperStringValue("HTTP_PROXY", ""),
		AzureOn:                    getViperBoolValue("AZURE_ON", false),
		AzureApiVersion:            getViperStringValue("AZURE_API_VERSION", "2023-03-15-preview"),
		AzureDeploymentName:        getViperStringValue("AZURE_DEPLOYMENT_NAME", ""),
		AzureResourceName:          getViperStringValue("AZURE_RESOURCE_NAME", ""),
		AzureOpenaiToken:           getViperStringValue("AZURE_OPENAI_TOKEN", ""),
	}

	return config
}

func getViperBoolValue(key string, defaultValue bool) bool {
	value := viper.GetString(key)

	if value == "" {
		return defaultValue
	}

	boolValue, err := strconv.ParseBool(value)
	if err != nil {
		fmt.Printf("Invalid value for %s, using default value %v\n", key, defaultValue)
		return defaultValue
	}
	return boolValue
}

func getViperIntValue(key string, defaultValue int) int {
	value := viper.GetInt(key)

	if value == 0 {
		return defaultValue
	}
	return value
}

func getViperStringArray(key string, defaultValue []string) []string {
	value := viper.GetStringSlice(key)

	if value == nil {
		return defaultValue
	}
	return value
}

func getViperStringValue(key string, defaultValue string) string {
	value := viper.GetString(key)

	if value == "" {
		return defaultValue
	}
	return value
}
