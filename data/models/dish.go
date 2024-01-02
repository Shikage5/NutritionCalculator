package models

type Dish struct {
	Name  string  `json:"name"`
	Foods []Food  `json:"foods"`
	Dish  []*Dish `json:"dish"`
}
