package fs

import (
	"encoding/json"
	"errors"
	"os"

	"github.com/sashabaranov/go-openai"
)

type memory struct {
	LoadedChatName string                         `json:"loadedChatName"`
	LoadedChat     []openai.ChatCompletionMessage `json:"loadedChat"`
}

var (
	mem        memory
	sampleChat = []openai.ChatCompletionMessage{
		{
			Role:    openai.ChatMessageRoleUser,
			Content: "Hello!",
		},
		{
			Role:    openai.ChatMessageRoleSystem,
			Content: "Hello!",
		},
	}
)

func EnsureMemory() {
	path := getConfigFolder() + "/memory"
	if e, _ := exists(path); !e {
		_, err := os.Create(path)
		if err != nil {
			panic(err)
		}
	}
}

func InitMemory() {
	path := getConfigFolder() + "/memory"
	content, err := os.ReadFile(path)
	if err != nil || len(content) == 0 {
		os.Create(path)
	}
	mem.LoadedChat = sampleChat
	mem.LoadedChatName = "sample"
	SaveChat(mem.LoadedChat, mem.LoadedChatName)
}

func LoadMemory() error {
	path := getConfigFolder() + "/memory"
	content, err := os.ReadFile(path)
	if err != nil {
		EnsureMemory()
		return err
	}
	if len(content) == 0 {
		return errors.New("no saved session")
	}
	json.Unmarshal(content, &mem)

	return nil
}

func SaveMemory() error {
	path := getConfigFolder() + "/memory"
	res, err := json.Marshal(mem)
	if err != nil {
		return err
	}
	err = os.WriteFile(path, res, os.ModePerm)
	if err != nil {
		return err
	}

	_, err = SaveChat(mem.LoadedChat, mem.LoadedChatName)
	if err != nil {
		return err
	}

	return nil
}

func UpdateMemory(chat []openai.ChatCompletionMessage) {
	mem.LoadedChat = chat
}

func SetMemoryFromFile(name string) error {
	messages, err := LoadChat(name)
	if err != nil {
		return err
	}
	SetMemory(name, messages)
	err = SaveMemory()
	if err != nil {
		return err
	}
	return nil
}

func NewMemory(name string) {
	mem.LoadedChatName = name
	mem.LoadedChat = []openai.ChatCompletionMessage{}
}

func SetMemory(name string, chat []openai.ChatCompletionMessage) {
	mem.LoadedChatName = name
	mem.LoadedChat = chat
}

func GetMemory() memory {
	return mem
}
