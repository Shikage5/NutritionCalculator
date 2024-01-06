package handlers

import (
	"NutritionCalculator/data/models"
	contextkeys "NutritionCalculator/pkg/contextKeys"
	"NutritionCalculator/pkg/services/userData"
	"net/http"
	"time"
)

func HomeHandler(userDataPath string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Get user data based on username from context
		username := r.Context().Value(contextkeys.UserKey).(string)
		userDataService := userData.NewUserDataService(username, userDataPath)
		if r.Method == http.MethodGet {
			userData, err := userDataService.GetUserData()
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			// Check if a Day object for today already exists
			todaysDate := time.Now().Format("2006-01-02")
			today := models.Day{Date: todaysDate}
			err = userDataService.AddDay(today)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			// Save the updated user data
			err = userDataService.SaveUserData(userData)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			// Calculate the nutritional values of the day
			dayNutritionalValues, err := userDataService.CalculateDayNutritionalValues(today)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			// Calculate the average nutritional values of the last seven days
			averageNutritionalValues, err := userDataService.CalculateLastSevenDaysNutritionalValues()
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			// Create a struct to hold the data for the overview page
			overviewData := struct {
				UserData                 models.UserData
				DayNutritionalValues     models.NutritionalValues
				AverageNutritionalValues models.NutritionalValues
			}{
				UserData:                 userData,
				DayNutritionalValues:     dayNutritionalValues,
				AverageNutritionalValues: averageNutritionalValues,
			}

			// Display the overview page with the calculated data
			DisplayPage(w, overviewData, "web/template/home.html")
		}
	}
}
