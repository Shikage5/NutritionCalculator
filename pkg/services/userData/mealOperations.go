package userData

import (
	"NutritionCalculator/data/models"
	"errors"
	"fmt"
)

var ErrMealAlreadyExists = errors.New("meal already exists")
var ErrMealNotFound = errors.New("meal not found")

func (s *DefaultUserDataService) GetMeals() ([]models.Meal, error) {
	savedData, err := s.GetUserData()
	if err != nil {
		return nil, err
	}
	return savedData.Meals, nil
}

func (s *DefaultUserDataService) AddMeal(meal models.Meal) error {
	savedData, err := s.GetUserData()
	if err != nil {
		return err
	}
	for _, m := range savedData.Meals {
		if m.Name == meal.Name {
			return ErrMealAlreadyExists
		}
	}
	savedData.Meals = append(savedData.Meals, meal)
	return s.SaveUserData(savedData)
}

func (s *DefaultUserDataService) UpdateMeal(meal models.Meal) error {
	savedData, err := s.GetUserData()
	if err != nil {
		return err
	}
	for i, m := range savedData.Meals {
		if m.Name == meal.Name {
			savedData.Meals[i] = meal
			break
		} else if i == len(savedData.Meals)-1 {
			return ErrMealNotFound
		}
	}
	return s.SaveUserData(savedData)
}

func (s *DefaultUserDataService) DeleteMeal(meal models.Meal) error {
	savedData, err := s.GetUserData()
	if err != nil {
		return err
	}
	for i, m := range savedData.Meals {
		if m.Name == meal.Name {
			savedData.Meals = append(savedData.Meals[:i], savedData.Meals[i+1:]...)
			break
		} else if i == len(savedData.Meals)-1 {
			return ErrMealNotFound
		}
	}
	return s.SaveUserData(savedData)
}

func (s *DefaultUserDataService) CalculateMealNutritionalValues(meal models.Meal, processedDishes map[string]bool) (models.NutritionalValues, error) {
	var totalMealNutritionalValues models.NutritionalValues

	/*==========================Add Nutritional Values of all Foods=============================*/

	for _, food := range meal.Foods {
		foodNutritionalValues, err := s.CalculateFoodNutritionalValues(food)
		if err != nil {
			return models.NutritionalValues{}, err
		}
		totalMealNutritionalValues = s.AddNutritions(totalMealNutritionalValues, foodNutritionalValues)
	}

	/*==========================Add Nutritional Values of all Dishes=============================*/
	for _, dish := range meal.Dishes {
		//Circular reference check
		if processedDishes[dish.Name] {
			return models.NutritionalValues{}, fmt.Errorf("circular reference detected with dish: %s", dish.Name)
		}
		processedDishes[dish.Name] = true

		dishNutritionalValues, err := s.CalculateDishNutritionalValues(dish, processedDishes)
		if err != nil {
			return models.NutritionalValues{}, err
		}
		totalMealNutritionalValues = s.AddNutritions(totalMealNutritionalValues, dishNutritionalValues)
	}

	return totalMealNutritionalValues, nil
}
