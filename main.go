package main

import (
	"os"
	"slices"

	"razorsh4rk.github.io/chatgpt-tui/ai"
	"razorsh4rk.github.io/chatgpt-tui/cmd"
	fshelper "razorsh4rk.github.io/chatgpt-tui/fs"
)

func main() {
	if !slices.Contains(os.Args, "set") {
		fshelper.LoadEnv()
		fshelper.EnsureConfigFolder()
		fshelper.EnsureMemory()
		fshelper.LoadMemory()
		ai.Setup()
	}

	cmd.App.Run(os.Args)
}
