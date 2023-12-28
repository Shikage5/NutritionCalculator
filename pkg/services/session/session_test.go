package session

import (
	"net/http/httptest"
	"testing"
)

func TestCreateSession(t *testing.T) {
	tests := []struct {
		name     string
		username string
		hasError bool
	}{
		{
			name:     "Valid username",
			username: "testuser",
			hasError: false,
		},
		{
			name:     "Empty username",
			username: "",
			hasError: false,
		},
	}

	SessionService := &DefaultSessionService{
		SessionMap: make(map[string]string),
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()

			err := SessionService.CreateSession(tt.username, w)
			if (err != nil) != tt.hasError {
				t.Errorf("CreateSession() error = %v, hasError %v", err, tt.hasError)
				return
			}

			// Check that the session was stored in the session map
			cookie := w.Result().Cookies()[0]
			sessionID := cookie.Value
			username, ok := SessionService.SessionMap[sessionID]
			if !ok {
				t.Errorf("Session not found in session map")
				return
			}
			if username != tt.username {
				t.Errorf("Expected username %s, got %s", tt.username, username)
			}
		})
	}
}
