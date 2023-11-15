package middleware

import (
	"NutritionCalculator/pkg/services/validation"
	"net/http"
)

// ValidateCredentialsMiddleware is a middleware that validates the username and password in the request body.
func ValidateUser(validator validation.CredentialsValidator, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		username := r.FormValue("username")
		password := r.FormValue("password")

		if !validator.ValidateCredentials(username, password) {
			http.Error(w, "Username and password are required", http.StatusBadRequest)
			return
		}

		// Call the next handler in the chain
		next.ServeHTTP(w, r)
	})
}
