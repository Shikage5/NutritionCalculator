package validation

import (
	"NutritionCalculator/data/models"
	"errors"
	"regexp"
)

type ValidationService interface {
	ValidateCredentials(username, password string) bool
	ValidateFoodData(foodData models.FoodData) error
}

type DefaultValidationService struct{}

func (v *DefaultValidationService) ValidateCredentials(username, password string) bool {
	return username != "" && password != ""
}

func (v *DefaultValidationService) ValidateFoodData(foodData models.FoodData) error {
	//if food data is not fully complete, return an error

	match, _ := regexp.MatchString("^[a-zA-Z\\s]*$", foodData.Name)
	if !match {
		return errors.New("name must only contain letters and spaces")
	}

	return nil
}

func (v *DefaultValidationService) ValidateDishData(dishData models.DishData) error {
	match, _ := regexp.MatchString("^[a-zA-Z\\s]*$", dishData.Name)
	if !match {
		return errors.New("name must only contain letters and spaces")
	}

	return nil
}
