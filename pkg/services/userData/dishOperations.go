package userData

import (
	"NutritionCalculator/data/models"
	"fmt"
)

func (s *DefaultUserDataService) CalculateDishNutritionalValues(username string, dish models.Dish) (models.NutritionalValues, error) {
	dishData, err := s.GetDishDataByName(username, dish.Name)
	if err != nil {
		return models.NutritionalValues{}, err
	}
	var dishNutritionalValues models.NutritionalValues
	totalFoodWeight := s.CalculateTotalFoodWeight(dishData.Foods)
	//add the nutritional values of each food in the dish
	for _, food := range dishData.Foods {
		foodWeight, err := s.CalculateFoodWeight(food)
		if err != nil {
			return models.NutritionalValues{}, err
		}
		ratio := foodWeight / totalFoodWeight

		foodNutritionalValues, err := s.CalculateFoodNutritionalValues(username, food)
		if err != nil {
			return models.NutritionalValues{}, err
		}
		dishNutritionalValues.Carbohydrates += foodNutritionalValues.Carbohydrates * ratio
		dishNutritionalValues.Energy += foodNutritionalValues.Energy * ratio
		dishNutritionalValues.Fat += foodNutritionalValues.Fat * ratio
		dishNutritionalValues.Fiber += foodNutritionalValues.Fiber * ratio
		dishNutritionalValues.Protein += foodNutritionalValues.Protein * ratio
		dishNutritionalValues.Salt += foodNutritionalValues.Salt * ratio
		dishNutritionalValues.SaturatedFattyAcids += foodNutritionalValues.SaturatedFattyAcids * ratio
		dishNutritionalValues.Sugar += foodNutritionalValues.Sugar * ratio
		dishNutritionalValues.Water += foodNutritionalValues.Water * ratio
	}

	if totalFoodWeight == 0 {
		return models.NutritionalValues{}, fmt.Errorf("no weight or quantity specified for dish %s", dish.Name)
	}

	return dishNutritionalValues, nil
}
