package handlers

import (
	"NutritionCalculator/data/models"
	contextkeys "NutritionCalculator/pkg/contextKeys"
	"NutritionCalculator/pkg/services/auth"
	"NutritionCalculator/pkg/services/hashing"
	"net/http"
)

func LoginHandler(authService auth.AuthService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			// Get the user request from the context
			userRequest, ok := r.Context().Value(contextkeys.UserRequestKey).(UserRequest)
			if !ok {
				http.Error(w, "Invalid form data", http.StatusBadRequest)
				return
			}
			hashingService := hashing.DefaultHashingService{}
			// Hash the password
			hashedPassword, err := hashingService.HashPassword(userRequest.Password)
			if err != nil {
				http.Error(w, "Invalid form data", http.StatusBadRequest)
				return
			}
			// Create a user object
			user := models.User{Username: userRequest.Username, PasswordHash: hashedPassword}

			// Authenticate the user
			authenticated, err := authService.Auth(user)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			if !authenticated {
				http.Error(w, "Invalid credentials", http.StatusUnauthorized)
				return
			}

			// If the user is authenticated, create a session for them

			w.WriteHeader(http.StatusOK)
			w.Write([]byte("Login successful"))
			return
		}

		w.Write([]byte("Login form"))
	}
}
