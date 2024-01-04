package userData

import (
	"NutritionCalculator/data/models"
	"fmt"
)

func (s *DefaultUserDataService) CalculateFoodNutritionalValues(username string, food models.Food) (models.NutritionalValues, error) {
	foodData, err := s.GetFoodDataByName(username, food.Name)
	if err != nil {
		return models.NutritionalValues{}, err
	}
	if food.Weight != nil {
		var foodNutritionalValues models.NutritionalValues
		ratio := *food.Weight / foodData.ReferenceWeight
		foodNutritionalValues.Carbohydrates = foodData.NutritionalValues.Carbohydrates * ratio
		foodNutritionalValues.Energy = foodData.NutritionalValues.Energy * ratio
		foodNutritionalValues.Fat = foodData.NutritionalValues.Fat * ratio
		foodNutritionalValues.Fiber = foodData.NutritionalValues.Fiber * ratio
		foodNutritionalValues.Protein = foodData.NutritionalValues.Protein * ratio
		foodNutritionalValues.Salt = foodData.NutritionalValues.Salt * ratio
		foodNutritionalValues.SaturatedFattyAcids = foodData.NutritionalValues.SaturatedFattyAcids * ratio
		foodNutritionalValues.Sugar = foodData.NutritionalValues.Sugar * ratio
		foodNutritionalValues.Water = foodData.NutritionalValues.Water * ratio

		return foodNutritionalValues, nil
	}
	if food.Quantity != nil {
		var foodNutritionalValues models.NutritionalValues
		ratio := *food.Quantity / foodData.MeasurementUnit.Weight
		foodNutritionalValues.Carbohydrates = foodData.NutritionalValues.Carbohydrates * ratio
		foodNutritionalValues.Energy = foodData.NutritionalValues.Energy * ratio
		foodNutritionalValues.Fat = foodData.NutritionalValues.Fat * ratio
		foodNutritionalValues.Fiber = foodData.NutritionalValues.Fiber * ratio
		foodNutritionalValues.Protein = foodData.NutritionalValues.Protein * ratio
		foodNutritionalValues.Salt = foodData.NutritionalValues.Salt * ratio
		foodNutritionalValues.SaturatedFattyAcids = foodData.NutritionalValues.SaturatedFattyAcids * ratio
		foodNutritionalValues.Sugar = foodData.NutritionalValues.Sugar * ratio
		foodNutritionalValues.Water = foodData.NutritionalValues.Water * ratio

		return foodNutritionalValues, nil
	}
	return models.NutritionalValues{}, fmt.Errorf("no weight or quantity specified for food %s", food.Name)
}

func (s *DefaultUserDataService) CalculateTotalFoodWeight(username string, foods []models.Food) float64 {
	var totalFoodWeight float64
	for _, food := range foods {
		foodWeight, err := s.CalculateFoodWeight(username, food)
		if err != nil {
			continue
		}
		totalFoodWeight += foodWeight
	}
	return totalFoodWeight
}
func (s *DefaultUserDataService) CalculateFoodWeight(username string, food models.Food) (float64, error) {
	if food.Weight != nil {
		return *food.Weight, nil
	} else if food.Quantity != nil {
		foodData, err := s.GetFoodDataByName(username, food.Name)
		if err != nil {
			return 0, err
		}
		return *food.Quantity * foodData.MeasurementUnit.Weight, nil
	}
	return 0, fmt.Errorf("no weight or quantity specified for food %s", food.Name)
}
