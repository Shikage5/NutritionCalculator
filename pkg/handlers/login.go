package handlers

import (
	"NutritionCalculator/data/models"
	contextkeys "NutritionCalculator/pkg/contextKeys"
	"NutritionCalculator/pkg/services/auth"
	"NutritionCalculator/pkg/services/session"
	"html/template"
	"net/http"
)

type LoginTemplateData struct {
	ErrMsg string
}

func LoginHandler(authService auth.AuthService, sessionService session.SessionService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var data LoginTemplateData
		if r.Method == http.MethodPost {
			// Get the user request from the context
			userRequest, ok := r.Context().Value(contextkeys.UserRequestKey).(models.UserRequest)
			if !ok {
				http.Error(w, "invalid form data", http.StatusBadRequest)
				return
			}

			// Authenticate the user
			authenticated, err := authService.Auth(userRequest)
			if err == auth.ErrInvalidCredentials {
				data.ErrMsg = err.Error()
				DisplayPage(w, data, "web/template/login.html")
				return
			} else if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			} else if !authenticated {
				data.ErrMsg = err.Error()
				DisplayPage(w, data, "web/template/login.html")
				return
			}
			// Create a session
			err = sessionService.CreateSession(userRequest.Username, w)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}

		//Display the login page template
		DisplayPage(w, data, "web/template/login.html")
		w.WriteHeader(http.StatusOK)
	}
}

func DisplayPage(w http.ResponseWriter, data any, templatePath string) {
	tmpl, err := template.ParseFiles(templatePath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
