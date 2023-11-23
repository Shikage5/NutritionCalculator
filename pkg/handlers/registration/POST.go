package registrationHandlers

import (
	"NutritionCalculator/pkg/services/registration"
	"net/http"
)

// RegisterPOSTHandler handles the POST request for user registration.
func RegisterPOSTHandler(registrationService registration.RegistrationService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Parse the form data from the HTTP request
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Failed to parse form data", http.StatusBadRequest)
			return
		}

		// Extract user input from the form
		username := r.FormValue("username")
		password := r.FormValue("password")
		filename := "data/users.json"

		// Call the function to register the user
		if err := registrationService.RegisterUser(username, password, filename); err != nil {
			http.Error(w, "Failed to register user: "+err.Error(), http.StatusInternalServerError)
			return
		}

		// Registration successful
		w.Write([]byte("Registration successful"))
	}
}
