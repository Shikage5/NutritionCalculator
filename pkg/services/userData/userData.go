package userData

import (
	"NutritionCalculator/data/models"
	"NutritionCalculator/utils"
)

//==========================Service for handling user data=============================

type UserDataService interface {
	// User operations
	GetUserData() (models.UserData, error)
	SaveUserData(userData models.UserData) error
	// FoodData operations
	GetFoodData() ([]models.FoodData, error)
	AddFoodData(foodData models.FoodData) error
	UpdateFoodData(foodData models.FoodData) error
	DeleteFoodData(foodData models.FoodData) error

	//Food Operations
	CalculateFoodNutritionalValues(food models.Food) (models.NutritionalValues, error)
	CalculateTotalFoodWeight(foods []models.Food) float64
	CalculateFoodWeight(food models.Food) (float64, error)

	// DishData operations
	GetDishData() ([]models.DishData, error)
	AddDishData(dishData models.DishData) error
	UpdateDishData(dishData models.DishData) error
	DeleteDishData(dishData models.DishData) error
	CalculateDishDataNutritionalValues(dishData models.DishData) (models.NutritionalValues, error)
	//Dish Operations
	CalculateDishNutritionalValues(dish models.Dish, processedDishes map[string]bool) (models.NutritionalValues, error)
	CalculateTotalDishWeight(dishes []models.Dish) float64
	CalculateDishWeight(dish models.Dish) (float64, error)
	// Meal operations

}

type DefaultUserDataService struct {
	UserDataPath string
	Username     string
}

func NewUserDataService(username string) *DefaultUserDataService {
	return &DefaultUserDataService{Username: username}
}

func (s *DefaultUserDataService) GetUserData() (models.UserData, error) {
	var userData models.UserData
	userData.Username = s.Username
	err := utils.ReadUserDataFromJSONFile(&userData, s.UserDataPath)
	if err != nil {
		return models.UserData{}, err
	}
	return userData, err
}

func (s *DefaultUserDataService) SaveUserData(userData models.UserData) error {
	err := utils.WriteUserDataToJSONFile(&userData, s.UserDataPath)
	if err != nil {
		return err
	}
	return nil
}
