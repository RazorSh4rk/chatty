package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/urfave/cli/v2"
	fshelper "razorsh4rk.github.io/chatty/fs"
)

var validVerbs = []string{
	"init",
	"load",
	"new",
	"send",
	"list",
}

var App = &cli.App{
	Name:  "chatty",
	Usage: "a tty interface for chatgpt",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:     "name",
			Aliases:  []string{"n"},
			Category: "Chat selection",
			Value:    "sample",
			Usage:    "chat name, without the .json",
		},
		&cli.StringFlag{
			Name:     "message",
			Aliases:  []string{"msg", "m"},
			Category: "Chat content",
			Value:    "",
			Usage:    "\"your message\" or omit to use stdin",
		},
		&cli.StringFlag{
			Name:     "key",
			Category: "App configuration",
			Value:    "",
			Usage:    "set you openai api key with the set command",
		},
	},
	Commands: []*cli.Command{
		{
			Name: "set",
			Action: func(ctx *cli.Context) error {
				key := ctx.String("key")
				if key == "" {
					fmt.Println("Please give me a --key=\"\"")
					return nil
				}
				fshelper.SaveKey(key)
				return nil
			},
		},
		{
			Name:    "init",
			Aliases: []string{"i"},
			Usage:   "initialize the app",
			Action: func(*cli.Context) error {
				printMessage("Initializing chatty...")
				Init()
				printMessage("App set up with defaults, use [list] to view available chats and [load] to load an existing chat.")
				return nil
			},
		},
		{
			Name:    "load",
			Aliases: []string{"l"},
			Usage:   "load an existing chat into context from --name",
			Action: func(ctx *cli.Context) error {
				cname := ctx.String("name")
				if cname == "sample" {
					printMessage("No chat name provided, loading the sample")
				} else {
					printMessage(fmt.Sprintf("Loading chat: %s.json", cname))
				}
				load(cname)
				return nil
			},
		},
		{
			Name:    "new",
			Aliases: []string{"n"},
			Usage:   "create a new, empty chat with --name",
			Action: func(ctx *cli.Context) error {
				cname := ctx.String("name")
				if cname == "sample" {
					printMessage("Please provide a name for the new chat with --name")
					panic("")
				}
				fshelper.NewMemory(cname)
				fshelper.SaveMemory()

				return nil
			},
		},
		{
			Name:    "list",
			Aliases: []string{"ls"},
			Usage:   "list existing chats",
			Action: func(ctx *cli.Context) error {
				chats, err := fshelper.GetAllChats()
				if err != nil {
					printErr(err)
					return err
				}
				printMessage(fmt.Sprintf("Saved chats: [ %s ]", strings.Join(chats, ", ")))
				printMessage(fmt.Sprintf("Currently loaded chat: [ %s ]", fshelper.GetMemory().LoadedChatName))
				return nil
			},
		},
		{
			Name:    "send",
			Aliases: []string{"s", "talk", "say"},
			Usage:   "send a message, either with --message=\"hello world\" or interactively",
			Action: func(ctx *cli.Context) error {
				message := ctx.String("message")
				var messageString string

				if message == "" {
					printMessage("Type your message, double return sends it:")
					var msgs []string
					scanner := bufio.NewScanner(os.Stdin)
					for {
						scanner.Scan()
						line := scanner.Text()
						if len(line) == 0 {
							break
						}
						msgs = append(msgs, line)
					}
					messageString = strings.Join(msgs, "\n")
				} else {
					messageString = message
				}
				talk(messageString)

				return nil
			},
		},
	},
}
