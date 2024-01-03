package handlers

import (
	"NutritionCalculator/data/models"
	contextkeys "NutritionCalculator/pkg/contextKeys"
	"NutritionCalculator/pkg/services/userData"
	"encoding/json"
	"html/template"
	"net/http"
)

func FoodHandler(userDataService userData.UserDataService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		//Get user data based on username from context
		username := r.Context().Value(contextkeys.UserKey).(string)

		if r.Method == http.MethodGet {

			//Get the user's data
			foodData, err := userDataService.GetFoodData(username)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			//Display the food page
			tmpl, err := template.ParseFiles("web/template/food.html")
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			err = tmpl.Execute(w, foodData)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			w.WriteHeader(http.StatusOK)
			return

		} else if r.Method == http.MethodPost {

			//Get the food from the request body
			var foodData models.FoodData
			err := json.NewDecoder(r.Body).Decode(&foodData)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			//Add the food to the user's data
			err = userDataService.AddFoodData(username, foodData)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			w.WriteHeader(http.StatusOK)
			w.Write([]byte("Food added!\n"))

		} else if r.Method == http.MethodPut {

			//Get the food from the request body
			var foodData models.FoodData
			err := json.NewDecoder(r.Body).Decode(&foodData)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			//Update the food in the user's data
			err = userDataService.UpdateFoodData(username, foodData)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			//Display the food page
			tmpl, err := template.ParseFiles("web/template/food.html")
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			err = tmpl.Execute(w, foodData)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			//Display a message saying the food was updated
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("Food updated!\n"))
			return

		} else if r.Method == http.MethodDelete {

			//Get the food from the request body
			var foodData models.FoodData
			err := json.NewDecoder(r.Body).Decode(&foodData)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			//Delete the food from the user's data
			err = userDataService.DeleteFoodData(username, foodData)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			//Display a message saying the food was deleted
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("Food deleted!\n"))
			//Display the food page
			tmpl, err := template.ParseFiles("web/template/food.html")
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			err = tmpl.Execute(w, foodData)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}

			return
		}

	}

}
