package services

import (
	"NutritionCalculator/data/models"
)

type UserDataService interface {
	AddFoodToUser(username string, food models.Food) error
}

type DefaultUserDataService struct {
	UserDataPath string
}

func (s *DefaultUserDataService) AddFoodToUser(username string, food models.Food) error {

}
