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

type DishHandler struct {
	DishDataService   userData.DefaultDishDataService
	ValidationService validation.DefaultValidationService
}

func NewDishHandler(dishDataService userData.DefaultDishDataService, validationService validation.DefaultValidationService) *DishHandler {
	return &DishHandler{DishDataService: dishDataService, ValidationService: validationService}
}

func (h *DishHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
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

func (h *DishHandler) handleGet(w http.ResponseWriter, username string) {
	h.displayDishPage(w, username)
}

func (h *DishHandler) handlePost(w http.ResponseWriter, r *http.Request, username string) {
	dishData, err := h.decodeAndValidateDish(r)
	if err != nil {
		h.handleError(w, err, http.StatusBadRequest)
		return
	}

	nutritionalValues, err := h.DishDataService.CalculateDishDataNutritionalValues(dishData)
	if err != nil {
		h.handleError(w, err, http.StatusInternalServerError)
		return
	}
	dishData.NutritionalValues = &nutritionalValues

	err = h.DishDataService.AddDishData(username, dishData)
	if err != nil {
		h.handleError(w, err, http.StatusInternalServerError)
		return
	}

	w.Write([]byte("Dish added!\n"))
	h.displayDishPage(w, username)
}

func (h *DishHandler) handlePut(w http.ResponseWriter, r *http.Request, username string) {
	dishData, err := h.decodeAndValidateDish(r)
	if err != nil {
		h.handleError(w, err, http.StatusBadRequest)
		return
	}

	nutritionalValues, err := h.DishDataService.CalculateDishDataNutritionalValues(dishData)
	if err != nil {
		h.handleError(w, err, http.StatusInternalServerError)
		return
	}
	dishData.NutritionalValues = &nutritionalValues

	err = h.DishDataService.UpdateDishData(username, dishData)
	if err != nil {
		h.handleError(w, err, http.StatusInternalServerError)
		return
	}

	w.Write([]byte("Dish updated!\n"))
	h.displayDishPage(w, username)
}

func (h *DishHandler) handleDelete(w http.ResponseWriter, r *http.Request, username string) {
	var dishData models.DishData
	err := json.NewDecoder(r.Body).Decode(&dishData)
	if err != nil {
		h.handleError(w, err, http.StatusBadRequest)
		return
	}

	err = h.ValidationService.ValidateDishDataForDeletion(dishData)
	if err != nil {
		h.handleError(w, err, http.StatusBadRequest)
		return
	}

	err = h.DishDataService.DeleteDishData(username, dishData)
	if err != nil {
		h.handleError(w, err, http.StatusInternalServerError)
		return
	}

	w.Write([]byte("Dish deleted!\n"))
	h.displayDishPage(w, username)
}

func (h *DishHandler) decodeAndValidateDish(r *http.Request) (models.DishData, error) {
	var dishData models.DishData
	err := json.NewDecoder(r.Body).Decode(&dishData)
	if err != nil {
		return dishData, err
	}

	err = h.ValidationService.ValidateDishData(dishData)
	if err != nil {
		log.Println(err)
		return dishData, err
	}

	return dishData, nil
}

func (h *DishHandler) displayDishPage(w http.ResponseWriter, username string) {
	dishData, err := h.DishDataService.GetDishData(username)
	if err != nil {
		h.handleError(w, err, http.StatusInternalServerError)
		return
	}

	utils.DisplayPage(w, dishData, "web/template/dish.html")
}

func (h *DishHandler) handleError(w http.ResponseWriter, err error, statusCode int) {
	log.Println(err)
	http.Error(w, err.Error(), statusCode)
}
