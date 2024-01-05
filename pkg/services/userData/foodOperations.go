package userData

import (
	"NutritionCalculator/data/models"
	"fmt"
)

func (s *DefaultUserDataService) CalculateFoodNutritionalValues(food models.Food) (models.NutritionalValues, error) {
	foodData, err := s.GetFoodDataByName(food.Name)
	if err != nil {
		return models.NutritionalValues{}, err
	}
	var foodNutritionalValues models.NutritionalValues
	foodWeight, err := s.CalculateFoodWeight(food)
	if err != nil {
		return models.NutritionalValues{}, err
	}
	ratio := foodWeight / foodData.ReferenceWeight

	foodNutritionalValues = s.AddNutritionsByRatio(ratio, foodData.NutritionalValues)

	return foodNutritionalValues, nil
}

func (s *DefaultUserDataService) CalculateTotalFoodWeight(foods []models.Food) float64 {
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
func (s *DefaultUserDataService) CalculateFoodWeight(food models.Food) (float64, error) {
	if food.Weight != nil {
		return *food.Weight, nil
	} else if food.Quantity != nil {
		foodData, err := s.GetFoodDataByName(food.Name)
		if err != nil {
			return 0, err
		}
		return *food.Quantity * foodData.MeasurementUnit.Weight, nil
	}
	return 0, fmt.Errorf("no weight or quantity specified for food %s", food.Name)
}
