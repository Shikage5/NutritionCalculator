package handlers

import (
	"NutritionCalculator/pkg/services/registration"
	"encoding/json"
	"fmt"
	"net/http"
)

type UserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func RegisterHandler(registrationService registration.RegistrationService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {

			var userRequest UserRequest
			err := json.NewDecoder(r.Body).Decode(&userRequest)
			if err != nil {
				http.Error(w, "Invalid request body", http.StatusBadRequest)
				return
			}
			err = registrationService.RegisterUser(userRequest.Username, userRequest.Password)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			w.WriteHeader(http.StatusCreated)
			w.Write([]byte("Registration successful form"))
			return
		}

		fmt.Fprintf(w, "Registration form")
	}
}
