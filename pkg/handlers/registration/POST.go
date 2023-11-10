package registration

import (
	"NutritionCalculator/data/models"
	"NutritionCalculator/pkg/services/hashing"
	"net/http"
)

func RegisterPOSTHandler(w http.ResponseWriter, r *http.Request) {
	// Parse the form data from the HTTP request
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Failed to parse form data", http.StatusBadRequest)
		return
	}

	// Extract user input from the form
	username := r.FormValue("username")
	hashedPassword, err := hashing.HashPassword(r.FormValue("password"))
	if err != nil {
		http.Error(w, "Failed to hash the password", http.StatusInternalServerError)
		return
	}
	user := models.User{
		Username:     username,
		PasswordHash: hashedPassword,
	}

	// Write the user to the "data/users.json" file
	err = models.WriteUserInJSONFile(user, "data/users.json")
	if err != nil {
		http.Error(w, "Failed to register user: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Registration successful
	w.Write([]byte("Registration successful"))
}
