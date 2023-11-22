package registrationHandlers

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRegisterPOSTHandler(t *testing.T) {
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
		{
			desc:               "Unsuccessful registration",
			requestBody:        "username=test&password=",
			expectedStatusCode: http.StatusInternalServerError,
			expectedResponse:   "Registration unsuccessful",
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
			RegisterPOSTHandler(mockRegistrationService).ServeHTTP(rr, req)

			// Assert the status code
			assert.Equal(t, tC.expectedStatusCode, rr.Code, "Unexpected status code")

			// Assert the response body
			assert.Equal(t, tC.expectedResponse, rr.Body.String(), "Unexpected response body")
		})
	}
}

type MockRegistrationService struct {
	shouldError bool
}

func (m *MockRegistrationService) RegisterUser(username, password, filename string) error {
	if m.shouldError {
		return errors.New("mock registration error")
	}
	return nil
}
