package fs_test

import (
	"os"
	"testing"

	"github.com/sashabaranov/go-openai"
	"razorsh4rk.github.io/chatty/fs"
)

func TestEncodeChat(t *testing.T) {
	testMessages := []openai.ChatCompletionMessage{
		{Role: openai.ChatMessageRoleUser, Content: "message"},
		{Role: openai.ChatMessageRoleSystem, Content: "message"},
	}
	expectedResult := `[{"role":"user","content":"message"},{"role":"system","content":"message"}]`

	result, err := fs.EncodeChat(testMessages)

	if err != nil {
		t.Errorf("EncodeChat() error = %v; want nil", err)
	}

	if result != expectedResult {
		t.Errorf("EncodeChat() = %v; want %v", result, expectedResult)
	}
}

func TestDecodeChat(t *testing.T) {
	testInput := `[{"role":"user","content":"message"},{"role":"system","content":"message"}]`
	expectedMessages := []openai.ChatCompletionMessage{
		{Role: openai.ChatMessageRoleUser, Content: "message"},
		{Role: openai.ChatMessageRoleSystem, Content: "message"},
	}

	result, err := fs.DecodeChat(testInput)

	if err != nil {
		t.Errorf("DecodeChat() error = %v; want nil", err)
	}

	for i := range expectedMessages {
		if result[i].Role != expectedMessages[i].Role && result[i].Content != expectedMessages[i].Content {
			t.Errorf("DecodeChat() message %v = %v; want %v", i, result[i], expectedMessages[i])
		}
	}
}

func TestSaveAndLoadChat(t *testing.T) {
	// Prepare sample chat messages
	messages := []openai.ChatCompletionMessage{
		{Role: openai.ChatMessageRoleUser, Content: "message"},
		{Role: openai.ChatMessageRoleSystem, Content: "message"},
	}

	// Save the chat messages
	fileName := "test_chat"
	_, saveErr := fs.SaveChat(messages, fileName)
	if saveErr != nil {
		t.Errorf("SaveChat() error = %v; want nil", saveErr)
	}

	// Load the chat messages
	loadedMessages, loadErr := fs.LoadChat(fileName)
	if loadErr != nil {
		t.Errorf("LoadChat() error = %v; want nil", loadErr)
	}

	// Compare the loaded messages with the original messages
	if len(loadedMessages) != len(messages) {
		t.Errorf("Loaded chat messages count = %d; want %d", len(loadedMessages), len(messages))
	}

	for i := range messages {
		if loadedMessages[i].Content != messages[i].Content {
			t.Errorf("Loaded chat message %d = %v; want %v", i, loadedMessages[i], messages[i])
		}
	}

	// Clean up by deleting the temporary test file
	err := os.Remove(fs.GetConfigFolder()() + "/" + fileName + ".json")
	if err != nil {
		t.Errorf("Error cleaning up test file: %v", err)
	}
}
