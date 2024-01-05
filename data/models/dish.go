package models

type Dish struct {
	Name              string             `json:"name"`
	Quantity          *float64           `json:"quantity"`
	Weight            *float64           `json:"weight"`
	NutritionalValues *NutritionalValues `json:"nutritionalValues,omitempty"`
}

type DishData struct {
	Name              string             `json:"name"`
	Foods             []Food             `json:"foods"`
	Dishes            []Dish             `json:"dishes"`
	NutritionalValues *NutritionalValues `json:"nutritionalValues,omitempty"`
}
