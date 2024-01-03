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
	// Food operations
	GetFoodData(username string) ([]models.FoodData, error)
	AddFoodData(username string, food models.FoodData) error
	UpdateFoodData(username string, food models.FoodData) error
	DeleteFoodData(username string, food models.FoodData) error
	// Dish operations
	GetDishData(username string) ([]models.DishData, error)
	AddDishData(username string, dish models.DishData) error
	UpdateDishData(username string, dish models.DishData) error
	DeleteDishData(username string, dish models.DishData) error
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
