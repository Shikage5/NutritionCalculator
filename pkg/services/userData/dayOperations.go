package userData

import (
	"NutritionCalculator/data/models"
	"errors"
)

var ErrDayAlreadyExists = errors.New("day already exists")
var ErrDayNotFound = errors.New("day not found")

func (s *DefaultUserDataService) GetDays() ([]models.Day, error) {
	savedData, err := s.GetUserData()
	if err != nil {
		return nil, err
	}
	return savedData.Days, nil
}

func (s *DefaultUserDataService) AddDay(day models.Day) error {
	savedData, err := s.GetUserData()
	if err != nil {
		return err
	}
	for _, d := range savedData.Days {
		if d.Date == day.Date {
			return ErrDayAlreadyExists
		}
	}
	savedData.Days = append(savedData.Days, day)
	return s.SaveUserData(savedData)
}

func (s *DefaultUserDataService) UpdateDay(day models.Day) error {
	savedData, err := s.GetUserData()
	if err != nil {
		return err
	}
	for i, d := range savedData.Days {
		if d.Date == day.Date {
			savedData.Days[i] = day
			break
		} else if i == len(savedData.Days)-1 {
			return ErrDayNotFound
		}
	}
	return s.SaveUserData(savedData)
}

func (s *DefaultUserDataService) DeleteDay(day models.Day) error {
	savedData, err := s.GetUserData()
	if err != nil {
		return err
	}
	for i, d := range savedData.Days {
		if d.Date == day.Date {
			savedData.Days = append(savedData.Days[:i], savedData.Days[i+1:]...)
			break
		} else if i == len(savedData.Days)-1 {
			return ErrDayNotFound
		}
	}
	return s.SaveUserData(savedData)
}

// func (s *DefaultUserDataService) CalculateDayNutritionalValues(day models.Day, processedDishes map[string]bool) (models.NutritionalValues, error){
// 	var dayNutritionalValues models.NutritionalValues
// 	for _, meal := range day.Meals {
// 		mealNutritionalValues, err := s.CalculateMealNutritionalValues(meal, processedDishes)
// 		if err != nil {
// 			return models.NutritionalValues{}, err
// 		}
// 		dayNutritionalValues = dayNutritionalValues.Add(mealNutritionalValues)
// 	}
// 	return dayNutritionalValues, nil
// }
