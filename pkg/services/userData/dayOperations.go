package userData

import (
	"NutritionCalculator/data/models"
	"errors"
)

type DayService interface {
	AddDay(username string, day models.Day) error
	GetDays(username string) ([]models.Day, error)
	UpdateDay(username string, day models.Day) error
	DeleteDay(username string, day models.Day) error
	GetDayByDate(username string, date string) (models.Day, error)
	GetLastSevenDays(username string) ([]models.Day, error)
	RecalculateNutritionalValuesOfDays([]models.Day) ([]models.Day, error)
	CalculateDayNutritionalValues(models.Day) (models.NutritionalValues, error)
	CalculateLastSevenDaysNutritionalValues(string) (models.NutritionalValues, error)
}

type DefaultDayService struct {
	UserDataService        UserDataService
	MealService            MealService
	NutritionValuesService NutritionValuesService
}

var ErrDayAlreadyExists = errors.New("day already exists")
var ErrDayNotFound = errors.New("day not found")

func NewDayService(userDataService UserDataService, mealService MealService, nutritionValuesService NutritionValuesService) *DefaultDayService {
	return &DefaultDayService{UserDataService: userDataService, MealService: mealService, NutritionValuesService: nutritionValuesService}
}

/*==========================CRUD=============================*/
func (s *DefaultDayService) AddDay(username string, day models.Day) error {
	savedData, err := s.UserDataService.GetUserData(username)
	if err != nil {
		return err
	}
	for _, d := range savedData.Days {
		if d.Date == day.Date {
			return ErrDayAlreadyExists
		}
	}
	savedData.Days = append(savedData.Days, day)
	return s.UserDataService.SaveUserData(savedData)
}
func (s *DefaultDayService) GetDays(username string) ([]models.Day, error) {
	savedData, err := s.UserDataService.GetUserData(username)
	if err != nil {
		return nil, err
	}
	return savedData.Days, nil
}

func (s *DefaultDayService) UpdateDay(username string, day models.Day) error {
	savedData, err := s.UserDataService.GetUserData(username)
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

	return s.UserDataService.SaveUserData(savedData)
}

func (s *DefaultDayService) DeleteDay(username string, day models.Day) error {
	savedData, err := s.UserDataService.GetUserData(username)
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
	return s.UserDataService.SaveUserData(savedData)
}

/*==========================Delete Helper Fucntions=============================*/

func (s *DefaultDayService) RecalculateNutritionalValuesOfDays(days []models.Day) ([]models.Day, error) {
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

func (s *DefaultDayService) GetDayByDate(username string, date string) (models.Day, error) {
	savedData, err := s.UserDataService.GetUserData(username)
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

func (s *DefaultDayService) GetLastSevenDays(username string) ([]models.Day, error) {
	savedData, err := s.UserDataService.GetUserData(username)
	if err != nil {
		return nil, err
	}
	if len(savedData.Days) < 7 {
		return savedData.Days, nil
	}
	return savedData.Days[len(savedData.Days)-7:], nil
}

/*==========================Day Nutritional Values=============================*/
func (s *DefaultDayService) CalculateDayNutritionalValues(day models.Day) (models.NutritionalValues, error) {
	var dayNutritionalValues models.NutritionalValues
	for _, meal := range day.Meals {
		mealNutritionalValues, err := s.MealService.CalculateMealNutritionalValues(meal)
		if err != nil {
			return models.NutritionalValues{}, err
		}
		dayNutritionalValues = s.NutritionValuesService.AddNutritions(dayNutritionalValues, mealNutritionalValues)
	}
	return dayNutritionalValues, nil
}

func (s *DefaultDayService) CalculateLastSevenDaysNutritionalValues(username string) (models.NutritionalValues, error) {
	var totalNutritionalValues models.NutritionalValues
	days, err := s.GetLastSevenDays(username)
	if err != nil {
		return models.NutritionalValues{}, err
	}
	for _, day := range days {
		dayNutritionalValues, err := s.CalculateDayNutritionalValues(day)
		if err != nil {
			return models.NutritionalValues{}, err
		}
		totalNutritionalValues = s.NutritionValuesService.AddNutritions(totalNutritionalValues, dayNutritionalValues)
	}
	return totalNutritionalValues, nil
}
