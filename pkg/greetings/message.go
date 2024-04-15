package greetings

import (
	"fmt"
	"strings"
)

type Greeting struct{}

func (g Greeting) GenerateMessage(name string) string {
	name = strings.TrimSpace(name)
	if name == "" {
		return "Hello, World"
	}
	return fmt.Sprintf("Hello, %s", name)
}
