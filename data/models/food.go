package models

type Food struct {
	Name                string  `json:"name"`
	Energy              float64 `json:"energy"`
	Fat                 float64 `json:"fat"`
	SaturatedFattyAcids float64 `json:"saturatedFattyAcids"`
	Carbohydrates       float64 `json:"carbohydrates"`
	Sugar               float64 `json:"sugar"`
	Protein             float64 `json:"protein"`
	Salt                float64 `json:"salt"`
	Fiber               float64 `json:"fiber"`
	Water               float64 `json:"water"`
	ReferenceWeight     float64 `json:"referenceWeight"`
	UnitWeight          float64 `json:"unitWeight"`
}
