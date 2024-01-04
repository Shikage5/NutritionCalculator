package middleware

import (
	contextkeys "NutritionCalculator/pkg/contextKeys"
	"NutritionCalculator/pkg/services/session"
	"context"
	"net/http"
)

func SessionMiddleware(sessionService session.SessionService, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Get the session ID from the cookie
		cookie, err := r.Cookie("session_id")
		if err != nil {
			// If there's an error, redirect to the login page
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}

		sessionID := cookie.Value

		// Get the username from the session map
		username, ok := sessionService.GetSession(sessionID)
		if !ok {
			// If the session ID is not in the session map, redirect to the login page
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}

		// If the session is valid, add the user data to the request context
		ctx := context.WithValue(r.Context(), contextkeys.UserKey, username)

		// Call the next handler with the updated context
		next.ServeHTTP(w, r.WithContext(ctx))
	}
}
