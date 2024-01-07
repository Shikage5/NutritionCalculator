package userData

import (
	"NutritionCalculator/data/models"
	"fmt"
)

type FoodService interface {
	CalculateFoodNutritionalValues(username string, food models.Food) (models.NutritionalValues, error)
	CalculateTotalFoodWeight([]models.Food) float64
	CalculateFoodWeight(models.Food) (float64, error)
}

type DefaultFoodService struct {
	FoodDataService        FoodDataService
	NutritionValuesService NutritionValuesService
}

func NewFoodService(foodDataService FoodDataService, nutritionValuesService NutritionValuesService) *DefaultFoodService {
	return &DefaultFoodService{FoodDataService: foodDataService, NutritionValuesService: nutritionValuesService}
}

/*===========================Food Operations=============================*/
func (s *DefaultFoodService) CalculateFoodNutritionalValues(username string, food models.Food) (models.NutritionalValues, error) {
	foodData, err := s.FoodDataService.GetFoodDataByName(username, food.Name)
	if err != nil {
		return models.NutritionalValues{}, err
	}
	var foodNutritionalValues models.NutritionalValues
	foodWeight, err := s.CalculateFoodWeight(food)
	if err != nil {
		return models.NutritionalValues{}, err
	}
	ratio := foodWeight / foodData.ReferenceWeight

	foodNutritionalValues = s.NutritionValuesService.AddNutritionsByRatio(ratio, *foodData.NutritionalValues)

	return foodNutritionalValues, nil
}

func (s *DefaultFoodService) CalculateTotalFoodWeight(foods []models.Food) float64 {
	var totalFoodWeight float64
	for _, food := range foods {
		foodWeight, err := s.CalculateFoodWeight(food)
		if err != nil {
			continue
		}
		totalFoodWeight += foodWeight
	}
	return totalFoodWeight
}
func (s *DefaultFoodService) CalculateFoodWeight(food models.Food) (float64, error) {
	if food.Weight != nil {
		return *food.Weight, nil
	} else if food.Quantity != nil {
		foodData, err := s.FoodDataService.GetFoodDataByName(food.Name)
		if err != nil {
			return 0, err
		}
		return *food.Quantity * foodData.MeasurementUnit.Weight, nil
	}
	return 0, fmt.Errorf("no weight or quantity specified for food %s", food.Name)
}
