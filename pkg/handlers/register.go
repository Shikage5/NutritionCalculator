package handlers

import (
	"NutritionCalculator/data/models"
	contextkeys "NutritionCalculator/pkg/contextKeys"
	"NutritionCalculator/pkg/services/registration"
	"NutritionCalculator/pkg/services/validation"
	"log"
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

			//validate userRequest

			validationService := &validation.DefaultValidationService{}
			err := validationService.ValidateUserRequest(userRequest)
			if err != nil {
				log.Println(err)
				data := PageData{
					error: err.Error(),
				}
				w.WriteHeader(http.StatusBadRequest)
				DisplayPage(w, data, "web/template/register.html")
				return
			}

			err = registrationService.RegisterUser(userRequest)
			if err != nil {
				log.Println(err)
				data := PageData{
					error: err.Error(),
				}
				w.WriteHeader(http.StatusBadRequest)
				DisplayPage(w, data, "web/template/register.html")
				return
			}

			w.WriteHeader(http.StatusCreated)
			w.Write([]byte("Registration successful form"))
			return
		}

		DisplayPage(w, nil, "web/template/register.html")
	}
}
