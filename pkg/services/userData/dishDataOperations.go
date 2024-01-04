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

func (s *DefaultUserDataService) CalculateDishDataNutritionalValues(username string, dishData models.DishData) (models.NutritionalValues, error) {
	for _, food := range dishData.Foods {
		foodNutritionalValues, err := s.CalculateFoodNutritionalValues(username, food)
		if err != nil {
			return models.NutritionalValues{}, err
		}
		dishData.NutritionalValues.Carbohydrates += foodNutritionalValues.Carbohydrates
		dishData.NutritionalValues.Energy += foodNutritionalValues.Energy
		dishData.NutritionalValues.Fat += foodNutritionalValues.Fat
		dishData.NutritionalValues.Fiber += foodNutritionalValues.Fiber
		dishData.NutritionalValues.Protein += foodNutritionalValues.Protein
		dishData.NutritionalValues.Salt += foodNutritionalValues.Salt
		dishData.NutritionalValues.SaturatedFattyAcids += foodNutritionalValues.SaturatedFattyAcids
		dishData.NutritionalValues.Sugar += foodNutritionalValues.Sugar
		dishData.NutritionalValues.Water += foodNutritionalValues.Water
	}
	for _, dish := range dishData.Dishes {
		dishNutritionalValues, err := s.CalculateDishNutritionalValues(username, dish, make(map[string]bool))
		if err != nil {
			return models.NutritionalValues{}, err
		}
		dishData.NutritionalValues.Carbohydrates += dishNutritionalValues.Carbohydrates
		dishData.NutritionalValues.Energy += dishNutritionalValues.Energy
		dishData.NutritionalValues.Fat += dishNutritionalValues.Fat
		dishData.NutritionalValues.Fiber += dishNutritionalValues.Fiber
		dishData.NutritionalValues.Protein += dishNutritionalValues.Protein
		dishData.NutritionalValues.Salt += dishNutritionalValues.Salt
		dishData.NutritionalValues.SaturatedFattyAcids += dishNutritionalValues.SaturatedFattyAcids
		dishData.NutritionalValues.Sugar += dishNutritionalValues.Sugar
		dishData.NutritionalValues.Water += dishNutritionalValues.Water
	}

	return dishData.NutritionalValues, nil
}
