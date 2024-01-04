package userData

import (
	"NutritionCalculator/data/models"
	"fmt"
	"log"
)

func (s *DefaultUserDataService) CalculateDishNutritionalValues(username string, dish models.Dish, processedDishes map[string]bool) (models.NutritionalValues, error) {
	var totalDishNutritionalValues models.NutritionalValues

	if processedDishes[dish.Name] {
		return models.NutritionalValues{}, fmt.Errorf("circular reference detected with dish: %s", dish.Name)
	}

	processedDishes[dish.Name] = true

	dishData, err := s.GetDishDataByName(username, dish.Name)
	if err != nil {
		return models.NutritionalValues{}, err
	}

	/*==========================Add Nutritional Values of all Foods=============================*/

	totalFoodWeight := s.CalculateTotalFoodWeight(username, dishData.Foods)
	for _, food := range dishData.Foods {
		foodWeight, err := s.CalculateFoodWeight(username, food)
		if err != nil {
			return models.NutritionalValues{}, err
		}
		ratio := foodWeight / totalFoodWeight

		foodNutritionalValues, err := s.CalculateFoodNutritionalValues(username, food)
		if err != nil {
			return models.NutritionalValues{}, err
		}
		totalDishNutritionalValues.Carbohydrates += foodNutritionalValues.Carbohydrates * ratio
		totalDishNutritionalValues.Energy += foodNutritionalValues.Energy * ratio
		totalDishNutritionalValues.Fat += foodNutritionalValues.Fat * ratio
		totalDishNutritionalValues.Fiber += foodNutritionalValues.Fiber * ratio
		totalDishNutritionalValues.Protein += foodNutritionalValues.Protein * ratio
		totalDishNutritionalValues.Salt += foodNutritionalValues.Salt * ratio
		totalDishNutritionalValues.SaturatedFattyAcids += foodNutritionalValues.SaturatedFattyAcids * ratio
		totalDishNutritionalValues.Sugar += foodNutritionalValues.Sugar * ratio
		totalDishNutritionalValues.Water += foodNutritionalValues.Water * ratio
	}

	/*==========================Add Nutritional Values of all Dishes=============================*/

	totalDishWeight := s.CalculateTotalDishWeight(username, dishData.Dishes)

	for _, dish := range dishData.Dishes {
		dishWeight, err := s.CalculateDishWeight(username, dish)
		if err != nil {
			return models.NutritionalValues{}, err
		}
		ratio := dishWeight / totalDishWeight

		dishNutritionalValues, err := s.CalculateDishNutritionalValues(username, dish, processedDishes)
		if err != nil {
			return models.NutritionalValues{}, err
		}
		totalDishNutritionalValues.Carbohydrates += dishNutritionalValues.Carbohydrates * ratio
		totalDishNutritionalValues.Energy += dishNutritionalValues.Energy * ratio
		totalDishNutritionalValues.Fat += dishNutritionalValues.Fat * ratio
		totalDishNutritionalValues.Fiber += dishNutritionalValues.Fiber * ratio
		totalDishNutritionalValues.Protein += dishNutritionalValues.Protein * ratio
		totalDishNutritionalValues.Salt += dishNutritionalValues.Salt * ratio
		totalDishNutritionalValues.SaturatedFattyAcids += dishNutritionalValues.SaturatedFattyAcids * ratio
		totalDishNutritionalValues.Sugar += dishNutritionalValues.Sugar * ratio
		totalDishNutritionalValues.Water += dishNutritionalValues.Water * ratio
	}

	return totalDishNutritionalValues, nil
}

func (s *DefaultUserDataService) CalculateTotalDishWeight(username string, dishes []models.Dish) float64 {
	var totalDishWeight float64
	for _, dish := range dishes {
		dishWeight, err := s.CalculateDishWeight(username, dish)
		if err != nil {
			log.Println(err)
			continue
		}
		totalDishWeight += dishWeight
	}
	return totalDishWeight
}

func (s *DefaultUserDataService) CalculateDishWeight(username string, dish models.Dish) (float64, error) {

	if dish.Weight != nil {
		return *dish.Weight, nil
	} else if dish.Quantity != nil {
		var totalDishWeight float64
		dishData, err := s.GetDishDataByName(username, dish.Name)
		if err != nil {
			return 0, err
		}
		totalDishWeight += s.CalculateTotalFoodWeight(username, dishData.Foods)
		totalDishWeight += s.CalculateTotalDishWeight(username, dishData.Dishes)

		return *dish.Quantity * totalDishWeight, nil
	}
	return 0, nil
}
