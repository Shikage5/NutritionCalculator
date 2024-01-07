package userData

import (
	"NutritionCalculator/data/models"
	"errors"
	"fmt"
)

type DishDataService interface {
	AddDishData(username string, dishData models.DishData) error
	GetDishData(username string) ([]models.DishData, error)
	UpdateDishData(username string, dishData models.DishData) error
	DeleteDishData(username string, dishData models.DishData) error
	GetDishDataByName(string) (models.DishData, error)
	RecalculateNutritionalValuesOfDishes([]models.DishData) ([]models.DishData, error)
	CalculateDishDataNutritionalValues(models.DishData) (models.NutritionalValues, error)
}

type DefaultDishDataService struct {
	UserDataService        UserDataService
	FoodService            FoodService
	DishService            DishService
	MealService            MealService
	DayService             DayService
	NutritionValuesService NutritionValuesService
}

var ErrDishAlreadyExists = errors.New("dish already exists")
var ErrDishNotFound = errors.New("dish not found")

func NewDishDataService(userDataService UserDataService, foodService FoodService, dishService DishService, mealService MealService, dayService DayService, nutritionValuesService NutritionValuesService) *DefaultDishDataService {
	return &DefaultDishDataService{UserDataService: userDataService, FoodService: foodService, DishService: dishService, MealService: mealService, DayService: dayService, NutritionValuesService: nutritionValuesService}
}

/*==========================CRUD=============================*/

func (s *DefaultDishDataService) AddDishData(username string, dishData models.DishData) error {
	savedData, err := s.UserDataService.GetUserData(username)
	if err != nil {
		return err
	}
	for _, f := range savedData.DishData {
		if f.Name == dishData.Name {
			return ErrDishAlreadyExists
		}
	}
	savedData.DishData = append(savedData.DishData, dishData)
	return s.UserDataService.SaveUserData(savedData)
}

func (s *DefaultDishDataService) GetDishData(username string) ([]models.DishData, error) {
	savedData, err := s.UserDataService.GetUserData(username)
	if err != nil {
		return nil, err
	}
	return savedData.DishData, nil

}

func (s *DefaultDishDataService) UpdateDishData(username string, dishData models.DishData) error {
	savedData, err := s.UserDataService.GetUserData(username)
	if err != nil {
		return err
	}
	for i, f := range savedData.DishData {
		if f.Name == dishData.Name {
			savedData.DishData[i] = dishData
			break
		} else if i == len(savedData.DishData)-1 {
			return fmt.Errorf("%w: %s", ErrDishNotFound, dishData.Name)
		}
	}

	//save the updated DishData
	err = s.UserDataService.SaveUserData(savedData)
	if err != nil {
		return err
	}
	savedData.DishData, err = s.RecalculateNutritionalValuesOfDishes(savedData.DishData)
	if err != nil {
		return err
	}
	savedData.Meals, err = s.MealService.RecalculateNutritionalValuesOfMeals(savedData.Meals)
	if err != nil {
		return err
	}
	savedData.Days, err = s.DayService.RecalculateNutritionalValuesOfDays(savedData.Days)
	if err != nil {
		return err
	}

	return s.UserDataService.SaveUserData(savedData)
}

func (s *DefaultDishDataService) DeleteDishData(username string, dishData models.DishData) error {
	savedData, err := s.UserDataService.GetUserData(username)
	if err != nil {
		return err
	}
	for i, d := range savedData.DishData {
		if d.Name == dishData.Name {
			savedData.DishData = append(savedData.DishData[:i], savedData.DishData[i+1:]...)

			// Delete the dish from all other dishes and meals
			savedData.DishData = s.deleteDishFromDishes(dishData.Name, savedData.DishData)
			savedData.Meals = s.deleteDishFromMeals(dishData.Name, savedData.Meals)

			//recalculate nutritional values of all dishes and meals
			savedData.DishData, err = s.RecalculateNutritionalValuesOfDishes(savedData.DishData)
			if err != nil {
				return err
			}
			savedData.Meals, err = s.MealService.RecalculateNutritionalValuesOfMeals(savedData.Meals)
			if err != nil {
				return err
			}

			break
		} else if i == len(savedData.DishData)-1 {
			return fmt.Errorf("%w: %s", ErrDishNotFound, dishData.Name)
		}
	}
	return s.UserDataService.SaveUserData(savedData)
}

/*==========================Delete Helper Functions=============================*/
func (s *DefaultDishDataService) deleteDishFromDishes(dishName string, dishes []models.DishData) []models.DishData {
	for i, dish := range dishes {
		for j, subDish := range dish.Dishes {
			if subDish.Name == dishName {
				dishes[i].Dishes = append(dish.Dishes[:j], dish.Dishes[j+1:]...)
				break
			}
		}
	}
	return dishes
}

func (s *DefaultDishDataService) deleteDishFromMeals(dishName string, meals []models.Meal) []models.Meal {
	for i, meal := range meals {
		for j, dish := range meal.Dishes {
			if dish.Name == dishName {
				meals[i].Dishes = append(meal.Dishes[:j], meal.Dishes[j+1:]...)
				break
			}
		}
	}
	return meals
}
func (s *DefaultDishDataService) RecalculateNutritionalValuesOfDishes(dishes []models.DishData) ([]models.DishData, error) {
	for i, dish := range dishes {
		nutritionalValues, err := s.CalculateDishDataNutritionalValues(dish)
		if err != nil {
			return nil, err
		}
		dishes[i].NutritionalValues = &nutritionalValues
	}
	return dishes, nil
}

/*==========================Specific DishData=============================*/
func (s *DefaultDishDataService) GetDishDataByName(name string) (models.DishData, error) {
	savedData, err := s.UserDataService.GetUserData()
	if err != nil {
		return models.DishData{}, err
	}
	for _, f := range savedData.DishData {
		if f.Name == name {
			return f, nil
		}
	}
	return models.DishData{}, fmt.Errorf("%w: %s", ErrDishNotFound, name)
}

/*==========================Nutritional Values=============================*/
func (s *DefaultDishDataService) CalculateDishDataNutritionalValues(dishData models.DishData) (models.NutritionalValues, error) {
	var totalDishNutritionalValues models.NutritionalValues

	// Add Nutritional Values of all Foods
	for i, food := range dishData.Foods {
		foodNutritionalValues, err := s.FoodService.CalculateFoodNutritionalValues(food)
		if err != nil {
			return models.NutritionalValues{}, err
		}
		foodWeight, err := s.FoodService.CalculateFoodWeight(food)
		if err != nil {
			return models.NutritionalValues{}, err
		}
		dishData.Foods[i].NutritionalValues = &foodNutritionalValues
		dishData.Foods[i].Weight = &foodWeight

		totalDishNutritionalValues = s.NutritionValuesService.AddNutritions(totalDishNutritionalValues, foodNutritionalValues)

	}

	// Add Nutritional Values of all Dishes
	for i, dish := range dishData.Dishes {
		dishNutritionalValues, err := s.DishService.CalculateDishNutritionalValues(dish, make(map[string]bool))
		if err != nil {
			return models.NutritionalValues{}, err
		}
		dishWeight, err := s.DishService.CalculateDishWeight(dish, make(map[string]bool))
		if err != nil {
			return models.NutritionalValues{}, err
		}
		dishData.Dishes[i].NutritionalValues = &dishNutritionalValues
		dishData.Dishes[i].Weight = &dishWeight

		totalDishNutritionalValues = s.NutritionValuesService.AddNutritions(totalDishNutritionalValues, dishNutritionalValues)
	}

	return totalDishNutritionalValues, nil
}
