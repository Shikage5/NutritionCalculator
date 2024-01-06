package utils

import (
	"NutritionCalculator/data/models"
	"encoding/csv"
	"fmt"
	"io"
	"sort"
	"time"
)

func WriteDaysToCSV(w io.Writer, days []models.Day) error {
	// Sort days by date
	sort.Slice(days, func(i, j int) bool {
		dateI, _ := time.Parse("2006-01-02", days[i].Date)
		dateJ, _ := time.Parse("2006-01-02", days[j].Date)
		return dateI.Before(dateJ)
	})

	writer := csv.NewWriter(w)
	defer writer.Flush()

	//write header
	err := writer.Write([]string{"Date", "Energy", "Fat", "Saturated Fatty Acids", "Carbohydrates", "Sugar", "Protein", "Salt", "Fiber", "Water"})
	if err != nil {
		return err
	}

	for _, day := range days {
		err = writer.Write([]string{
			day.Date,
			fmt.Sprintf("%f", day.NutritionalValues.Energy),
			fmt.Sprintf("%f", day.NutritionalValues.Fat),
			fmt.Sprintf("%f", day.NutritionalValues.SaturatedFattyAcids),
			fmt.Sprintf("%f", day.NutritionalValues.Carbohydrates),
			fmt.Sprintf("%f", day.NutritionalValues.Sugar),
			fmt.Sprintf("%f", day.NutritionalValues.Protein),
			fmt.Sprintf("%f", day.NutritionalValues.Salt),
			fmt.Sprintf("%f", day.NutritionalValues.Fiber),
			fmt.Sprintf("%f", day.NutritionalValues.Water),
		})
		if err != nil {
			return err
		}
	}
	return nil
}
