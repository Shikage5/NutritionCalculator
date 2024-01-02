package userData

import (
	"NutritionCalculator/data/models"
	"errors"
)

var ErrFoodAlreadyExists = errors.New("food already exists")
var ErrFoodNotFound = errors.New("food not found")

func (s *DefaultUserDataService) GetFoods(username string) ([]models.Food, error) {
	savedData, err := s.GetUserData(username)
	if err != nil {
		return nil, err
	}
	return savedData.Foods, nil
}

// AddFood adds a food to the user's data
func (s *DefaultUserDataService) AddFood(username string, food models.Food) error {
	savedData, err := s.GetUserData(username)
	if err != nil {
		return err
	}
	for _, f := range savedData.Foods {
		if f.Name == food.Name {
			return ErrFoodAlreadyExists
		}
	}
	savedData.Foods = append(savedData.Foods, food)
	return s.SaveUserData(savedData, username)
}

func (s *DefaultUserDataService) UpdateFood(username string, food models.Food) error {
	savedData, err := s.GetUserData(username)
	if err != nil {
		return err
	}
	for i, f := range savedData.Foods {
		if f.Name == food.Name {
			savedData.Foods[i] = food
			break
		} else if i == len(savedData.Foods)-1 {
			return ErrFoodNotFound
		}
	}
	return s.SaveUserData(savedData, username)
}

func (s *DefaultUserDataService) DeleteFood(username string, food models.Food) error {
	savedData, err := s.GetUserData(username)
	if err != nil {
		return err
	}
	for i, f := range savedData.Foods {
		if f.Name == food.Name {
			savedData.Foods = append(savedData.Foods[:i], savedData.Foods[i+1:]...)
			break
		} else if i == len(savedData.Foods)-1 {
			return ErrFoodNotFound
		}
	}
	return s.SaveUserData(savedData, username)
}
