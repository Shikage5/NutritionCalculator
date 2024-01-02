package models

type Dish struct {
	Name  string
	Foods []Food
	Dish  []*Dish
}
