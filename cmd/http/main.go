package main

import (
	"hello-ai-world/pkg/api"
	"log"
	"net/http"
)

func main() {
	// Define new HTTP endpoint
	http.HandleFunc("/generateMessage", api.GenerateMessageHandler)

	// Run the server
	log.Fatal(http.ListenAndServe(":8080", nil))
}
