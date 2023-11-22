package registrationHandlers

import (
	"NutritionCalculator/pkg/middleware"
	"NutritionCalculator/pkg/services/registration"
	"NutritionCalculator/pkg/services/validation"
	"net/http"
)

// RegisterHandler is the intermediate handler that dispatches the request to the appropriate specialized handler.
func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		RegisterGETHandler(w, r)
	case http.MethodPost:
		filePath := "data/users.json"
		validator := validation.NewCredentialsValidator()
		registrator := registration.NewRegistrationService(filePath)
		middleware.ValidateUser(validator, RegisterPOSTHandler(registrator))
	default:
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}
