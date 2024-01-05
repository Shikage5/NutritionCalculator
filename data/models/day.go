package models

type Day struct {
	Date              string             `json:"date"`
	Meals             []Meal             `json:"meals"`
	NutritionalValues *NutritionalValues `json:"nutritionalValues,omitempty"`
}
