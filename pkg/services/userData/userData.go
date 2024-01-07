package userData

import (
	"NutritionCalculator/data/models"
	"NutritionCalculator/utils"
)

type UserDataService interface {
	//userdata operations
	GetUserData(username string) (models.UserData, error)
	SaveUserData(models.UserData) error
}

//==========================Service for handling user data=============================

type DefaultUserDataService struct {
	UserDataPath string
	Username     string
}

func NewUserDataService(userDataPath string) *DefaultUserDataService {
	return &DefaultUserDataService{
		UserDataPath: userDataPath,
	}
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

func (s *DefaultUserDataService) SaveUserData(userData models.UserData) error {
	err := utils.WriteUserDataToJSONFile(&userData, s.UserDataPath)
	if err != nil {
		return err
	}
	return nil
}
