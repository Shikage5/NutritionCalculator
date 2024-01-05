package userData

import (
	"NutritionCalculator/data/models"
	"errors"
)

var ErrFoodAlreadyExists = errors.New("food already exists")
var ErrFoodNotFound = errors.New("food not found")

/*==========================CRUD=============================*/
func (s *DefaultUserDataService) GetFoodData() ([]models.FoodData, error) {
	savedData, err := s.GetUserData()
	if err != nil {
		return nil, err
	}
	return savedData.FoodData, nil
}

func (s *DefaultUserDataService) AddFoodData(foodData models.FoodData) error {
	savedData, err := s.GetUserData()
	if err != nil {
		return err
	}
	for _, f := range savedData.FoodData {
		if f.Name == foodData.Name {
			return ErrFoodAlreadyExists
		}
	}
	savedData.FoodData = append(savedData.FoodData, foodData)
	return s.SaveUserData(savedData)
}

func (s *DefaultUserDataService) UpdateFoodData(foodData models.FoodData) error {
	savedData, err := s.GetUserData()
	if err != nil {
		return err
	}
	for i, f := range savedData.FoodData {
		if f.Name == foodData.Name {
			savedData.FoodData[i] = foodData
			break
		} else if i == len(savedData.FoodData)-1 {
			return ErrFoodNotFound
		}
	}
	return s.SaveUserData(savedData)
}

func (s *DefaultUserDataService) DeleteFoodData(foodData models.FoodData) error {
	savedData, err := s.GetUserData()
	if err != nil {
		return err
	}
	for i, f := range savedData.FoodData {
		if f.Name == foodData.Name {
			savedData.FoodData = append(savedData.FoodData[:i], savedData.FoodData[i+1:]...)
			break
		} else if i == len(savedData.FoodData)-1 {
			return ErrFoodNotFound
		}
	}
	return s.SaveUserData(savedData)
}

/*==========================Specific FoodData OP=============================*/
func (s *DefaultUserDataService) GetFoodDataByName(foodName string) (models.FoodData, error) {
	savedData, err := s.GetUserData()
	if err != nil {
		return models.FoodData{}, err
	}
	for _, f := range savedData.FoodData {
		if f.Name == foodName {
			return f, nil
		}
	}
	return models.FoodData{}, ErrFoodNotFound
}
