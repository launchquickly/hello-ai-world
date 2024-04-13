package greetings_test

import (
	"fmt"
	"hello-ai-world/pkg/greetings"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	fmt.Println("Setting up...")
	result := m.Run() // Runs all tests in the suite.
	fmt.Println("Tearing down...")
	os.Exit(result) // Exit with exit code from running tests.
}

func TestGenerateMessage(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected string
	}{
		{"Hello World", "", "Hello, World"},
		{"Hello Name", "Paul", "Hello, Paul"},
		{"Hello Empty Name", "", "Hello, World"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			msg := greetings.Greeting{}.GenerateMessage(tc.input)
			assert.Equal(t, tc.expected, msg)
		})
	}
}
