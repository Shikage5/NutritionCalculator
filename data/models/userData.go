package models

type UserData struct {
	Username  string     `json:"username"`
	FoodData  []FoodData `json:"foods,omitempty"`
	DishData  []DishData `json:"dishes,omitempty"`
	Meals     []Meal     `json:"meals,omitempty"`
	FoodDiary []Day      `json:"foodDiary"`
}
