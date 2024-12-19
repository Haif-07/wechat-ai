package chat

import (
	"context"

	"github.com/spf13/viper"
	"github.com/ysicing/openai/openai"
)

func RunChatAi(content, prompt string) (string, error) {

	provider := viper.GetString("openai.provider")
	token := viper.GetString("openai.token")
	model := viper.GetString("openai.model")
	// prompt := viper.GetString("openai.prompt")

	client, err := openai.New(
		openai.WithToken(token),
		openai.WithProvider(provider),
		openai.WithModel(model),
	)
	if err != nil {
		return "", err
	}

	resp, err := client.Completion(context.Background(), prompt, content)
	if err != nil {
		return "", err
	}
	return resp.Content, nil

}
