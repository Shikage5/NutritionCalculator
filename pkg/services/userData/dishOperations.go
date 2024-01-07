package userData

import (
	"NutritionCalculator/data/models"
	"fmt"
	"log"
)

type DishService interface {
	CalculateDishNutritionalValues(username, dish models.Dish, processedDishes map[string]bool) (models.NutritionalValues, error)
	CalculateTotalDishWeight(dishes []models.Dish, processedDishes map[string]bool) float64
	CalculateDishWeight(dish models.Dish, processedDishes map[string]bool) (float64, error)
}

type DefaultDishService struct {
	DishDataService        DishDataService
	FoodService            FoodService
	NutritionValuesService NutritionValuesService
}

func NewDishService(dishDataService DishDataService, foodService FoodService, nutritionValuesService NutritionValuesService) *DefaultDishService {
	return &DefaultDishService{DishDataService: dishDataService, FoodService: foodService, NutritionValuesService: nutritionValuesService}
}

/*===========================Dish Operations=============================*/

func (s *DefaultDishService) CalculateDishNutritionalValues(username string, dish models.Dish, processedDishes map[string]bool) (models.NutritionalValues, error) {
	var totalDishNutritionalValues models.NutritionalValues
	//Circular reference check
	if processedDishes[dish.Name] {
		return models.NutritionalValues{}, fmt.Errorf("circular reference detected with dish: %s", dish.Name)
	}

	processedDishes[dish.Name] = true

	dishData, err := s.DishDataService.GetDishData(username, dish.Name)
	if err != nil {
		return models.NutritionalValues{}, err
	}

	// Add Nutritional Values of all Foods
	for _, food := range dishData.Foods {
		if err != nil {
			return models.NutritionalValues{}, err
		}

		foodNutritionalValues, err := s.FoodService.CalculateFoodNutritionalValues(username, food)
		if err != nil {
			return models.NutritionalValues{}, err
		}

		totalDishNutritionalValues = s.NutritionValuesService.AddNutritions(totalDishNutritionalValues, foodNutritionalValues)
	}

	// Add Nutritional Values of all Dishes
	totalDishWeight := s.CalculateTotalDishWeight(dishData.Dishes, processedDishes)

	for _, dish := range dishData.Dishes {
		dishWeight, err := s.CalculateDishWeight(dish, processedDishes)
		if err != nil {
			return models.NutritionalValues{}, err
		}
		ratio := dishWeight / totalDishWeight

		dishNutritionalValues, err := s.CalculateDishNutritionalValues(username, dish, processedDishes)
		if err != nil {
			return models.NutritionalValues{}, err
		}
		dishNutritionalValues = s.NutritionValuesService.AddNutritionsByRatio(ratio, dishNutritionalValues)
		totalDishNutritionalValues = s.NutritionValuesService.AddNutritions(totalDishNutritionalValues, dishNutritionalValues)
	}

	return totalDishNutritionalValues, nil
}

func (s *DefaultDishService) CalculateTotalDishWeight(dishes []models.Dish, processedDishes map[string]bool) float64 {
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

func (s *DefaultDishService) CalculateDishWeight(dish models.Dish, processedDishes map[string]bool) (float64, error) {

	if dish.Weight != nil {
		return *dish.Weight, nil
	} else if dish.Quantity != nil {
		//Circular reference check
		if processedDishes[dish.Name] {
			return 0, fmt.Errorf("circular reference detected with dish: %s", dish.Name)
		}
		processedDishes[dish.Name] = true

		var totalDishWeight float64
		dishData, err := s.DishDataService.GetDishDataByName(dish.Name)
		if err != nil {
			return 0, err
		}
		totalDishWeight += s.FoodService.CalculateTotalFoodWeight(dishData.Foods)
		totalDishWeight += s.CalculateTotalDishWeight(dishData.Dishes, processedDishes)

		return *dish.Quantity * totalDishWeight, nil
	}
	return 0, nil
}
