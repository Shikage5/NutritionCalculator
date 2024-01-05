package models

type UserData struct {
	Username string     `json:"username"`
	FoodData []FoodData `json:"foodData"`
	DishData []DishData `json:"dishData"`
	Meals    []Meal     `json:"meals"`
	Days     []Day      `json:"foodDiary"`
}
