package cmd

import (
	"github.com/sashabaranov/go-openai"
	"razorsh4rk.github.io/chatgpt-tui/ai"
	fshelper "razorsh4rk.github.io/chatgpt-tui/fs"
)

func talk(message string) {
	resp, err := ai.Talk(message)
	if err != nil {
		printErr(err)
	} else {
		printMessage(resp)
	}

	mem := fshelper.GetMemory().LoadedChat
	mem = append(mem, openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleUser,
		Content: message,
	})
	mem = append(mem, openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleSystem,
		Content: resp,
	})

	fshelper.UpdateMemory(mem)
	fshelper.SaveMemory()
}

func load(name string) {
	err := fshelper.SetMemoryFromFile(name)
	if err != nil {
		printErr(err)
	}
}

func Init() {
	fshelper.EnsureConfigFolder()
	fshelper.EnsureMemory()
	fshelper.InitMemory()
	fshelper.SaveMemory()
}
