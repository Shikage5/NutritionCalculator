package models

import (
	"encoding/json"
	"os"
)

type UserData struct {
	Username  string `json:"username"`
	Foods     []Food `json:"foods,omitempty"`
	Dishes    []Dish `json:"dishes,omitempty"`
	Meals     []Meal `json:"meals,omitempty"`
	FoodDiary []Day  `json:"foodDiary"`
}

func (u *UserData) SaveToFile(userDataPath string) error {
	data, err := json.Marshal(u)
	if err != nil {
		return err
	}
	filepath := userDataPath + u.Username + ".json"
	return os.WriteFile(filepath, data, 0644)
}

func (u *UserData) LoadFromFile(userDataPath string) error {
	filepath := userDataPath + u.Username + ".json"
	data, err := os.ReadFile(filepath)
	if err != nil {
		return err
	}

	return json.Unmarshal(data, u)
}
