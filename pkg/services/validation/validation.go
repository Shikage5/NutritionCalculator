package validation

import (
	"NutritionCalculator/data/models"
)

type ValidationService interface {
	ValidateCredentials(username, password string) bool
	ValidateFoodData(foodData models.FoodData) error
}

type DefaultValidationService struct{}

func (v *DefaultValidationService) ValidateCredentials(username, password string) bool {
	return username != "" && password != ""
}

//TODO: Add validation

func (v *DefaultValidationService) ValidateFoodData(foodData models.FoodData) error {
	return nil
}

func (v *DefaultValidationService) ValidateDishData(dishData models.DishData) error {
	return nil
}

func (v *DefaultValidationService) ValidateMeal(meal models.Meal) error {
	return nil
}

func (v *DefaultValidationService) ValidateDay(day models.Day) error {
	return nil
}
