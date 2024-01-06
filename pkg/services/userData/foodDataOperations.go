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

			// Delete the food from all dishes and meals
			savedData.DishData = s.deleteFoodFromDishes(foodData.Name, savedData.DishData)
			savedData.Meals = s.deleteFoodFromMeals(foodData.Name, savedData.Meals)

			// Recalculate the nutritional values of all dishes and meals
			savedData.DishData, err = s.recalculateNutritionalValuesOfDishes(savedData.DishData)
			if err != nil {
				return err
			}
			savedData.Meals, err = s.recalculateNutritionalValuesOfMeals(savedData.Meals)
			if err != nil {
				return err
			}

			break
		} else if i == len(savedData.FoodData)-1 {
			return ErrFoodNotFound
		}
	}
	return s.SaveUserData(savedData)
}

/*==========================Delete Helper Functions=============================*/
func (s *DefaultUserDataService) deleteFoodFromDishes(foodName string, dishes []models.DishData) []models.DishData {
	for i, dish := range dishes {
		for j, food := range dish.Foods {
			if food.Name == foodName {
				dishes[i].Foods = append(dish.Foods[:j], dish.Foods[j+1:]...)
				break
			}
		}
	}
	return dishes
}

func (s *DefaultUserDataService) deleteFoodFromMeals(foodName string, meals []models.Meal) []models.Meal {
	for i, meal := range meals {
		for j, food := range meal.Foods {
			if food.Name == foodName {
				meals[i].Foods = append(meal.Foods[:j], meal.Foods[j+1:]...)
				break
			}
		}
	}
	return meals
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
