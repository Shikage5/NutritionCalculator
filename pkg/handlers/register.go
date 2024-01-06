package handlers

import (
	"NutritionCalculator/data/models"
	contextkeys "NutritionCalculator/pkg/contextKeys"
	"NutritionCalculator/pkg/services/registration"
	"net/http"
)

func RegisterHandler(registrationService registration.RegistrationService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			userRequest, ok := r.Context().Value(contextkeys.UserRequestKey).(models.UserRequest)
			if !ok {
				http.Error(w, "Registration failed", http.StatusBadRequest)
				return
			}
			err := registrationService.RegisterUser(userRequest)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			w.WriteHeader(http.StatusCreated)
			w.Write([]byte("Registration successful form"))
			return
		}

		DisplayPage(w, nil, "web/template/register.html")
	}
}
