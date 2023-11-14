package registrationHandlers

import (
	"NutritionCalculator/pkg/services/registration"
	"net/http"
)

func RegisterPOSTHandler(w http.ResponseWriter, r *http.Request, registrationService registration.RegistrationService) {
	// Parse the form data from the HTTP request
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Failed to parse form data", http.StatusBadRequest)
		return
	}

	// Extract user input from the form
	username := r.FormValue("username")
	password := r.FormValue("password")

	// Call the function to register the user
	if err := registrationService.RegisterUser(username, password); err != nil {
		http.Error(w, "Failed to register user: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Registration successful
	w.Write([]byte("Registration successful"))
}
