package middleware

import (
	"NutritionCalculator/data/models"
	"NutritionCalculator/pkg/services/validation"
	"NutritionCalculator/utils"
	"net/http"
)

// ValidateUser is a middleware that validates the username and password in the request body.
func ValidateUser(validator validation.CredentialsValidator, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user models.User

		err := utils.DecodeJSON(r.Body, &user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if !validator.ValidateCredentials(user.Username, user.PasswordHash) {
			http.Error(w, "Username and password are required", http.StatusBadRequest)
			return
		}

		next(w, r)
	}
}
