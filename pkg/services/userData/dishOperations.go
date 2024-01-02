package userData

import (
	"NutritionCalculator/data/models"
	"errors"
)

var ErrDishAlreadyExists = errors.New("dish already exists")
var ErrDishNotFound = errors.New("dish not found")

func (s *DefaultUserDataService) GetDish(username string) ([]models.Dish, error) {
	savedData, err := s.GetUserData(username)
	if err != nil {
		return nil, err
	}
	return savedData.Dishes, nil
}

// AddDish adds a dish to the user's data
func (s *DefaultUserDataService) AddDish(username string, dish models.Dish) error {
	savedData, err := s.GetUserData(username)
	if err != nil {
		return err
	}
	for _, f := range savedData.Dishes {
		if f.Name == dish.Name {
			return ErrDishAlreadyExists
		}
	}
	savedData.Dishes = append(savedData.Dishes, dish)
	return s.SaveUserData(savedData, username)
}

func (s *DefaultUserDataService) UpdateDish(username string, dish models.Dish) error {
	savedData, err := s.GetUserData(username)
	if err != nil {
		return err
	}
	for i, f := range savedData.Dishes {
		if f.Name == dish.Name {
			savedData.Dishes[i] = dish
			break
		} else if i == len(savedData.Dishes)-1 {
			return ErrDishNotFound
		}
	}
	return s.SaveUserData(savedData, username)
}

func (s *DefaultUserDataService) DeleteDish(username string, dish models.Dish) error {
	savedData, err := s.GetUserData(username)
	if err != nil {
		return err
	}
	for i, f := range savedData.Dishes {
		if f.Name == dish.Name {
			savedData.Dishes = append(savedData.Dishes[:i], savedData.Dishes[i+1:]...)
			break
		} else if i == len(savedData.Dishes)-1 {
			return ErrDishNotFound
		}
	}
	return s.SaveUserData(savedData, username)
}
