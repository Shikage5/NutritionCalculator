package models

type UserData struct {
	Username  string     `json:"username"`
	FoodData  []FoodData `json:"foods"`
	DishData  []DishData `json:"dishes"`
	Meals     []Meal     `json:"meals"`
	FoodDiary []Day      `json:"foodDiary"`
}
