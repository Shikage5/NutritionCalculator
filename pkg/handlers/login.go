package handlers

import (
	"NutritionCalculator/data/models"
	contextkeys "NutritionCalculator/pkg/contextKeys"
	"NutritionCalculator/pkg/services/auth"
	"NutritionCalculator/pkg/services/session"
	"net/http"
)

func LoginHandler(authService auth.AuthService, sessionService session.SessionService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			// Get the user request from the context
			userRequest, ok := r.Context().Value(contextkeys.UserRequestKey).(UserRequest)
			if !ok {
				http.Error(w, "invalid form data", http.StatusBadRequest)
				return
			}
			// Create a userCredentials object
			userCredentials := models.UserCredentials{Username: userRequest.Username, PasswordHash: userRequest.Password}

			// Authenticate the user
			authenticated, err := authService.Auth(userCredentials)
			if err == auth.ErrInvalidCredentials {
				http.Error(w, err.Error(), http.StatusUnauthorized)
				return
			} else if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			} else if !authenticated {
				http.Error(w, auth.ErrInvalidCredentials.Error(), http.StatusUnauthorized)
				return
			}
			// Create a session
			err = sessionService.CreateSession(userCredentials.Username, w)
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
