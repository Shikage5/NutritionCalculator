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

type FoodHandler struct {
	FoodDataService   userData.DefaultFoodDataService
	ValidationService validation.DefaultValidationService
}

func NewFoodHandler(foodDataService userData.DefaultFoodDataService, validationService validation.DefaultValidationService) *FoodHandler {
	return &FoodHandler{FoodDataService: foodDataService, ValidationService: validationService}
}

func (h *FoodHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
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

func (h *FoodHandler) handleGet(w http.ResponseWriter, username string) {
	h.displayFoodPage(w, username)
}

func (h *FoodHandler) handlePost(w http.ResponseWriter, r *http.Request, username string) {
	foodData, err := h.decodeAndValidateFood(r)
	if err != nil {
		h.handleError(w, err, http.StatusBadRequest)
		return
	}

	err = h.FoodDataService.AddFoodData(username, foodData)
	if err != nil {
		h.handleError(w, err, http.StatusInternalServerError)
		return
	}

	w.Write([]byte("Food added!\n"))
	h.displayFoodPage(w, username)
}

func (h *FoodHandler) handlePut(w http.ResponseWriter, r *http.Request, username string) {
	foodData, err := h.decodeAndValidateFood(r)
	if err != nil {
		h.handleError(w, err, http.StatusBadRequest)
		return
	}

	err = h.FoodDataService.UpdateFoodData(username, foodData)
	if err != nil {
		h.handleError(w, err, http.StatusInternalServerError)
		return
	}

	w.Write([]byte("Food updated!\n"))
	h.displayFoodPage(w, username)
}

func (h *FoodHandler) handleDelete(w http.ResponseWriter, r *http.Request, username string) {
	var foodData models.FoodData
	err := json.NewDecoder(r.Body).Decode(&foodData)
	if err != nil {
		h.handleError(w, err, http.StatusBadRequest)
		return
	}

	err = h.ValidationService.ValidateFoodDataForDeletion(foodData)
	if err != nil {
		h.handleError(w, err, http.StatusBadRequest)
		return
	}

	err = h.FoodDataService.DeleteFoodData(username, foodData)
	if err != nil {
		h.handleError(w, err, http.StatusInternalServerError)
		return
	}

	w.Write([]byte("Food deleted!\n"))
	h.displayFoodPage(w, username)
}

func (h *FoodHandler) decodeAndValidateFood(r *http.Request) (models.FoodData, error) {
	var foodData models.FoodData
	err := json.NewDecoder(r.Body).Decode(&foodData)
	if err != nil {
		return models.FoodData{}, err
	}

	err = h.ValidationService.ValidateFoodData(foodData)
	if err != nil {
		return models.FoodData{}, err
	}
	return foodData, nil
}

func (h *FoodHandler) displayFoodPage(w http.ResponseWriter, username string) {
	foodData, err := h.FoodDataService.GetFoodData(username)
	if err != nil {
		h.handleError(w, err, http.StatusInternalServerError)
		return
	}

	utils.DisplayPage(w, foodData, "web/template/food.html")
}

func (h *FoodHandler) handleError(w http.ResponseWriter, err error, statusCode int) {
	log.Println(err)
	http.Error(w, err.Error(), statusCode)
}
