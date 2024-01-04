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
	match, _ := regexp.MatchString("^[a-zA-Z\\s]*$", foodData.Name)
	if !match {
		return errors.New("invalid name. Only letters and spaces are allowed")
	}

	if foodData.ReferenceWeight <= 0 {
		return errors.New("invalid reference weight. Must be a positive number")
	}

	return nil
}
