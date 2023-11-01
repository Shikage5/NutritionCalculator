package handlers

import (
	"NutritionCalculator/pkg/models"
	"html/template"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		username := r.FormValue("username")
		password := r.FormValue("password")

		users, err := models.ReadUsersFromJSONFile("../users.json")
		if err != nil {
			// Handle error
			http.Error(w, "Failed to read user data.", http.StatusInternalServerError)
			return
		}
		var authenticated bool
		for _, user := range users {
			if user.Username == username {
				err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
				if err == nil {
					// Passwords match, user is authenticated
					authenticated = true
					// Here you can set a session or token to manage user sessions
					// Redirect the user to their profile or another page
					http.Redirect(w, r, "/profile", http.StatusSeeOther)
					return
				}
			}
		}

		// Authentication failed, display an error message
		if !authenticated {
			// Authentication failed, display an error message
			http.Error(w, "Authentication failed. Please check your username and password.", http.StatusUnauthorized)
			return
		}

	}

	// If the request method is not POST or authentication failed, show the login form
	tmpl, err := template.ParseFiles("pkg/templates/login.html")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, nil)
}
