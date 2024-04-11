package main

import (
	"fmt"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	fmt.Println("Setting up...")
	result := m.Run() // Runs all tests in the suite.
	fmt.Println("Tearing down...")
	os.Exit(result) // Exit with exit code from running tests.
}

func TestHelloWorld(t *testing.T) {
	expectedMsg := "Hello, World"
	t.Run("Testing Hello World", func(t *testing.T) {
		msg := generateMessage()
		if msg != expectedMsg {
			t.Errorf("Expected '%s', but got '%s'", expectedMsg, msg)
		}
	})
}

func TestHelloName(t *testing.T) {
	expectedMsg := "Hello, Paul"
	t.Run("Testing Hello Name", func(t *testing.T) {
		msg := generateMessage("Paul")
		if msg != expectedMsg {
			t.Errorf("Expected '%s', but got '%s'", expectedMsg, msg)
		}
	})
}

func TestHelloEmptyName(t *testing.T) {
	expectedMsg := "Hello, World"
	t.Run("Testing Hello Empty Name", func(t *testing.T) {
		msg := generateMessage("")
		if msg != expectedMsg {
			t.Errorf("Expected '%s', but got '%s'", expectedMsg, msg)
		}
	})
}
