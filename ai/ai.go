package ai

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"strings"

	"github.com/sashabaranov/go-openai"
	fshelper "razorsh4rk.github.io/chatgpt-tui/fs"
)

var client *openai.Client

func Setup() {
	key := os.Getenv("OPENAI_KEY")
	if key == "" {
		return
	}
	client = openai.NewClient(key)
}

func GetChatTitle(messages string) string {
	res, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: "Return a single word title for this chat, return nothing else but the word: " + messages,
				},
			},
		},
	)

	if err != nil {
		return fmt.Sprint(rand.Int())
	}

	return strings.ToLower(res.Choices[0].Message.Content)
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
