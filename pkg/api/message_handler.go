package api

import (
	"fmt"
	"hello-ai-world/pkg/greetings"
	"net/http"
	"strings"
)

// MessageHandler is a type for handler functions that can handle messages
type MessageHandler func(message string) error

func GenerateMessageHandler(w http.ResponseWriter, r *http.Request) {
	// Extract arguments from request query parameters
	name := strings.TrimSpace(r.URL.Query().Get("name")) // Get name "name" and trim whitespaces
	if name == "" {
		name = "" // default value if "name" is not in the query
	}

	// Call your function and return its result as JSON
	fmt.Fprintf(w, `{"message": "%s"}`, greetings.Greeting{}.GenerateMessage(name))
}
