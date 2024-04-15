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
			{"GET with arguments", "/generateMessage?arg=Alice&arg=Bob", http.StatusOK, `{"message": "Hello, Alice"}`},
			{"GET without arguments", "/generateMessage", http.StatusOK, `{"message": "Hello, World"}`},
			{"GET with empty arguments", "/generateMessage?name=", http.StatusOK, `{"message": "Hello, World"}`},
		}

		for _, tc := range testCases {
			t.Run(tc.name, func(t *testing.T) {
				// Setup a new request to the handler
				var req *http.Request
				if tc.url == "/generateMessage?arg=Alice&arg=Bob" {
					req, _ = http.NewRequest("GET", tc.url, nil)
				} else {
					req, _ = http.NewRequest(http.MethodPost, tc.url, nil)
				}

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
