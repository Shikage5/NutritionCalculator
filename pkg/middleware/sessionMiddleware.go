package middleware

import (
	"NutritionCalculator/pkg/services/session"
	"context"
	"net/http"
)

func sessionMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get the session ID from the cookie
		cookie, err := r.Cookie("session_id")
		if err != nil {
			// If there's an error, redirect to the login page
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}

		sessionID := cookie.Value

		// Get the user data from the session map
		userData, ok := session.SessionMap[sessionID]

		if !ok {
			// If the session ID is not in the session map, redirect to the login page
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}

		// If the session is valid, add the user data to the request context
		ctx := context.WithValue(r.Context(), "userData", userData)

		// Call the next handler with the updated context
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
