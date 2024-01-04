package models

type Meal struct {
	Name              string            `json:"name"`
	Foods             []Food            `json:"foods"`
	Dishes            []Dish            `json:"dishes"`
	NutritionalValues NutritionalValues `json:"nutritionalValues"`
}
