package handlers

import (
	"NutritionCalculator/data/models"
	"NutritionCalculator/pkg/services/registration"
	"NutritionCalculator/pkg/services/validation"
	"NutritionCalculator/utils"
	"encoding/json"
	"errors"
	"log"
	"net/http"
)

type RegisterHandler struct {
	RegistrationService registration.RegistrationService
}

func NewRegisterHandler(registrationService registration.RegistrationService) *RegisterHandler {
	return &RegisterHandler{RegistrationService: registrationService}
}

func (h *RegisterHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.handleGet(w)
	case http.MethodPost:
		h.handlePost(w, r)
	default:
		h.handleError(w, errors.New("invalid method"), http.StatusMethodNotAllowed)
	}
}

func (h *RegisterHandler) handleGet(w http.ResponseWriter) {
	utils.DisplayPage(w, nil, "web/template/register.html")
}

func (h *RegisterHandler) handlePost(w http.ResponseWriter, r *http.Request) {
	userRequest, err := h.decodeAndValidateUserRequest(r)
	if err != nil {
		h.handleError(w, err, http.StatusBadRequest)
		return
	}

	err = h.RegistrationService.RegisterUser(userRequest)
	if err != nil {
		h.handleError(w, err, http.StatusBadRequest)
		return
	}

	log.Println("User registered successfully!")
	utils.DisplayPage(w, nil, "web/template/register.html")
}

func (h *RegisterHandler) decodeAndValidateUserRequest(r *http.Request) (models.UserRequest, error) {
	var userRequest models.UserRequest
	err := json.NewDecoder(r.Body).Decode(&userRequest)
	if err != nil {
		return models.UserRequest{}, err
	}

	validationService := &validation.DefaultValidationService{}
	err = validationService.ValidateUserRequest(userRequest)
	if err != nil {
		return models.UserRequest{}, err
	}

	return userRequest, nil
}

func (h *RegisterHandler) handleError(w http.ResponseWriter, err error, statusCode int) {
	log.Println(err)
	http.Error(w, err.Error(), statusCode)
}
