package handlers

import (
	"NutritionCalculator/data/models"
	contextkeys "NutritionCalculator/pkg/contextKeys"
	"NutritionCalculator/pkg/services/userData"
	"NutritionCalculator/utils"
	"errors"
	"log"
	"net/http"
	"time"
)

type HomeHandler struct {
	UserDataService userData.DefaultUserDataService
	DayService      userData.DefaultDayService
}

func NewHomeHandler(userDataService userData.DefaultUserDataService, dayService userData.DefaultDayService) *HomeHandler {
	return &HomeHandler{UserDataService: userDataService, DayService: dayService}
}

func (h *HomeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	username := r.Context().Value(contextkeys.UserKey).(string)

	switch r.Method {
	case http.MethodGet:
		h.handleGet(w, username)
	default:
		h.handleError(w, errors.New("invalid method"), http.StatusMethodNotAllowed)
	}
}

func (h *HomeHandler) handleGet(w http.ResponseWriter, username string) {
	userData, err := h.UserDataService.GetUserData(username)
	if err != nil {
		h.handleError(w, err, http.StatusInternalServerError)
		return
	}

	todaysDate := time.Now().Format("2006-01-02")
	today := models.Day{Date: todaysDate}
	err = h.DayService.AddDay(username, today)
	if err != nil {
		h.handleError(w, err, http.StatusInternalServerError)
		return
	}

	dayNutritionalValues, err := h.DayService.CalculateDayNutritionalValues(today)
	if err != nil {
		h.handleError(w, err, http.StatusInternalServerError)
		return
	}

	averageNutritionalValues, err := h.DayService.CalculateLastSevenDaysNutritionalValues(username)
	if err != nil {
		h.handleError(w, err, http.StatusInternalServerError)
		return
	}

	overviewData := struct {
		UserData                 models.UserData
		DayNutritionalValues     models.NutritionalValues
		AverageNutritionalValues models.NutritionalValues
	}{
		UserData:                 userData,
		DayNutritionalValues:     dayNutritionalValues,
		AverageNutritionalValues: averageNutritionalValues,
	}

	utils.DisplayPage(w, overviewData, "web/template/home.html")
}

func (h *HomeHandler) handleError(w http.ResponseWriter, err error, statusCode int) {
	log.Println(err)
	http.Error(w, err.Error(), statusCode)
}
