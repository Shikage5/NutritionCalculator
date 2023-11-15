package registrationHandlers

import (
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
		{"Valid GET Request", http.MethodGet, http.StatusOK},
		{"Valid POST Request", http.MethodPost, http.StatusOK},
		{"Invalid Method", http.MethodPut, http.StatusMethodNotAllowed},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			req, err := http.NewRequest(tc.method, "/test", nil)
			assert.NoError(t, err, "Failed to create request")

			recorder := httptest.NewRecorder()
			RegisterHandler(recorder, req)

			// Assert the status code
			assert.Equal(t, tc.expectedStatus, recorder.Code, "Unexpected status code")

			// You can add more assertions if needed based on your specific requirements.
		})
	}
}
