package main

import (
	"NutritionCalculator/pkg/handlers"
	"NutritionCalculator/pkg/middleware"
	"NutritionCalculator/pkg/services/auth"
	"NutritionCalculator/pkg/services/hashing"
	"NutritionCalculator/pkg/services/registration"
	"NutritionCalculator/pkg/services/validation"
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"time"
)

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World! %s", time.Now())
}

func main() {
	dataFilePath := "data/users.json"

	hashingService := &hashing.DefaultHashingService{}
	registrationService := &registration.DefaultRegistrationService{
		HashingService: hashingService,
		DataFilePath:   dataFilePath,
	}
	validationService := &validation.DefaultCredentialsValidationService{}
	authService := &auth.DefaultAuthService{
		FilePath: dataFilePath,
	}

	http.HandleFunc("/", greet)
	http.HandleFunc("/register", middleware.ValidateUser(validationService, handlers.RegisterHandler(registrationService)))
	http.HandleFunc("/login", middleware.ValidateUser(validationService, handlers.LoginHandler(authService)))

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
