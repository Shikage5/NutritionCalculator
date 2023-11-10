package registration

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRegisterPOSTHandler(t *testing.T) {
	// Create a test server
	server := httptest.NewServer(http.HandlerFunc(RegisterPOSTHandler))
	defer server.Close()

	// Test case: valid POST request
	validRequest := `{"username": "testuser", "password": "testpassword"}`
	resp, err := http.Post(server.URL, "application/json", bytes.NewBuffer([]byte(validRequest)))
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	// Test case: invalid POST request (missing username)
	invalidRequest := `{"password": "testpassword"}`
	resp, err = http.Post(server.URL, "application/json", bytes.NewBuffer([]byte(invalidRequest)))
	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)

	// Add more test cases as needed
}
