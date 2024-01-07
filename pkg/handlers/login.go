package handlers

import (
	"NutritionCalculator/data/models"
	contextkeys "NutritionCalculator/pkg/contextKeys"
	"NutritionCalculator/pkg/services/auth"
	"NutritionCalculator/pkg/services/session"
	"NutritionCalculator/pkg/services/validation"
	"NutritionCalculator/utils"
	"log"
	"net/http"
)

type LoginHandler struct {
	ValidationService validation.DefaultValidationService
	SessionService    session.DefaultSessionService
	AuthService       auth.DefaultAuthService
}

func NewLoginHandler(validationService validation.DefaultValidationService, sessionService session.DefaultSessionService, authService auth.DefaultAuthService) *LoginHandler {
	return &LoginHandler{ValidationService: validationService, SessionService: sessionService, AuthService: authService}
}

type LoginTemplateData struct {
	ErrMsg string
}

func (h *LoginHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var data LoginTemplateData
	if r.Method == http.MethodPost {
		h.handlePost(w, r, &data)
	} else {
		utils.DisplayPage(w, data, "web/template/login.html")
	}
}

func (h *LoginHandler) handlePost(w http.ResponseWriter, r *http.Request, data *LoginTemplateData) {
	userRequest, ok := r.Context().Value(contextkeys.UserRequestKey).(models.UserRequest)
	if !ok {
		h.handleError(w, "invalid form data", http.StatusBadRequest)
		return
	}

	err := h.ValidationService.ValidateUserRequest(userRequest)
	if err != nil {
		h.handleError(w, err.Error(), http.StatusBadRequest)
		return
	}

	authenticated, err := h.AuthService.Auth(userRequest)
	if err == auth.ErrInvalidCredentials {
		data.ErrMsg = err.Error()
		utils.DisplayPage(w, data, "web/template/login.html")
		return
	} else if err != nil {
		h.handleError(w, err.Error(), http.StatusInternalServerError)
		return
	} else if !authenticated {
		data.ErrMsg = "Authentication failed"
		utils.DisplayPage(w, data, "web/template/login.html")
		return
	}

	err = h.SessionService.CreateSession(userRequest.Username, w)
	if err != nil {
		h.handleError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/home", http.StatusSeeOther)
}

func (h *LoginHandler) handleError(w http.ResponseWriter, err string, statusCode int) {
	log.Println(err)
	http.Error(w, err, statusCode)
}
