package validation

import (
	"NutritionCalculator/data/models"
	"errors"
	"fmt"
	"time"
)

type ValidationService interface {
	ValidateUserRequest(models.UserRequest) error
	ValidateFoodData(foodData models.FoodData) error
}

type DefaultValidationService struct{}

func (v *DefaultValidationService) ValidateUserRequest(userRequest models.UserRequest) error {
	if userRequest.Username != "" && userRequest.Password != "" {
		return nil
	}

	return errors.New("username and password should not be empty")
}

//TODO: Add validation

func (v *DefaultValidationService) ValidateFoodData(foodData models.FoodData) error {
	if foodData.Name == "" {
		return errors.New("food name should not be empty")
	}

	if foodData.NutritionalValues == nil {
		return errors.New("nutritional values should not be nil")
	}

	if foodData.NutritionalValues.Energy < 0 ||
		foodData.NutritionalValues.Fat < 0 ||
		foodData.NutritionalValues.SaturatedFattyAcids < 0 ||
		foodData.NutritionalValues.Carbohydrates < 0 ||
		foodData.NutritionalValues.Sugar < 0 ||
		foodData.NutritionalValues.Protein < 0 ||
		foodData.NutritionalValues.Salt < 0 ||
		foodData.NutritionalValues.Fiber < 0 ||
		foodData.NutritionalValues.Water < 0 {
		return errors.New("all nutritional values should be non-negative")
	}

	if foodData.ReferenceWeight <= 0 {
		return errors.New("reference weight should be positive")
	}

	if foodData.MeasurementUnit == nil {
		return errors.New("measurement unit should not be nil")
	}

	if foodData.MeasurementUnit.Name == "" {
		return errors.New("measurement unit name should not be empty")
	}

	if foodData.MeasurementUnit.Weight < 0 {
		return errors.New("measurement unit weight should be positive")
	}

	return nil
}

func (v *DefaultValidationService) ValidateFood(food models.Food) error {
	if food.Name == "" {
		return errors.New("food name should not be empty")
	}

	if food.Quantity != nil && food.Weight != nil {
		return errors.New("food should have either quantity or weight")
	}

	if food.Quantity != nil && *food.Quantity < 0 {
		return errors.New("food quantity should be non-negative")
	}

	if food.Weight != nil && *food.Weight < 0 {
		return errors.New("food weight should be non-negative")
	}

	return nil
}

func (v *DefaultValidationService) ValidateDishData(dishData models.DishData) error {
	if dishData.Name == "" {
		return errors.New("dish name should not be empty")
	}

	if dishData.Foods == nil && dishData.Dishes == nil {
		return errors.New("foods and dishes should not be nil")
	}

	for _, food := range dishData.Foods {
		if err := v.ValidateFood(food); err != nil {
			return fmt.Errorf("invalid food: %w", err)
		}
	}

	for _, dish := range dishData.Dishes {
		if err := v.ValidateDish(dish); err != nil {
			return fmt.Errorf("invalid dish: %w", err)
		}
	}
	return nil
}

func (v *DefaultValidationService) ValidateDish(dish models.Dish) error {
	if dish.Name == "" {
		return errors.New("food name should not be empty")
	}

	if dish.Quantity != nil && dish.Weight != nil {
		return errors.New("food should have either quantity or weight")
	}

	if dish.Quantity != nil && *dish.Quantity < 0 {
		return errors.New("food quantity should be non-negative")
	}

	if dish.Weight != nil && *dish.Weight < 0 {
		return errors.New("food weight should be non-negative")
	}

	return nil
}

func (v *DefaultValidationService) ValidateMeal(meal models.Meal) error {
	if meal.Name == "" {
		return errors.New("meal name should not be empty")
	}

	if meal.Dishes == nil && meal.Foods == nil {
		return errors.New("foods and dishes should not be nil")
	}

	for _, food := range meal.Foods {
		if err := v.ValidateFood(food); err != nil {
			return fmt.Errorf("invalid food: %w", err)
		}
	}

	for _, dish := range meal.Dishes {
		if err := v.ValidateDish(dish); err != nil {
			return fmt.Errorf("invalid dish: %w", err)
		}
	}
	return nil
}

func (v *DefaultValidationService) ValidateDay(day models.Day) error {
	if day.Date == "" {
		return errors.New("date should not be empty")
	}

	// Check if the date is valid
	_, err := time.Parse("2006-01-02", day.Date)
	if err != nil {
		return errors.New("date should be a valid date")
	}

	if day.Meals == nil {
		return errors.New("meals should not be nil")
	}

	for _, meal := range day.Meals {
		if err := v.ValidateMeal(meal); err != nil {
			return fmt.Errorf("invalid meal: %w", err)
		}
	}
	return nil
}

func (v *DefaultValidationService) ValidateFoodDataForDeletion(foodData models.FoodData) error {
	if foodData.Name == "" {
		return errors.New("food name should not be empty")
	}

	return nil
}

func (v *DefaultValidationService) ValidateDishDataForDeletion(dishData models.DishData) error {
	if dishData.Name == "" {
		return errors.New("dish name should not be empty")
	}

	return nil
}

func (v *DefaultValidationService) ValidateMealForDeletion(meal models.Meal) error {
	if meal.Name == "" {
		return errors.New("meal name should not be empty")
	}

	return nil
}

func (v *DefaultValidationService) ValidateDayForDeletion(day models.Day) error {
	if day.Date == "" {
		return errors.New("date should not be empty")
	}

	return nil
}
