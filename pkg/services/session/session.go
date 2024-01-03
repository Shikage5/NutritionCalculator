package session

import (
	"crypto/rand"
	"encoding/base64"
	"net/http"
	"time"
)

// SessionService defines the interface for session management.
type SessionService interface {
	CreateSession(username string, w http.ResponseWriter) error
	GenerateSessionID() (string, error)
	GetSession(sessionID string) (string, bool)
}

// InMemorySessionService implements the SessionService interface using an in-memory map.
type DefaultSessionService struct {
	SessionMap map[string]string
}

func (m *DefaultSessionService) CreateSession(username string, w http.ResponseWriter) error {
	// Generate a new session ID
	sessionID, err := m.GenerateSessionID()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return err
	}

	// Add the session to the session map
	m.SessionMap[sessionID] = username

	// Set the session duration
	sessionDuration := 1 * time.Minute

	// Set a cookie with the session ID
	http.SetCookie(w, &http.Cookie{
		Name:     "session_id",
		Value:    sessionID,
		MaxAge:   int(sessionDuration.Seconds()),
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteNoneMode,
	})

	// Start a goroutine to delete the session data when the session duration is over
	go func() {
		time.Sleep(sessionDuration)
		delete(m.SessionMap, sessionID)
	}()

	return nil
}

// GenerateSessionID generates a session ID for a user.
func (m *DefaultSessionService) GenerateSessionID() (string, error) {
	// Generate a random byte slice with 32 bytes
	randomBytes := make([]byte, 32)
	_, err := rand.Read(randomBytes)
	if err != nil {
		return "", err
	}

	// Encode the random bytes to base64 string
	sessionID := base64.URLEncoding.EncodeToString(randomBytes)

	return sessionID, nil
}

// GetSession returns the username for a given session ID.
func (m *DefaultSessionService) GetSession(sessionID string) (string, bool) {
	username, ok := m.SessionMap[sessionID]
	if !ok {
		return "", false
	}
	return username, true
}
