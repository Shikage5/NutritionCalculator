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
		userData, err := userDataService.GetUserData(username)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if r.Method == http.MethodGet {

			foods := userData.Foods
			//Display the food page
			tmpl, err := template.ParseFiles("web/template/food.html")
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			err = tmpl.Execute(w, foods)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			w.WriteHeader(http.StatusOK)
			return

		} else if r.Method == http.MethodPost {

			//Get the food from the request body
			var food models.Food
			err = json.NewDecoder(r.Body).Decode(&food)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			//Add the food to the user's data
			userDataService.AddFood(username, food)

			//Display the food page

			tmpl, err := template.ParseFiles("web/template/food.html")
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			err = tmpl.Execute(w, userData.Foods)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			//Display a message saying the food was added
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("Food added!"))
			return

		} else if r.Method == http.MethodPut {

			//Get the food from the request body
			var food models.Food
			err = json.NewDecoder(r.Body).Decode(&food)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			//Update the food in the user's data
			userDataService.UpdateFood(username, food)

			//Display the food page
			tmpl, err := template.ParseFiles("web/template/food.html")
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			err = tmpl.Execute(w, userData.Foods)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			//Display a message saying the food was updated
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("Food updated!"))
			return

		} else if r.Method == http.MethodDelete {

			//Get the food from the request body
			var food models.Food
			err = json.NewDecoder(r.Body).Decode(&food)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			//Delete the food from the user's data
			userDataService.DeleteFood(username, food)

			//Display the food page
			tmpl, err := template.ParseFiles("web/template/food.html")
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			err = tmpl.Execute(w, userData.Foods)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			//Display a message saying the food was deleted
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("Food deleted!"))

		}

	}

}
