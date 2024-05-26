package ai

import (
	"context"
	"os"

	"github.com/sashabaranov/go-openai"
	fshelper "razorsh4rk.github.io/chatty/fs"
)

var client *openai.Client

func Setup() {
	key := os.Getenv("OPENAI_KEY")
	if key == "" {
		return
	}
	client = openai.NewClient(key)
}

func Talk(message string) (string, error) {
	chat := fshelper.GetMemory()
	extended := append(chat.LoadedChat, openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleUser,
		Content: message,
	})

	res, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model:    openai.GPT3Dot5Turbo,
			Messages: extended,
		},
	)

	if err != nil {
		return "", err
	}

	return res.Choices[0].Message.Content, nil
}
