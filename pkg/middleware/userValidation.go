package middleware

import (
	contextKeys "NutritionCalculator/pkg/contextKeys"
	"NutritionCalculator/pkg/handlers"
	"NutritionCalculator/pkg/services/validation"
	"context"
	"encoding/json"
	"net/http"
)

// ValidateUser is a middleware that validates the username and password in the request body.
func ValidateUser(validator validation.CredentialsValidationService, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			var userRequest handlers.UserRequest

			err := json.NewDecoder(r.Body).Decode(&userRequest)
			if err != nil {
				http.Error(w, "Invalid request body", http.StatusBadRequest)
				return
			}
			if !validator.ValidateCredentials(userRequest.Username, userRequest.Password) {
				http.Error(w, "Username and password are required", http.StatusBadRequest)
				return
			}
			ctx := context.WithValue(r.Context(), contextKeys.UserRequestKey, userRequest)

			// Create a new request with the context
			r = r.WithContext(ctx)

			// Pass the new request to the next handler
			next(w, r)
		} else {
			// If the request method is not POST, pass it to the next handler
			next(w, r)
		}
	}
}
