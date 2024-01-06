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

func FoodHandler(userDataPath string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		//Get user data based on username from context

		username := r.Context().Value(contextkeys.UserKey).(string)
		userDataService := userData.NewUserDataService(username, userDataPath)
		/*==========================GET=============================*/
		if r.Method == http.MethodGet {

			//Get the user's data
			foodData, err := userDataService.GetFoodData()
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			DisplayPage(w, foodData, "web/template/foodData.html")

			return

			/*==========================POST=============================*/
		} else if r.Method == http.MethodPost {

			//Get the food data from the request body
			var foodData models.FoodData
			err := json.NewDecoder(r.Body).Decode(&foodData)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			valiationService := &validation.DefaultValidationService{}
			err = valiationService.ValidateFoodData(foodData)
			if err != nil {
				log.Println(err)
				http.Error(w, "Invalid User Input: "+err.Error(), http.StatusBadRequest)
				return
			}

			//Add the food to the user's data
			err = userDataService.AddFoodData(foodData)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			w.Write([]byte("Food added!\n"))
			DisplayPage(w, foodData, "web/template/food.html")

			/*==========================PUT=============================*/
		} else if r.Method == http.MethodPut {

			//Get the food from the request body
			var foodData models.FoodData
			err := json.NewDecoder(r.Body).Decode(&foodData)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			valiationService := &validation.DefaultValidationService{}
			err = valiationService.ValidateFoodData(foodData)
			if err != nil {
				log.Println(err)
				http.Error(w, "Invalid User Input: "+err.Error(), http.StatusBadRequest)
				return
			}

			//Update the food in the user's data
			err = userDataService.UpdateFoodData(foodData)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			w.Write([]byte("Food updated!\n"))
			DisplayPage(w, foodData, "web/template/food.html")

			return

			/*==========================DELETE=============================*/
		} else if r.Method == http.MethodDelete {

			//Get the food from the request body
			var foodData models.FoodData
			err := json.NewDecoder(r.Body).Decode(&foodData)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			//Delete the food from the user's data
			err = userDataService.DeleteFoodData(foodData)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			//Display a message saying the food was deleted
			w.Write([]byte("Food deleted!\n"))
			DisplayPage(w, foodData, "web/template/food.html")
		}

	}

}
