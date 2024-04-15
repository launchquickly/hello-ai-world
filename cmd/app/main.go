package main

import (
	"fmt"
	"hello-ai-world/pkg/greetings"
	"os"
	"strings"
)

func main() {
	args := os.Args[1:]
	var message string = ""
	if len(args) > 0 {
		message = strings.TrimSpace(args[0]) // Trims whitespace from the first arg
	}
	fmt.Println(greetings.Greeting{}.GenerateMessage(message))
}
