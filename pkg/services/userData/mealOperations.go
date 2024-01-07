package userData

import (
	"NutritionCalculator/data/models"
	"errors"
)

type MealService interface {
	AddMeal(username string, meal models.Meal) error
	GetMeals(username string) ([]models.Meal, error)
	UpdateMeal(username string, meal models.Meal) error
	DeleteMeal(username string, meal models.Meal) error
	RecalculateNutritionalValuesOfMeals([]models.Meal) ([]models.Meal, error)
	CalculateMealNutritionalValues(models.Meal) (models.NutritionalValues, error)
}

type DefaultMealService struct {
	UserDataService        UserDataService
	FoodService            FoodService
	DishService            DishService
	DayService             DayService
	NutritionValuesService NutritionValuesService
}

var ErrMealAlreadyExists = errors.New("meal already exists")
var ErrMealNotFound = errors.New("meal not found")

func NewMealService(userDataService UserDataService, foodService FoodService, dishService DishService, dayService DayService, nutritionValuesService NutritionValuesService) *DefaultMealService {
	return &DefaultMealService{UserDataService: userDataService, FoodService: foodService, DishService: dishService, DayService: dayService, NutritionValuesService: nutritionValuesService}
}

/*==========================CRUD=============================*/
func (s *DefaultMealService) AddMeal(username string, meal models.Meal) error {
	savedData, err := s.UserDataService.GetUserData(username)
	if err != nil {
		return err
	}
	for _, m := range savedData.Meals {
		if m.Name == meal.Name {
			return ErrMealAlreadyExists
		}
	}
	savedData.Meals = append(savedData.Meals, meal)
	return s.UserDataService.SaveUserData(savedData)
}

func (s *DefaultMealService) GetMeals(username string) ([]models.Meal, error) {
	savedData, err := s.UserDataService.GetUserData(username)
	if err != nil {
		return nil, err
	}
	return savedData.Meals, nil
}

func (s *DefaultMealService) UpdateMeal(username string, meal models.Meal) error {
	savedData, err := s.UserDataService.GetUserData(username)
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
	savedData.Days, err = s.DayService.RecalculateNutritionalValuesOfDays(savedData.Days)
	if err != nil {
		return err
	}
	return s.UserDataService.SaveUserData(savedData)

}

func (s *DefaultMealService) DeleteMeal(username string, meal models.Meal) error {
	savedData, err := s.UserDataService.GetUserData(username)
	if err != nil {
		return err
	}
	for i, m := range savedData.Meals {
		if m.Name == meal.Name {
			savedData.Meals = append(savedData.Meals[:i], savedData.Meals[i+1:]...)

			// Delete the meal from all days
			savedData.Days = s.deleteMealFromDays(meal.Name, savedData.Days)

			//recalculate nutritional values of all meals
			savedData.Days, err = s.DayService.RecalculateNutritionalValuesOfDays(savedData.Days)
			if err != nil {
				return err
			}

			break
		} else if i == len(savedData.Meals)-1 {
			return ErrMealNotFound
		}
	}
	return s.UserDataService.SaveUserData(savedData)
}

/* ========================== Delete Helper Function============================= */
func (s *DefaultMealService) deleteMealFromDays(mealName string, days []models.Day) []models.Day {
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

func (s *DefaultMealService) RecalculateNutritionalValuesOfMeals(meals []models.Meal) ([]models.Meal, error) {
	for i, meal := range meals {
		nutritionalValues, err := s.CalculateMealNutritionalValues(username ,meal)
		if err != nil {
			return nil, err
		}
		meals[i].NutritionalValues = &nutritionalValues
	}
	return meals, nil
}

/*==========================Nutritional Values=============================*/
func (s *DefaultMealService) CalculateMealNutritionalValues(username, meal models.Meal) (models.NutritionalValues, error) {
	var totalMealNutritionalValues models.NutritionalValues

	// Add Nutritional Values of all Foods

	for i, food := range meal.Foods {
		foodNutritionalValues, err := s.FoodService.CalculateFoodNutritionalValues(username, food)
		if err != nil {
			return models.NutritionalValues{}, err
		}

		foodWeight, err := s.FoodService.CalculateFoodWeight(food)
		if err != nil {
			return models.NutritionalValues{}, err
		}
		meal.Foods[i].NutritionalValues = &foodNutritionalValues
		meal.Foods[i].Weight = &foodWeight

		totalMealNutritionalValues = s.NutritionValuesService.AddNutritions(totalMealNutritionalValues, foodNutritionalValues)
	}

	// Add Nutritional Values of all Dishes
	for i, dish := range meal.Dishes {

		dishNutritionalValues, err := s.DishService.CalculateDishNutritionalValues(dish, make(map[string]bool))
		if err != nil {
			return models.NutritionalValues{}, err
		}

		foodWeight, err := s.DishService.CalculateDishWeight(dish, make(map[string]bool))
		if err != nil {
			return models.NutritionalValues{}, err
		}
		meal.Dishes[i].NutritionalValues = &dishNutritionalValues
		meal.Dishes[i].Weight = &foodWeight

		totalMealNutritionalValues = s.NutritionValuesService.AddNutritions(totalMealNutritionalValues, dishNutritionalValues)
	}

	return totalMealNutritionalValues, nil
}
