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
	"flag"
	"log"
	"net/http"
	"path/filepath"
)

func greet(w http.ResponseWriter, r *http.Request) {

	user := r.Context().Value(contextkeys.UserKey)
	handlers.DisplayPage(w, user, "web/template/greet.html")
}

func main() {
	port := flag.String("port", "8080", "port to run the server on")
	userDataPath := "data/user_data/"
	credentialsDataPath := "data/userCredentials.json"
	flag.Parse()

	s := startServices(userDataPath, credentialsDataPath)
	http.HandleFunc("/", middleware.SessionMiddleware(s.SessionService, greet))
	http.HandleFunc("/register", middleware.ValidateUser(s.ValidationService, handlers.RegisterHandler(s.RegistrationService)))
	http.HandleFunc("/login", middleware.ValidateUser(s.ValidationService, handlers.LoginHandler(s.AuthService, s.SessionService)))
	http.HandleFunc("/food", middleware.SessionMiddleware(s.SessionService, handlers.FoodHandler(userDataPath)))
	http.HandleFunc("/dish", middleware.SessionMiddleware(s.SessionService, handlers.DishHandler(userDataPath)))
	//ignore this
	http.HandleFunc("/testUserData", middleware.SessionMiddleware(s.SessionService, handlers.TestUserData(userDataPath)))

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

	log.Println("Starting server on port:", *port)
	err = http.ListenAndServeTLS(":"+*port, certFile, keyFile, nil)
	if err != nil {
		log.Fatal("ListenAndServeTLS: ", err)
	}
}

type Services struct {
	HashingService      *hashing.DefaultHashingService
	RegistrationService *registration.DefaultRegistrationService
	ValidationService   *validation.DefaultValidationService
	AuthService         *auth.DefaultAuthService
	SessionService      *session.DefaultSessionService
}

func startServices(userDataPath, credentialsDataPath string) *Services {
	HashingService := &hashing.DefaultHashingService{}
	services := &Services{
		HashingService: HashingService,
		RegistrationService: &registration.DefaultRegistrationService{
			HashingService: HashingService,
			FilePath:       credentialsDataPath,
			UserDataPath:   userDataPath,
		},
		ValidationService: &validation.DefaultValidationService{},
		AuthService: &auth.DefaultAuthService{
			FilePath: credentialsDataPath,
		},
		SessionService: &session.DefaultSessionService{
			SessionMap: make(map[string]string),
		},
	}
	return services
}
