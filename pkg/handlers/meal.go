package handlers

import (
	"NutritionCalculator/data/models"
	contextkeys "NutritionCalculator/pkg/contextKeys"
	"NutritionCalculator/pkg/services/userData"
	"NutritionCalculator/pkg/services/validation"
	"encoding/json"
	"log"
	"net/http"
	"strings"
)

func MealHandler(userDataPath string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		//Get user data based on username from context
		username := r.Context().Value(contextkeys.UserKey).(string)
		userDataService := userData.NewUserDataService(username, userDataPath)

		/*==========================GET=============================*/
		if r.Method == http.MethodGet {

			//Get the user's data
			meals, err := userDataService.GetMeals()
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			DisplayPage(w, meals, "web/template/meals.html")
			return

			/*==========================POST=============================*/
		} else if r.Method == http.MethodPost {

			//Get the meal data from the request body
			var meal models.Meal
			err := json.NewDecoder(r.Body).Decode(&meal)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			valiationService := &validation.DefaultValidationService{}
			err = valiationService.ValidateMeal(meal)
			if err != nil {
				log.Println(err)
				http.Error(w, "Invalid User Input: "+err.Error(), http.StatusBadRequest)
				return
			}

			//Calculate the meal's nutrition
			nutritionalValues, err := userDataService.CalculateMealNutritionalValues(meal)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			meal.NutritionalValues = &nutritionalValues
			//Add the meal to the user's data
			err = userDataService.AddMeal(meal)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			w.Write([]byte("Meal added!\n"))

			DisplayPage(w, meal, "web/template/meals.html")
			return

			/*==========================PUT=============================*/
		} else if r.Method == http.MethodPut {

			//Get the meal from the request body
			var meal models.Meal
			err := json.NewDecoder(r.Body).Decode(&meal)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			valiationService := &validation.DefaultValidationService{}
			err = valiationService.ValidateMeal(meal)
			if err != nil {
				log.Println(err)
				http.Error(w, "Invalid User Input: "+err.Error(), http.StatusBadRequest)
				return
			}

			//calculate the meal's nutrition
			nutritionalValues, err := userDataService.CalculateMealNutritionalValues(meal)
			if err != nil {
				if strings.Contains(err.Error(), "circular reference") {
					http.Error(w, "Invalid input: "+err.Error(), http.StatusBadRequest)
				} else {
					http.Error(w, err.Error(), http.StatusInternalServerError)
				}
				return
			}
			meal.NutritionalValues = &nutritionalValues

			//Update the meal in the user's data
			err = userDataService.UpdateMeal(meal)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			//Display a message saying the meal was updated
			w.Write([]byte("Meal updated!\n"))

			DisplayPage(w, meal, "web/template/meals.html")

			return

			/*==========================DELETE=============================*/
		} else if r.Method == http.MethodDelete {

			//Get the meal from the request body
			var meal models.Meal
			err := json.NewDecoder(r.Body).Decode(&meal)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			validationService := &validation.DefaultValidationService{}
			err = validationService.ValidateMeal(meal)
			if err != nil {
				log.Println(err)
				http.Error(w, "Invalid User Input: "+err.Error(), http.StatusBadRequest)
				return
			}

			//Delete the meal from the user's data
			err = userDataService.DeleteMeal(meal)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			w.Write([]byte("Meal deleted!\n"))

			DisplayPage(w, meal, "web/template/meals.html")

			return
		}

	}

}
