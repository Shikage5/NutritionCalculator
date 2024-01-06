package middleware

import (
	"NutritionCalculator/data/models"
	contextKeys "NutritionCalculator/pkg/contextKeys"
	"NutritionCalculator/pkg/services/validation"
	"context"
	"net/http"
)

// ValidateUser is a middleware that validates the username and password in the request body.
func ValidateUser(validator validation.ValidationService, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			// Decode the request body into a UserRequest struct
			err := r.ParseForm()
			if err != nil {
				http.Error(w, "Failed to parse form", http.StatusBadRequest)
				return
			}
			userRequest := models.UserRequest{
				Username: r.FormValue("username"),
				Password: r.FormValue("password"),
			}
			// Validate the user request
			err = validator.ValidateUserRequest(userRequest)
			if err != nil {
				http.Error(w, "Invalid User Input: "+err.Error(), http.StatusBadRequest)
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
