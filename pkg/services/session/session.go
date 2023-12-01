package session

import (
	"crypto/rand"
	"encoding/base64"
	"net/http"
	"time"
)

var SessionMap = make(map[string]string)

func CreateSession(username string, w http.ResponseWriter) error {
	// Generate a new session ID
	sessionID, err := GenerateSessionID()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return err
	}

	// Add the session to the session map
	SessionMap[sessionID] = username

	// Set the session duration
	sessionDuration := 1 * time.Minute

	// Set a cookie with the session ID
	http.SetCookie(w, &http.Cookie{
		Name:   "session_id",
		Value:  sessionID,
		MaxAge: int(sessionDuration.Seconds()),
	})

	// Start a goroutine to delete the session data when the session duration is over
	go func() {
		time.Sleep(sessionDuration)
		delete(SessionMap, sessionID)
	}()

	return nil
}

// GenerateSessionID generates a session ID for a user.
func GenerateSessionID() (string, error) {
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
