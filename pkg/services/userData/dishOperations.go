package userData

import (
	"NutritionCalculator/data/models"
	"fmt"
	"log"
)

/*===========================Dish Operations=============================*/

func (s *DefaultUserDataService) CalculateDishNutritionalValues(dish models.Dish, processedDishes map[string]bool) (models.NutritionalValues, error) {
	var totalDishNutritionalValues models.NutritionalValues
	//Circular reference check
	if processedDishes[dish.Name] {
		return models.NutritionalValues{}, fmt.Errorf("circular reference detected with dish: %s", dish.Name)
	}

	processedDishes[dish.Name] = true

	dishData, err := s.GetDishDataByName(dish.Name)
	if err != nil {
		return models.NutritionalValues{}, err
	}

	// Add Nutritional Values of all Foods
	for _, food := range dishData.Foods {
		if err != nil {
			return models.NutritionalValues{}, err
		}

		foodNutritionalValues, err := s.CalculateFoodNutritionalValues(food)
		if err != nil {
			return models.NutritionalValues{}, err
		}

		totalDishNutritionalValues = s.AddNutritions(totalDishNutritionalValues, foodNutritionalValues)
	}

	// Add Nutritional Values of all Dishes
	totalDishWeight := s.CalculateTotalDishWeight(dishData.Dishes, processedDishes)

	for _, dish := range dishData.Dishes {
		dishWeight, err := s.CalculateDishWeight(dish, processedDishes)
		if err != nil {
			return models.NutritionalValues{}, err
		}
		ratio := dishWeight / totalDishWeight

		dishNutritionalValues, err := s.CalculateDishNutritionalValues(dish, processedDishes)
		if err != nil {
			return models.NutritionalValues{}, err
		}
		dishNutritionalValues = s.AddNutritionsByRatio(ratio, dishNutritionalValues)
		totalDishNutritionalValues = s.AddNutritions(totalDishNutritionalValues, dishNutritionalValues)
	}

	return totalDishNutritionalValues, nil
}

func (s *DefaultUserDataService) CalculateTotalDishWeight(dishes []models.Dish, processedDishes map[string]bool) float64 {
	var totalDishWeight float64
	for _, dish := range dishes {
		dishWeight, err := s.CalculateDishWeight(dish, processedDishes)
		if err != nil {
			log.Println(err)
			continue
		}
		totalDishWeight += dishWeight
	}
	return totalDishWeight
}

func (s *DefaultUserDataService) CalculateDishWeight(dish models.Dish, processedDishes map[string]bool) (float64, error) {

	if dish.Weight != nil {
		return *dish.Weight, nil
	} else if dish.Quantity != nil {
		//Circular reference check
		if processedDishes[dish.Name] {
			return 0, fmt.Errorf("circular reference detected with dish: %s", dish.Name)
		}
		processedDishes[dish.Name] = true

		var totalDishWeight float64
		dishData, err := s.GetDishDataByName(dish.Name)
		if err != nil {
			return 0, err
		}
		totalDishWeight += s.CalculateTotalFoodWeight(dishData.Foods)
		totalDishWeight += s.CalculateTotalDishWeight(dishData.Dishes, processedDishes)

		return *dish.Quantity * totalDishWeight, nil
	}
	return 0, nil
}
