package models

import (
	"encoding/json"
	"os"
)

type User struct {
	Username  string
	Foods     []Food
	Dishes    []Dish
	Meals     []Meal
	FoodDiary []FoodDiary
}

func (u *User) SaveToFile(filename string) error {
	data, err := json.Marshal(u)
	if err != nil {
		return err
	}

	return os.WriteFile(filename, data, 0644)
}

func (u *User) LoadFromFile(filename string) error {
	data, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	return json.Unmarshal(data, u)
}
