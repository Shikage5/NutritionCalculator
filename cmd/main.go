package main

import (
	contextkeys "NutritionCalculator/pkg/contextKeys"
	"NutritionCalculator/pkg/handlers"
	"NutritionCalculator/pkg/middleware"
	"NutritionCalculator/pkg/services/auth"
	"NutritionCalculator/pkg/services/hashing"
	"NutritionCalculator/pkg/services/registration"
	"NutritionCalculator/pkg/services/session"
	"NutritionCalculator/pkg/services/validation"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

func greet(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("web/template/greet.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	user := r.Context().Value(contextkeys.UserKey)
	err = tmpl.Execute(w, user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	const userDataPath = "data/user_data/"
	const credentialsDataPath = "data/userCredentials.json"

	hashingService := &hashing.DefaultHashingService{}
	registrationService := &registration.DefaultRegistrationService{
		HashingService: hashingService,
		FilePath:       credentialsDataPath,
		UserDataPath:   userDataPath,
	}
	validationService := &validation.CredentialsValidationService{}
	authService := &auth.DefaultAuthService{
		FilePath: credentialsDataPath,
	}
	sessionService := &session.DefaultSessionService{
		SessionMap: make(map[string]string),
	}

	http.HandleFunc("/", middleware.SessionMiddleware(sessionService, greet))
	http.HandleFunc("/register", middleware.ValidateUser(validationService, handlers.RegisterHandler(registrationService)))
	http.HandleFunc("/login", middleware.ValidateUser(validationService, handlers.LoginHandler(authService, sessionService)))

	// Get the absolute path to the certificate file
	certFile, err := filepath.Abs("server.crt")
	if err != nil {
		log.Fatal("Failed to get absolute path for server.crt: ", err)
	}

	// Get the absolute path to the key file
	keyFile, err := filepath.Abs("server.key")
	if err != nil {
		log.Fatal("Failed to get absolute path for server.key: ", err)
	}

	err = http.ListenAndServeTLS(":443", certFile, keyFile, nil)
	if err != nil {
		log.Fatal("ListenAndServeTLS: ", err)
	}
}
