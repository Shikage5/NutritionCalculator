package userData

import (
	"NutritionCalculator/data/models"
	"errors"
)

var ErrDishAlreadyExists = errors.New("dish already exists")
var ErrDishNotFound = errors.New("dish not found")

func (s *DefaultUserDataService) GetDishData(username string) ([]models.DishData, error) {
	savedData, err := s.GetUserData(username)
	if err != nil {
		return nil, err
	}
	return savedData.DishData, nil

}

// AddDish adds a dish to the user's data
func (s *DefaultUserDataService) AddDishData(username string, dishData models.DishData) error {
	savedData, err := s.GetUserData(username)
	if err != nil {
		return err
	}
	for _, f := range savedData.DishData {
		if f.Name == dishData.Name {
			return ErrDishAlreadyExists
		}
	}
	savedData.DishData = append(savedData.DishData, dishData)
	return s.SaveUserData(savedData, username)
}

func (s *DefaultUserDataService) UpdateDishData(username string, dishData models.DishData) error {
	savedData, err := s.GetUserData(username)
	if err != nil {
		return err
	}
	for i, f := range savedData.DishData {
		if f.Name == dishData.Name {
			savedData.DishData[i] = dishData
			break
		} else if i == len(savedData.DishData)-1 {
			return ErrDishNotFound
		}
	}
	return s.SaveUserData(savedData, username)
}

func (s *DefaultUserDataService) DeleteDishData(username string, dishData models.DishData) error {
	savedData, err := s.GetUserData(username)
	if err != nil {
		return err
	}
	for i, f := range savedData.DishData {
		if f.Name == dishData.Name {
			savedData.DishData = append(savedData.DishData[:i], savedData.DishData[i+1:]...)
			break
		} else if i == len(savedData.DishData)-1 {
			return ErrDishNotFound
		}
	}
	return s.SaveUserData(savedData, username)
}

func (s *DefaultUserDataService) GetDishDataByName(username, name string) (models.DishData, error) {
	savedData, err := s.GetUserData(username)
	if err != nil {
		return models.DishData{}, err
	}
	for _, f := range savedData.DishData {
		if f.Name == name {
			return f, nil
		}
	}
	return models.DishData{}, ErrDishNotFound
}
