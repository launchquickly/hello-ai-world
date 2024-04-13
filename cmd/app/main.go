package main

import (
	"fmt"
	"hello-ai-world/pkg/greetings"
	"os"
)

func main() {
	args := os.Args[1:]
	fmt.Println(greetings.Greeting{}.GenerateMessage(args...))
}
