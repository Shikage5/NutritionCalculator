package handlers

import (
	"NutritionCalculator/pkg/services/registration"
	"fmt"
	"net/http"
)

func RegisterHandler(registrationService registration.RegistrationService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			username := r.FormValue("username")
			password := r.FormValue("password")
			err := registrationService.RegisterUser(username, password, "data/users.json")
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
