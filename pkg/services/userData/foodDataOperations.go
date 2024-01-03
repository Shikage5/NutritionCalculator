package userData

import (
	"NutritionCalculator/data/models"
	"errors"
)

var ErrFoodAlreadyExists = errors.New("food already exists")
var ErrFoodNotFound = errors.New("food not found")

// GetFoodData gets the user's food data
func (s *DefaultUserDataService) GetFoodData(username string) ([]models.FoodData, error) {
	savedData, err := s.GetUserData(username)
	if err != nil {
		return nil, err
	}
	return savedData.FoodData, nil
}

// AddFood adds a food to the user's data
func (s *DefaultUserDataService) AddFoodData(username string, foodData models.FoodData) error {
	savedData, err := s.GetUserData(username)
	if err != nil {
		return err
	}
	for _, f := range savedData.FoodData {
		if f.Name == foodData.Name {
			return ErrFoodAlreadyExists
		}
	}
	savedData.FoodData = append(savedData.FoodData, foodData)
	return s.SaveUserData(savedData, username)
}

// UpdateFood updates a food in the user's data
func (s *DefaultUserDataService) UpdateFoodData(username string, food models.FoodData) error {
	savedData, err := s.GetUserData(username)
	if err != nil {
		return err
	}
	for i, f := range savedData.FoodData {
		if f.Name == food.Name {
			savedData.FoodData[i] = food
			break
		} else if i == len(savedData.FoodData)-1 {
			return ErrFoodNotFound
		}
	}
	return s.SaveUserData(savedData, username)
}

// DeleteFood deletes a food from the user's data
func (s *DefaultUserDataService) DeleteFoodData(username string, food models.FoodData) error {
	savedData, err := s.GetUserData(username)
	if err != nil {
		return err
	}
	for i, f := range savedData.FoodData {
		if f.Name == food.Name {
			savedData.FoodData = append(savedData.FoodData[:i], savedData.FoodData[i+1:]...)
			break
		} else if i == len(savedData.FoodData)-1 {
			return ErrFoodNotFound
		}
	}
	return s.SaveUserData(savedData, username)
}
