package main

import (
	"NutritionCalculator/pkg/handlers"
	"NutritionCalculator/pkg/middleware"
	"NutritionCalculator/pkg/services/auth"
	"NutritionCalculator/pkg/services/hashing"
	"NutritionCalculator/pkg/services/registration"
	"NutritionCalculator/pkg/services/session"
	"NutritionCalculator/pkg/services/userData"
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

	http.HandleFunc("/login", middleware.ValidateUser(s.ValidationService, handlers.NewLoginHandler(s.ValidationService, s.SessionService, s.AuthService)))
	http.HandleFunc("/register", middleware.ValidateUser(s.ValidationService, handlers.NewRegisterHandler(s.RegistrationService)))

	http.HandleFunc("/home", middleware.SessionMiddleware(s.SessionService, handlers.NewHomeHandler(userDataPath)))

	http.HandleFunc("/food", middleware.SessionMiddleware(s.SessionService, handlers.NewFoodHandler(userDataPath)))
	http.HandleFunc("/dish", middleware.SessionMiddleware(s.SessionService, handlers.NewDishHandler(userDataPath)))
	http.HandleFunc("/meal", middleware.SessionMiddleware(s.SessionService, handlers.NewMealHandler(userDataPath)))
	http.HandleFunc("/day", middleware.SessionMiddleware(s.SessionService, handlers.NewDayHandler()))

	http.HandleFunc("/export", middleware.SessionMiddleware(s.SessionService, handlers.NewExportHandler(userDataPath)))

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
	HashingService      hashing.DefaultHashingService
	AuthService         auth.DefaultAuthService
	ValidationService   validation.DefaultValidationService
	SessionService      session.DefaultSessionService
	RegistrationService registration.DefaultRegistrationService

	UserDataService        userData.DefaultUserDataService
	FoodService            userData.DefaultFoodService
	FoodDataService        userData.DefaultFoodDataService
	DishService            userData.DefaultDishService
	DishDataService        userData.DefaultDishDataService
	MealService            userData.DefaultMealService
	DayService             userData.DefaultDayService
	NutritionValuesService userData.DefaultNutritionValuesService
}

func startServices(userDataPath string, credentialsDataPath string) *Services {
    hashingService := &hashing.DefaultHashingService{}
    authService := &auth.DefaultAuthService{
        FilePath: credentialsDataPath,
    }
    validationService := &validation.DefaultValidationService{}
    sessionService := &session.DefaultSessionService{
        SessionMap: make(map[string]string),
    }
    registrationService := &registration.DefaultRegistrationService{
        HashingService:      hashingService,
        CredentialsFilepath: credentialsDataPath,
        UserDataPath:        userDataPath,
    }
    userDataService := &userData.DefaultUserDataService{
        UserDataPath: userDataPath,
    }
    nutritionValuesService := &userData.DefaultNutritionValuesService{}

    // Initialize the services that don't have any dependencies first
    foodService := &userData.DefaultFoodService{
        NutritionValuesService: nutritionValuesService,
    }
    dishService := &userData.DefaultDishService{
        NutritionValuesService: nutritionValuesService,
    }
    mealService := &userData.DefaultMealService{
        NutritionValuesService: nutritionValuesService,
    }
    dayService := &userData.DefaultDayService{
        NutritionValuesService: nutritionValuesService,
    }

    // Initialize DishDataService first without setting its FoodService field
    dishDataService := &userData.DefaultDishDataService{
        UserDataService:        userDataService,
        DishService:            dishService,
        MealService:            mealService,
        DayService:             dayService,
        NutritionValuesService: nutritionValuesService,
    }

    // Now initialize FoodDataService
    foodDataService := &userData.DefaultFoodDataService{
        FoodService:     foodService,
        DishDataService: dishDataService,
        MealService:     mealService,
        DayService:      dayService,
        UserDataService: userDataService,
    }

    // Finally, update the FoodService field of DishDataService
    dishDataService.FoodService = foodService

    // Update the dependent services with their dependencies
    foodService.FoodDataService = foodDataService
    dishService.DishDataService = dishDataService
    mealService.DayService = dayService
    dayService.MealService = mealService

    services := &Services{
        HashingService:         *hashingService,
        AuthService:            *authService,
        ValidationService:      *validationService,
        SessionService:         *sessionService,
        RegistrationService:    *registrationService,
        UserDataService:        *userDataService,
        FoodService:            *foodService,
        FoodDataService:        *foodDataService,
        DishService:            *dishService,
        DishDataService:        *dishDataService,
        MealService:            *mealService,
        DayService:             *dayService,
        NutritionValuesService: *nutritionValuesService,
    }

    return services
}