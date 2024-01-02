package models

type Day struct {
	Date   string `json:"date"`
	Foods  []Food `json:"foods"`
	Dishes []Dish `json:"dishes"`
	Meals  []Meal `json:"meals"`
}
