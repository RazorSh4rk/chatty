package cmd

import "fmt"

func printMessage[T any](msg T) {
	fmt.Println(msg)
}

func printErr(err error) {
	fmt.Println(err.Error())
}
