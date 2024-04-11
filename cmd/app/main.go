package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	args := os.Args[1:]
	fmt.Println(generateMessage(args...))
}

func generateMessage(args ...string) string {
	for i := range args {
		if strings.TrimSpace(args[i]) == "" {
			continue
		} else {
			return fmt.Sprintf("Hello, %s", strings.Join(args[:], ", "))
		}
	}
	return "Hello, World"
}
