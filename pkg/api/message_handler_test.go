package api_test

import (
	"hello-ai-world/pkg/api"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGenerateMessageHandler(t *testing.T) {
	t.Run("Test GenerateMessageHandler", func(t *testing.T) {
		testCases := []struct {
			name       string
			url        string
			statusCode int
			body       string
		}{
			{"GET with single name value", "/generateMessage?name=Bob", http.StatusOK, `{"message": "Hello, Bob"}`},
			{"GET with multiple name values, first one taken", "/generateMessage?name=Alice&arg=Bob", http.StatusOK, `{"message": "Hello, Alice"}`},
			{"GET without arguments", "/generateMessage", http.StatusOK, `{"message": "Hello, World"}`},
			{"GET with empty name", "/generateMessage?name=", http.StatusOK, `{"message": "Hello, World"}`},
			{"GET with argument that isn't name", "/generateMessage?arg=", http.StatusOK, `{"message": "Hello, World"}`},
		}

		for _, tc := range testCases {
			t.Run(tc.name, func(t *testing.T) {
				// Setup a new request to the handler
				req, _ := http.NewRequest("GET", tc.url, nil)

				// Create a ResponseRecorder which satisfies http.ResponseWriter to record the response.
				rr := httptest.NewRecorder()
				handler := http.HandlerFunc(api.GenerateMessageHandler)

				// Call the handler with the mock request and recorder
				handler.ServeHTTP(rr, req)

				// Check if the status code is as expected
				if status := rr.Code; status != tc.statusCode {
					t.Errorf("Handler returned wrong status code: got %v want %v",
						status, tc.statusCode)
				}

				// Check if the response body matches the expected one (if any)
				if rr.Body.String() != tc.body {
					t.Errorf("Handler returned unexpected body: got %v want %v",
						rr.Body.String(), tc.body)
				}
			})
		}
	})
}
