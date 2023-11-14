package registrationHandlers

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	testCases := []struct {
		desc               string
		requestBody        string
		expectedStatusCode int
		expectedResponse   string
	}{
		{
			desc:               "Successful registration",
			requestBody:        "username=test&password=testpass",
			expectedStatusCode: http.StatusOK,
			expectedResponse:   "Registration successful",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			// Create a request with the specified form data
			req, err := http.NewRequest("POST", "/register", strings.NewReader(tC.requestBody))
			assert.NoError(t, err, "Error creating request")

			// Create a recorder to capture the response
			rr := httptest.NewRecorder()

			// Create an instance of MockRegistrationService for testing
			mockRegistrationService := &MockRegistrationService{}

			// Call the handler with the mock registration service
			RegisterPOSTHandler(rr, req, mockRegistrationService)

			// Assert the status code
			assert.Equal(t, tC.expectedStatusCode, rr.Code, "Unexpected status code")

			// Assert the response body
			assert.Equal(t, tC.expectedResponse, rr.Body.String(), "Unexpected response body")
		})
	}
}

// MockRegistrationService is a mock implementation of RegistrationService for testing.
type MockRegistrationService struct{}

// RegisterUser implements the mock registration logic.
func (s *MockRegistrationService) RegisterUser(username, password string) error {
	// Simulate successful registration
	return nil
}
