package handlers

import (
	"NutritionCalculator/pkg/models"
	"encoding/json"
	"html/template"
	"net/http"
	"os"

	"golang.org/x/crypto/bcrypt"
)

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	// Handle user registration
	if r.Method == http.MethodPost {
		username := r.FormValue("username")
		password := r.FormValue("password")

		// Check if the username is already taken
		if isUsernameTaken("users.json", username) {
			// Handle the case where the username is already taken
			http.Error(w, "Username is already taken. Please choose a different one.", http.StatusConflict)
			return
		} else {
			newUser := models.User{
				Username: username,
			}

			// Hash the password before storing it in your JSON file
			hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
			if err != nil {
				// Handle error
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				return
			}

			newUser.PasswordHash = string(hashedPassword)

			// Append the new user to your JSON file
			if err := appendUserToJSONFile("users.json", newUser); err != nil {
				// Handle error
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				return
			}

			// Redirect the user to a login page, profile page, or any other appropriate location
			http.Redirect(w, r, "/login", http.StatusSeeOther)
		}
	}

	// If the request method is not POST or registration is unsuccessful, show the registration form
	tmpl, err := template.ParseFiles("pkg/templates/register.html")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, nil)
}

func appendUserToJSONFile(filename string, user models.User) error {
	// Read the existing user data from the JSON file
	users, err := models.ReadUsersFromJSONFile(filename)
	if err != nil {
		return err
	}

	// Append the new user to the slice of users
	users = append(users, user)

	// Encode the updated user data as JSON
	newData, err := json.Marshal(users)
	if err != nil {
		return err
	}

	// Write the updated data back to the JSON file
	if err := os.WriteFile(filename, newData, 0644); err != nil {
		return err
	}

	return nil
}

func isUsernameTaken(filename, username string) bool {
	// Read the existing user data from the JSON file
	users, err := models.ReadUsersFromJSONFile(filename)
	if err != nil {
		// Handle error, for example, by logging the error
		return true // Assume the username is taken due to the error
	}

	// Check if the username is already in use
	for _, user := range users {
		if user.Username == username {
			return true
		}
	}

	return false
}
