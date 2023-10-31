package handlers

import (
	"NutritionCalculator/models"
	"html/template"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		username := r.FormValue("username")
		password := r.FormValue("password")

		users, err := models.ReadUsersFromJSONFile("users.json")
		if err != nil {
			// Handle error
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		for _, user := range users {
			if user.Username == username {
				err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
				if err == nil {
					// Passwords match, user is authenticated
					// Here you can set a session or token to manage user sessions
					// Redirect the user to their profile or another page
					http.Redirect(w, r, "/profile", http.StatusSeeOther)
					return
				}
			}
		}

		// Authentication failed, display an error message
	}

	// If the request method is not POST or authentication failed, show the login form
	tmpl, err := template.ParseFiles("templates/login.html")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, nil)
}
