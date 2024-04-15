package api

import (
	"fmt"
	"hello-ai-world/pkg/greetings"
	"net/http"
)

// MessageHandler is a type for handler functions that can handle messages
type MessageHandler func(message string) error

func GenerateMessageHandler(w http.ResponseWriter, r *http.Request) {
	// Extract arguments from request query parameters
	var args []string
	for k, _ := range r.URL.Query() {
		if v := r.FormValue(k); v != "" {
			args = append(args, v)
		}
	}

	// Call your function and return its result as JSON
	fmt.Fprintf(w, `{"message": "%s"}`, greetings.Greeting{}.GenerateMessage(args...))
}
