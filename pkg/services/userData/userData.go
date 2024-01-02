package userData

import (
	"NutritionCalculator/data/models"
)

// UserDataService handles the logic for getting and saving user data

type UserDataService interface {
	GetUserData(username string) (*models.UserData, error)
	SaveUserData(data *models.UserData, username string) error
	// Food operations
	GetFoods(username string) ([]models.Food, error)
	AddFood(username string, food models.Food) error
	UpdateFood(username string, food models.Food) error
	DeleteFood(username string, food models.Food) error
}

type DefaultUserDataService struct {
	UserDataPath string
}

func (s *DefaultUserDataService) GetUserData(username string) (*models.UserData, error) {
	data := &models.UserData{}

	err := data.LoadFromFile(s.UserDataPath)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (s *DefaultUserDataService) SaveUserData(data *models.UserData, username string) error {
	return data.SaveToFile(s.UserDataPath)
}
