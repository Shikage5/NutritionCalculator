package userData

import (
	"NutritionCalculator/data/models"
	"errors"
)

var ErrMealAlreadyExists = errors.New("meal already exists")
var ErrMealNotFound = errors.New("meal not found")

/*==========================CRUD=============================*/
func (s *DefaultUserDataService) AddMeal(meal models.Meal) error {
	savedData, err := s.GetUserData()
	if err != nil {
		return err
	}
	for _, m := range savedData.Meals {
		if m.Name == meal.Name {
			return ErrMealAlreadyExists
		}
	}
	savedData.Meals = append(savedData.Meals, meal)
	return s.SaveUserData(savedData)
}

func (s *DefaultUserDataService) GetMeals() ([]models.Meal, error) {
	savedData, err := s.GetUserData()
	if err != nil {
		return nil, err
	}
	return savedData.Meals, nil
}

func (s *DefaultUserDataService) UpdateMeal(meal models.Meal) error {
	savedData, err := s.GetUserData()
	if err != nil {
		return err
	}
	for i, m := range savedData.Meals {
		if m.Name == meal.Name {
			savedData.Meals[i] = meal
			break
		} else if i == len(savedData.Meals)-1 {
			return ErrMealNotFound
		}
	}
	return s.SaveUserData(savedData)
}

func (s *DefaultUserDataService) DeleteMeal(meal models.Meal) error {
	savedData, err := s.GetUserData()
	if err != nil {
		return err
	}
	for i, m := range savedData.Meals {
		if m.Name == meal.Name {
			savedData.Meals = append(savedData.Meals[:i], savedData.Meals[i+1:]...)

			// Delete the meal from all days
			savedData.Days = s.deleteMealFromDays(meal.Name, savedData.Days)

			//recalculate nutritional values of all meals
			savedData.Days, err = s.recalculateNutritionalValuesOfDays(savedData.Days)
			if err != nil {
				return err
			}

			break
		} else if i == len(savedData.Meals)-1 {
			return ErrMealNotFound
		}
	}
	return s.SaveUserData(savedData)
}

/* ========================== Delete Helper Function============================= */
func (s *DefaultUserDataService) deleteMealFromDays(mealName string, days []models.Day) []models.Day {
	for i, day := range days {
		for j, meal := range day.Meals {
			if meal.Name == mealName {
				days[i].Meals = append(day.Meals[:j], day.Meals[j+1:]...)
				break
			}
		}
	}
	return days
}

func (s *DefaultUserDataService) recalculateNutritionalValuesOfMeals(meals []models.Meal) ([]models.Meal, error) {
	for i, meal := range meals {
		nutritionalValues, err := s.CalculateMealNutritionalValues(meal)
		if err != nil {
			return nil, err
		}
		meals[i].NutritionalValues = &nutritionalValues
	}
	return meals, nil
}

/*==========================Nutritional Values=============================*/
func (s *DefaultUserDataService) CalculateMealNutritionalValues(meal models.Meal) (models.NutritionalValues, error) {
	var totalMealNutritionalValues models.NutritionalValues

	// Add Nutritional Values of all Foods

	for _, food := range meal.Foods {
		foodNutritionalValues, err := s.CalculateFoodNutritionalValues(food)
		if err != nil {
			return models.NutritionalValues{}, err
		}
		totalMealNutritionalValues = s.AddNutritions(totalMealNutritionalValues, foodNutritionalValues)
	}

	// Add Nutritional Values of all Dishes
	for _, dish := range meal.Dishes {

		dishNutritionalValues, err := s.CalculateDishNutritionalValues(dish, make(map[string]bool))
		if err != nil {
			return models.NutritionalValues{}, err
		}
		totalMealNutritionalValues = s.AddNutritions(totalMealNutritionalValues, dishNutritionalValues)
	}

	return totalMealNutritionalValues, nil
}
