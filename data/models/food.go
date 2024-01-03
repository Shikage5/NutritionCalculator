package models

type Food struct {
	Name              string            `json:"name"`
	Quantity          float64           `json:"quantity"`
	Weight            float64           `json:"weight"`
	NutritionalValues NutritionalValues `json:"nutritionalValues"`
}

type FoodData struct {
	Name              string            `json:"name"`
	NutritionalValues NutritionalValues `json:"nutritionalValues"`
	ReferenceWeight   float64           `json:"referenceWeight"`
	MeasurementUnit   MeasurementUnit   `json:"measurementUnit"`
}
type NutritionalValues struct {
	Energy              float64 `json:"energy"`
	Fat                 float64 `json:"fat"`
	SaturatedFattyAcids float64 `json:"saturatedFattyAcids"`
	Carbohydrates       float64 `json:"carbohydrates"`
	Sugar               float64 `json:"sugar"`
	Protein             float64 `json:"protein"`
	Salt                float64 `json:"salt"`
	Fiber               float64 `json:"fiber"`
	Water               float64 `json:"water"`
}

type MeasurementUnit struct {
	Name   string  `json:"name"`
	Weight float64 `json:"weight"`
}
