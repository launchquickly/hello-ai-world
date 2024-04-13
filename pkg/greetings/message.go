package greetings

import (
	"fmt"
	"strings"
)

type Greeting struct{}

func (g Greeting) GenerateMessage(args ...string) string {
	for i := range args {
		if strings.TrimSpace(args[i]) == "" {
			continue
		} else {
			return fmt.Sprintf("Hello, %s", strings.Join(args[:], ", "))
		}
	}
	return "Hello, World"
}
