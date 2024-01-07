package handlers

import (
	"NutritionCalculator/data/models"
	contextkeys "NutritionCalculator/pkg/contextKeys"
	"NutritionCalculator/pkg/services/userData"
	"NutritionCalculator/pkg/services/validation"
	"NutritionCalculator/utils"
	"encoding/json"
	"errors"
	"log"
	"net/http"
)

type DayHandler struct {
	DayService        userData.DefaultDayService
	MealService       userData.DefaultMealService
	ValidationService validation.DefaultValidationService
}

func NewDayHandler(dayService userData.DefaultDayService, mealService userData.DefaultMealService, validationService validation.DefaultValidationService) *DayHandler {
	return &DayHandler{DayService: dayService, MealService: mealService, ValidationService: validationService}
}

func (h *DayHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	username := r.Context().Value(contextkeys.UserKey).(string)

	switch r.Method {
	case http.MethodGet:
		h.handleGet(w, username)
	case http.MethodPost:
		h.handlePost(w, r, username)
	case http.MethodPut:
		h.handlePut(w, r, username)
	case http.MethodDelete:
		h.handleDelete(w, r, username)
	default:
		h.handleError(w, errors.New("invalid method"), http.StatusMethodNotAllowed)
	}
}

func (h *DayHandler) handleGet(w http.ResponseWriter, username string) {
	h.displayDayPage(w, username)
}

func (h *DayHandler) handlePost(w http.ResponseWriter, r *http.Request, username string) {
	day, err := h.decodeAndValidateDay(r)
	if err != nil {
		h.handleError(w, err, http.StatusBadRequest)
		return
	}

	err = h.DayService.AddDay(username, day)
	if err != nil {
		h.handleError(w, err, http.StatusInternalServerError)
		return
	}

	w.Write([]byte("Day added!\n"))
	h.displayDayPage(w, username)
}

func (h *DayHandler) handlePut(w http.ResponseWriter, r *http.Request, username string) {
	day, err := h.decodeAndValidateDay(r)
	if err != nil {
		h.handleError(w, err, http.StatusBadRequest)
		return
	}

	nutritionalValues, err := h.DayService.CalculateDayNutritionalValues(day)
	if err != nil {
		h.handleError(w, err, http.StatusInternalServerError)
		return
	}
	day.NutritionalValues = &nutritionalValues

	err = h.DayService.UpdateDay(username, day)
	if err != nil {
		h.handleError(w, err, http.StatusInternalServerError)
		return
	}

	w.Write([]byte("Day updated!\n"))
	h.displayDayPage(w, username)
}

func (h *DayHandler) handleDelete(w http.ResponseWriter, r *http.Request, username string) {
	var day models.Day
	err := json.NewDecoder(r.Body).Decode(&day)
	if err != nil {
		h.handleError(w, err, http.StatusBadRequest)
		return
	}

	err = h.ValidationService.ValidateDayForDeletion(day)
	if err != nil {
		h.handleError(w, err, http.StatusBadRequest)
		return
	}

	err = h.DayService.DeleteDay(username, day)
	if err != nil {
		h.handleError(w, err, http.StatusInternalServerError)
		return
	}

	w.Write([]byte("Day deleted!\n"))
	h.displayDayPage(w, username)
}

func (h *DayHandler) decodeAndValidateDay(r *http.Request) (models.Day, error) {
	var day models.Day
	err := json.NewDecoder(r.Body).Decode(&day)
	if err != nil {
		return models.Day{}, err
	}

	err = h.ValidationService.ValidateDay(day)
	if err != nil {
		return models.Day{}, err
	}

	return day, nil
}

func (h *DayHandler) displayDayPage(w http.ResponseWriter, username string) {
	days, err := h.DayService.GetDays(username)
	if err != nil {
		h.handleError(w, err, http.StatusInternalServerError)
		return
	}

	utils.DisplayPage(w, days, "web/template/day.html")
}

func (h *DayHandler) handleError(w http.ResponseWriter, err error, statusCode int) {
	log.Println(err)
	http.Error(w, err.Error(), statusCode)
}
