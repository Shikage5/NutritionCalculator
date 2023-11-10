package registration

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRegisterHandler(t *testing.T) {
	testCases := []struct {
		name           string
		method         string
		expectedStatus int
	}{
		{"GET request", http.MethodGet, http.StatusOK},
		{"POST request", http.MethodPost, http.StatusOK},
		{"Unsupported method", http.MethodPut, http.StatusMethodNotAllowed},
	}

	// Create a test server
	server := httptest.NewServer(http.HandlerFunc(RegisterHandler))
	defer server.Close()

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var resp *http.Response
			var err error

			switch tc.method {
			case http.MethodGet:
				resp, err = http.Get(server.URL)
			case http.MethodPost:
				// Provide a request body for the POST request
				req, _ := http.NewRequest(tc.method, server.URL, bytes.NewBuffer([]byte("some content")))
				resp, err = http.DefaultClient.Do(req)
			default:
				req, _ := http.NewRequest(tc.method, server.URL, nil)
				resp, err = http.DefaultClient.Do(req)
			}

			assert.NoError(t, err)
			assert.Equal(t, tc.expectedStatus, resp.StatusCode)
		})
	}
}
