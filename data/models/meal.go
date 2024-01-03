package models

type Meal struct {
	Name              string            `json:"name"`
	Foods             []Food            `json:"foods"`
	Dish              []Dish            `json:"dish"`
	NutritionalValues NutritionalValues `json:"nutritionalValues"`
}
