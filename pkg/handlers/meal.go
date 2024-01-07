package handlers

import (
	"NutritionCalculator/data/models"
	contextkeys "NutritionCalculator/pkg/contextKeys"
	"NutritionCalculator/pkg/services/userData"
	"NutritionCalculator/pkg/services/validation"
	"NutritionCalculator/utils"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"
)

type MealHandler struct {
	MealService       userData.DefaultMealService
	ValidationService validation.DefaultValidationService
}

func NewMealHandler(mealService userData.DefaultMealService, validationService validation.DefaultValidationService) *MealHandler {
	return &MealHandler{MealService: mealService, ValidationService: validationService}
}

func (h *MealHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
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

func (h *MealHandler) handleGet(w http.ResponseWriter, username string) {
	h.displayMealPage(w, username)
}

func (h *MealHandler) handlePost(w http.ResponseWriter, r *http.Request, username string) {
	var meal models.Meal
	err := h.decodeAndValidate(r, &meal)
	if err != nil {
		h.handleError(w, err, http.StatusBadRequest)
		return
	}

	nutritionalValues, err := h.MealService.CalculateMealNutritionalValues(meal)
	if err != nil {
		h.handleError(w, err, http.StatusInternalServerError)
		return
	}

	meal.NutritionalValues = &nutritionalValues
	err = h.MealService.AddMeal(username, meal)
	if err != nil {
		h.handleError(w, err, http.StatusInternalServerError)
		return
	}

	w.Write([]byte("Meal added!\n"))
	h.displayMealPage(w, username)
}

func (h *MealHandler) handlePut(w http.ResponseWriter, r *http.Request, username string) {
	var meal models.Meal
	err := h.decodeAndValidate(r, &meal)
	if err != nil {
		h.handleError(w, err, http.StatusBadRequest)
		return
	}

	nutritionalValues, err := h.MealService.CalculateMealNutritionalValues(meal)
	if err != nil {
		if strings.Contains(err.Error(), "circular reference") {
			h.handleError(w, err, http.StatusBadRequest)
		} else {
			h.handleError(w, err, http.StatusInternalServerError)
		}
		return
	}

	meal.NutritionalValues = &nutritionalValues
	err = h.MealService.UpdateMeal(username, meal)
	if err != nil {
		h.handleError(w, err, http.StatusInternalServerError)
		return
	}

	w.Write([]byte("Meal updated!\n"))
	h.displayMealPage(w, username)
}

func (h *MealHandler) handleDelete(w http.ResponseWriter, r *http.Request, username string) {
	var meal models.Meal
	err := json.NewDecoder(r.Body).Decode(&meal)
	if err != nil {
		h.handleError(w, err, http.StatusBadRequest)
		return
	}

	err = h.ValidationService.ValidateMealForDeletion(meal)
	if err != nil {
		h.handleError(w, err, http.StatusBadRequest)
		return
	}

	err = h.MealService.DeleteMeal(username, meal)
	if err != nil {
		h.handleError(w, err, http.StatusInternalServerError)
		return
	}

	w.Write([]byte("Meal deleted!\n"))
	h.displayMealPage(w, username)
}

func (h *MealHandler) handleError(w http.ResponseWriter, err error, statusCode int) {
	log.Println(err)
	http.Error(w, err.Error(), statusCode)
}

func (h *MealHandler) decodeAndValidate(r *http.Request, meal *models.Meal) error {
	err := json.NewDecoder(r.Body).Decode(meal)
	if err != nil {
		return fmt.Errorf("failed to decode meal: %w", err)
	}

	err = h.ValidationService.ValidateMeal(*meal)
	if err != nil {
		return fmt.Errorf("invalid meal: %w", err)
	}

	return nil
}

func (h *MealHandler) displayMealPage(w http.ResponseWriter, username string) {
	meals, err := h.MealService.GetMeals(username)
	if err != nil {
		h.handleError(w, err, http.StatusInternalServerError)
		return
	}

	utils.DisplayPage(w, meals, "web/template/meal.html")
}
