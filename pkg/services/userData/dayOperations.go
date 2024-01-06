package userData

import (
	"NutritionCalculator/data/models"
	"errors"
)

var ErrDayAlreadyExists = errors.New("day already exists")
var ErrDayNotFound = errors.New("day not found")

/*==========================CRUD=============================*/
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
func (s *DefaultUserDataService) GetDays() ([]models.Day, error) {
	savedData, err := s.GetUserData()
	if err != nil {
		return nil, err
	}
	return savedData.Days, nil
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

/*==========================Delete Helper Fucntions=============================*/

func (s *DefaultUserDataService) recalculateNutritionalValuesOfDays(days []models.Day) ([]models.Day, error) {
	for i, day := range days {
		nutritionalValues, err := s.CalculateDayNutritionalValues(day)
		if err != nil {
			return nil, err
		}
		days[i].NutritionalValues = &nutritionalValues
	}
	return days, nil
}

/*==========================Specific Days=============================*/

func (s *DefaultUserDataService) GetDayByDate(date string) (models.Day, error) {
	savedData, err := s.GetUserData()
	if err != nil {
		return models.Day{}, err
	}
	for _, day := range savedData.Days {
		if day.Date == date {
			return day, nil
		}
	}
	return models.Day{}, ErrDayNotFound
}

func (s *DefaultUserDataService) GetLastSevenDays() ([]models.Day, error) {
	savedData, err := s.GetUserData()
	if err != nil {
		return nil, err
	}
	if len(savedData.Days) < 7 {
		return savedData.Days, nil
	}
	return savedData.Days[len(savedData.Days)-7:], nil
}

/*==========================Day Nutritional Values=============================*/
func (s *DefaultUserDataService) CalculateDayNutritionalValues(day models.Day) (models.NutritionalValues, error) {
	var dayNutritionalValues models.NutritionalValues
	for _, meal := range day.Meals {
		mealNutritionalValues, err := s.CalculateMealNutritionalValues(meal)
		if err != nil {
			return models.NutritionalValues{}, err
		}
		dayNutritionalValues = s.AddNutritions(dayNutritionalValues, mealNutritionalValues)
	}
	return dayNutritionalValues, nil
}

func (s *DefaultUserDataService) CalculateLastSevenDaysNutritionalValues() (models.NutritionalValues, error) {
	var totalNutritionalValues models.NutritionalValues
	days, err := s.GetLastSevenDays()
	if err != nil {
		return models.NutritionalValues{}, err
	}
	for _, day := range days {
		dayNutritionalValues, err := s.CalculateDayNutritionalValues(day)
		if err != nil {
			return models.NutritionalValues{}, err
		}
		totalNutritionalValues = s.AddNutritions(totalNutritionalValues, dayNutritionalValues)
	}
	return totalNutritionalValues, nil
}
