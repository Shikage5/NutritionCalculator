package handlers

import (
	"NutritionCalculator/data/models"
	contextkeys "NutritionCalculator/pkg/contextKeys"
	"NutritionCalculator/pkg/services/auth"
	"NutritionCalculator/pkg/services/session"
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
			// Create a user object
			user := models.User{Username: userRequest.Username, PasswordHash: userRequest.Password}

			// Authenticate the user
			authenticated, err := authService.Auth(user)
			if err == auth.ErrInvalidCredentials {
				http.Error(w, err.Error(), http.StatusUnauthorized)
				return
			} else if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			} else if !authenticated {
				http.Error(w, "invalid credentials", http.StatusUnauthorized)
				return
			}

			err = session.CreateSession(user.Username, w)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			w.WriteHeader(http.StatusOK)
			w.Write([]byte("Login successful"))
			return
		}

		w.Write([]byte("Login form"))
	}
}
