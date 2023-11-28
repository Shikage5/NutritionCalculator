package middleware

import (
	"NutritionCalculator/pkg/services/validation"
	"encoding/json"
	"net/http"
)

type UserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// ValidateUser is a middleware that validates the username and password in the request body.
func ValidateUser(validator validation.CredentialsValidator, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var userRequest UserRequest

		err := json.NewDecoder(r.Body).Decode(&userRequest)
		if err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}
		if !validator.ValidateCredentials(userRequest.Username, userRequest.Password) {
			http.Error(w, "Username and password are required", http.StatusBadRequest)
			return
		}

		next(w, r)
	}
}
