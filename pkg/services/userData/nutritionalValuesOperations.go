package userData

import "NutritionCalculator/data/models"

// AddNutritionsByRatio adds two nutritional values by ratio
func (s *DefaultUserDataService) AddNutritionsByRatio(ratio float64, nutritionalValues models.NutritionalValues) models.NutritionalValues {
	var result models.NutritionalValues
	result.Carbohydrates = nutritionalValues.Carbohydrates * ratio
	result.Energy = nutritionalValues.Energy * ratio
	result.Fat = nutritionalValues.Fat * ratio
	result.Fiber = nutritionalValues.Fiber * ratio
	result.Protein = nutritionalValues.Protein * ratio
	result.Salt = nutritionalValues.Salt * ratio
	result.SaturatedFattyAcids = nutritionalValues.SaturatedFattyAcids * ratio
	result.Sugar = nutritionalValues.Sugar * ratio
	result.Water = nutritionalValues.Water * ratio
	return result
}

// AddNutritions adds two nutritional values
func (s *DefaultUserDataService) AddNutritions(n1, n2 models.NutritionalValues) models.NutritionalValues {
	var result models.NutritionalValues
	result.Carbohydrates = n1.Carbohydrates + n2.Carbohydrates
	result.Energy = n1.Energy + n2.Energy
	result.Fat = n1.Fat + n2.Fat
	result.Fiber = n1.Fiber + n2.Fiber
	result.Protein = n1.Protein + n2.Protein
	result.Salt = n1.Salt + n2.Salt
	result.SaturatedFattyAcids = n1.SaturatedFattyAcids + n2.SaturatedFattyAcids
	result.Sugar = n1.Sugar + n2.Sugar
	result.Water = n1.Water + n2.Water
	return result
}
