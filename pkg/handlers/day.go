package handlers

import (
	"NutritionCalculator/data/models"
	contextkeys "NutritionCalculator/pkg/contextKeys"
	"NutritionCalculator/pkg/services/userData"
	"NutritionCalculator/pkg/services/validation"
	"encoding/json"
	"log"
	"net/http"
)

type PageData struct {
	error string
}

func DayHandler(userDataPath string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		// Get user data based on username from context
		username := r.Context().Value(contextkeys.UserKey).(string)
		userDataService := userData.NewUserDataService(username, userDataPath)

		/*==========================GET=============================*/
		if r.Method == http.MethodGet {
			// Get the meals from the user data
			meals, err := userDataService.GetMeals()
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			DisplayPage(w, meals, "web/template/day.html")
		}

		/*==========================POST=============================*/
		if r.Method == http.MethodPost {
			// Get the day from the form

			var day models.Day
			err := json.NewDecoder(r.Body).Decode(&day)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			//validate the day
			valiationService := &validation.DefaultValidationService{}
			err = valiationService.ValidateDay(day)
			if err != nil {
				log.Println(err)
				http.Error(w, "Invalid User Input: "+err.Error(), http.StatusBadRequest)
				return
			}

			// Add the day to the user data
			err = userDataService.AddDay(day)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			w.Write([]byte("Day added!\n"))

			// Display the day page
			days, err := userDataService.GetDays()
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			DisplayPage(w, days, "web/template/day.html")
		}

		/*==========================PUT=============================*/
		if r.Method == http.MethodPut {
			// Get the day from the form
			var day models.Day
			err := json.NewDecoder(r.Body).Decode(&day)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			valiationService := &validation.DefaultValidationService{}
			err = valiationService.ValidateDay(day)
			if err != nil {
				log.Println(err)
				http.Error(w, "Invalid User Input: "+err.Error(), http.StatusBadRequest)
				return
			}

			//calculate the day's nutritional values
			nutritionalValues, err := userDataService.CalculateDayNutritionalValues(day)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			day.NutritionalValues = &nutritionalValues

			// Update the day to the user data
			err = userDataService.UpdateDay(day)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			w.Write([]byte("Day updated!\n"))

			// Display the day page
			days, err := userDataService.GetDays()
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			DisplayPage(w, days, "web/template/day.html")

		}
		/*==========================DELETE=============================*/
		if r.Method == http.MethodDelete {
			// Get the day from the form
			var day models.Day
			err := json.NewDecoder(r.Body).Decode(&day)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			valiationService := &validation.DefaultValidationService{}
			err = valiationService.ValidateDayForDeletion(day)
			if err != nil {
				log.Println(err)
				http.Error(w, "Invalid User Input: "+err.Error(), http.StatusBadRequest)
				return
			}

			// Delete the day from the user data
			err = userDataService.DeleteDay(day)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			w.Write([]byte("Day deleted!\n"))

			// Display the day page
			days, err := userDataService.GetDays()
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			DisplayPage(w, days, "web/template/day.html")
		}

	}
}
