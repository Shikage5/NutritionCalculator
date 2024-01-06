package main

import (
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
)

func main() {
	port := flag.String("port", "443", "port to run the server on")
	userDataPath := "data/user_data/"
	credentialsDataPath := "data/userCredentials.json"
	flag.Parse()

	s := startServices(userDataPath, credentialsDataPath)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	})

	http.HandleFunc("/login", middleware.ValidateUser(s.ValidationService, handlers.LoginHandler(s.AuthService, s.SessionService)))
	http.HandleFunc("/register", middleware.ValidateUser(s.ValidationService, handlers.RegisterHandler(s.RegistrationService)))

	http.HandleFunc("/home", middleware.SessionMiddleware(s.SessionService, handlers.HomeHandler(userDataPath)))

	http.HandleFunc("/food", middleware.SessionMiddleware(s.SessionService, handlers.FoodHandler(userDataPath)))
	http.HandleFunc("/dish", middleware.SessionMiddleware(s.SessionService, handlers.DishHandler(userDataPath)))
	http.HandleFunc("/meal", middleware.SessionMiddleware(s.SessionService, handlers.MealHandler(userDataPath)))
	http.HandleFunc("/day", middleware.SessionMiddleware(s.SessionService, handlers.DayHandler(userDataPath)))

	http.HandleFunc("/export", middleware.SessionMiddleware(s.SessionService, handlers.ExportHandler(userDataPath)))

	// Get the absolute path to the certificate file
	certFilePath := "server.crt"
	keyFilePath := "server.key"

	log.Println("Starting server on port:", *port)
	err := http.ListenAndServeTLS(":"+*port, certFilePath, keyFilePath, nil)
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
