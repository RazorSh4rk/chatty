package fs

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"github.com/sashabaranov/go-openai"
)

func LoadEnv() {
	path := getConfigFolder()
	err := godotenv.Load(path + "/.env")
	if err != nil {
		fmt.Printf("Error loading .env from %s: %s\n", path, err.Error())
		os.Create(path + "/.env")
	}
	if os.Getenv("OPENAI_KEY") == "" {
		s := fmt.Sprintf(`Please give me an openai key!
		Either set OPENAI_KEY environment variable,
		call 'chatty --key="..." set', 
		or set up your dotenv file in %s`, path)
		fmt.Println(s)
		os.Exit(1)
	}
}

func EncodeChat(messages []openai.ChatCompletionMessage) (string, error) {
	res, err := json.Marshal(messages)
	return string(res), err
}

func encodeChatPretty(messages []openai.ChatCompletionMessage) (string, error) {
	res, err := json.MarshalIndent(messages, "", "  ")
	return string(res), err
}

func DecodeChat(messages string) ([]openai.ChatCompletionMessage, error) {
	var box []openai.ChatCompletionMessage
	err := json.Unmarshal([]byte(messages), &box)
	return box, err
}

func EnsureConfigFolder() {
	path := getConfigFolder()

	if exists, err := exists(path); !exists || err != nil {
		err = os.Mkdir(path, os.ModePerm)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func SaveChat(messages []openai.ChatCompletionMessage, title string) (string, error) {
	res, err := EncodeChat(messages)
	if err != nil {
		return "", err
	}

	path := getConfigFolder() + "/" + title + ".json"

	err = os.WriteFile(path, []byte(res), os.ModePerm)
	if err != nil {
		return "", err
	}

	return path, nil
}

func LoadChat(name string) ([]openai.ChatCompletionMessage, error) {
	path := getConfigFolder() + "/" + name + ".json"
	content, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return DecodeChat(string(content))
}

func PPChat(name string) {
	chat, err := LoadChat(name)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(encodeChatPretty(chat))
}

func GetRawChat(name string) string {
	chat, err := LoadChat(name)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	chatString, _ := EncodeChat(chat)
	return chatString

}

func GetAllChats() ([]string, error) {
	path := getConfigFolder()

	files, err := os.ReadDir(path)

	if err != nil {
		return nil, err
	}

	var ret []string

	for _, file := range files {
		if !file.IsDir() && strings.HasSuffix(file.Name(), ".json") {
			ret = append(ret, strings.Replace(file.Name(), ".json", "", -1))
		}
	}

	return ret, nil
}

func getConfigFolder() string {
	path := os.Getenv("CHAT_CONFIG_FOLDER")
	var err error
	if path == "" {
		path, err = os.UserHomeDir()
		if err != nil {
			return "."
		}
	}
	path += "/.chatconfig"

	return path
}

func GetConfigFolder() func() string {
	return getConfigFolder
}

func SaveKey(key string) {
	path := getConfigFolder() + "/.env"
	content := fmt.Sprintf("OPENAI_KEY=%s\n", key)

	err := os.WriteFile(path, []byte(content), os.ModePerm)
	if err != nil {
		panic(err)
	}
}

func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
