package userData

import (
	"NutritionCalculator/data/models"
	"errors"
)

var ErrDishAlreadyExists = errors.New("dish already exists")
var ErrDishNotFound = errors.New("dish not found")

func (s *DefaultUserDataService) GetDishData() ([]models.DishData, error) {
	savedData, err := s.GetUserData()
	if err != nil {
		return nil, err
	}
	return savedData.DishData, nil

}

// AddDish adds a dish to the user's data
func (s *DefaultUserDataService) AddDishData(dishData models.DishData) error {
	savedData, err := s.GetUserData()
	if err != nil {
		return err
	}
	for _, f := range savedData.DishData {
		if f.Name == dishData.Name {
			return ErrDishAlreadyExists
		}
	}
	savedData.DishData = append(savedData.DishData, dishData)
	return s.SaveUserData(savedData)
}

func (s *DefaultUserDataService) UpdateDishData(dishData models.DishData) error {
	savedData, err := s.GetUserData()
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
	return s.SaveUserData(savedData)
}

func (s *DefaultUserDataService) DeleteDishData(dishData models.DishData) error {
	savedData, err := s.GetUserData()
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
	return s.SaveUserData(savedData)
}

func (s *DefaultUserDataService) GetDishDataByName(name string) (models.DishData, error) {
	savedData, err := s.GetUserData()
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

func (s *DefaultUserDataService) CalculateDishDataNutritionalValues(dishData models.DishData) (models.NutritionalValues, error) {
	var totalDishNutritionalValues models.NutritionalValues

	/*==========================Add Nutritional Values of all Foods=============================*/
	for i, food := range dishData.Foods {
		foodNutritionalValues, err := s.CalculateFoodNutritionalValues(food)
		if err != nil {
			return models.NutritionalValues{}, err
		}
		foodWeight, err := s.CalculateFoodWeight(food)
		if err != nil {
			return models.NutritionalValues{}, err
		}
		dishData.Foods[i].NutritionalValues = &foodNutritionalValues
		dishData.Foods[i].Weight = &foodWeight

		totalDishNutritionalValues = s.AddNutritions(totalDishNutritionalValues, foodNutritionalValues)

	}

	/*==========================Add Nutritional Values of all Dishes=============================*/

	for i, dish := range dishData.Dishes {
		dishNutritionalValues, err := s.CalculateDishNutritionalValues(dish, make(map[string]bool))
		if err != nil {
			return models.NutritionalValues{}, err
		}
		dishWeight, err := s.CalculateDishWeight(dish, make(map[string]bool))
		if err != nil {
			return models.NutritionalValues{}, err
		}
		dishData.Dishes[i].NutritionalValues = &dishNutritionalValues
		dishData.Dishes[i].Weight = &dishWeight

		totalDishNutritionalValues = s.AddNutritions(totalDishNutritionalValues, dishNutritionalValues)
	}

	return totalDishNutritionalValues, nil
}
