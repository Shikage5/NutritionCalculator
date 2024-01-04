package userData

import (
	"NutritionCalculator/data/models"
	"NutritionCalculator/utils"
)

//==========================Service for handling user data=============================

type UserDataService interface {
	// User operations
	GetUserData(username string) (models.UserData, error)
	SaveUserData(userData models.UserData, username string) error
	// FoodData operations
	GetFoodData(username string) ([]models.FoodData, error)
	AddFoodData(username string, foodData models.FoodData) error
	UpdateFoodData(username string, foodData models.FoodData) error
	DeleteFoodData(username string, foodData models.FoodData) error

	//Food Operations
	CalculateFoodNutritionalValues(food models.Food) (models.NutritionalValues, error)
	CalculateTotalFoodWeight(foods []models.Food) float64
	CalculateFoodWeight(food models.Food) (float64, error)

	// DishData operations
	GetDishData(username string) ([]models.DishData, error)
	AddDishData(username string, dishData models.DishData) error
	UpdateDishData(username string, dishData models.DishData) error
	DeleteDishData(username string, dishData models.DishData) error
	CalculateDishDataNutritionalValues(username string, dishData models.DishData) (models.NutritionalValues, error)
	//Dish Operations
	CalculateDishNutritionalValues(username string, dish models.Dish, processedDishes map[string]bool) (models.NutritionalValues, error)
	CalculateTotalDishWeight(dishes []models.Dish) float64
	CalculateDishWeight(dish models.Dish) (float64, error)
	// Meal operations

}

type DefaultUserDataService struct {
	UserDataPath string
}

func (s *DefaultUserDataService) GetUserData(username string) (models.UserData, error) {
	var userData models.UserData
	userData.Username = username
	err := utils.ReadUserDataFromJSONFile(&userData, s.UserDataPath)
	if err != nil {
		return models.UserData{}, err
	}
	return userData, err
}

func (s *DefaultUserDataService) SaveUserData(userData models.UserData, username string) error {
	err := utils.WriteUserDataToJSONFile(&userData, s.UserDataPath)
	if err != nil {
		return err
	}
	return nil
}
