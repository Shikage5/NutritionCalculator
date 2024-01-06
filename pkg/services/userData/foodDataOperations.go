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

func (s *DefaultUserDataService) UpdateFoodData(newFoodData models.FoodData) error {
	savedData, err := s.GetUserData()
	if err != nil {
		return err
	}

	for i, f := range savedData.FoodData {
		if f.Name == newFoodData.Name {
			savedData.FoodData[i] = newFoodData
			break
		} else if i == len(savedData.FoodData)-1 {
			return ErrFoodNotFound
		}
	}
	//save the updated FoodData
	err = s.SaveUserData(savedData)
	if err != nil {
		return err
	}

	// Recalculate the nutritional values of all food items, dishes and meals
	savedData, err = s.recalculateNutritionalValuesOfFoods(newFoodData.Name, savedData)
	if err != nil {
		return err
	}
	savedData.DishData, err = s.recalculateNutritionalValuesOfDishes(savedData.DishData)
	if err != nil {
		return err
	}
	savedData.Meals, err = s.recalculateNutritionalValuesOfMeals(savedData.Meals)
	if err != nil {
		return err
	}
	savedData.Days, err = s.recalculateNutritionalValuesOfDays(savedData.Days)
	if err != nil {
		return err
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

/*==========================Recalculate Helper Functions=============================*/
func (s *DefaultUserDataService) recalculateNutritionalValuesOfFoods(foodName string, savedData models.UserData) (models.UserData, error) {
	for i, dish := range savedData.DishData {
		for j, food := range dish.Foods {
			if food.Name == foodName {
				foodNutritionalValues, err := s.CalculateFoodNutritionalValues(food)
				if err != nil {
					return models.UserData{}, err
				}
				savedData.DishData[i].Foods[j].NutritionalValues = &foodNutritionalValues
				break
			}
		}
	}
	for i, meal := range savedData.Meals {
		for j, food := range meal.Foods {
			if food.Name == foodName {
				foodNutritionalValues, err := s.CalculateFoodNutritionalValues(food)
				if err != nil {
					return models.UserData{}, err
				}
				savedData.Meals[i].Foods[j].NutritionalValues = &foodNutritionalValues
				break
			}
		}
	}
	return savedData, nil
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
