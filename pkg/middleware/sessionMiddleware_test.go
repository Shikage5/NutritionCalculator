package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	contextkeys "NutritionCalculator/pkg/contextKeys"

	"github.com/stretchr/testify/assert"
)

func TestSessionMiddleware(t *testing.T) {
	testCases := []struct {
		desc           string
		sessionID      string
		shouldFail     bool
		expectedStatus int
	}{
		{
			desc:           "Valid Session",
			sessionID:      "validSessionID",
			shouldFail:     false,
			expectedStatus: http.StatusOK,
		},
		{
			desc:           "Invalid Session",
			sessionID:      "invalidSessionID",
			shouldFail:     true,
			expectedStatus: http.StatusUnauthorized,
		},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			// Setup
			req, err := http.NewRequest("GET", "/", nil)
			if err != nil {
				t.Fatalf("Failed to create request: %v", err)
			}
			req.AddCookie(&http.Cookie{Name: "session_id", Value: tC.sessionID})

			sessionService := &MockSessionService{shouldFail: tC.shouldFail}
			nextHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				value := r.Context().Value(contextkeys.UserKey)
				assert.NotNil(t, value)
			})

			handler := SessionMiddleware(sessionService, nextHandler)

			// Execute
			rr := httptest.NewRecorder()
			handler.ServeHTTP(rr, req)

			// Assert
			assert.Equal(t, tC.expectedStatus, rr.Code, "Unexpected status code")
		})
	}
}

type MockSessionService struct {
	shouldFail bool
}

func (m *MockSessionService) CreateSession(username string, w http.ResponseWriter) error {
	return nil
}

func (m *MockSessionService) GenerateSessionID() (string, error) {
	return "", nil
}

func (m *MockSessionService) GetSession(sessionID string) (string, bool) {
	if m.shouldFail {
		return "", false
	}

	return "validUsername", true
}
