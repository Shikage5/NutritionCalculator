package handlers

import (
	"NutritionCalculator/data/models"
	contextkeys "NutritionCalculator/pkg/contextKeys"
	"NutritionCalculator/pkg/services/userData"
	"NutritionCalculator/pkg/services/validation"
	"encoding/json"
	"html/template"
	"log"
	"net/http"
)

func DishHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		//Get user data based on username from context
		username := r.Context().Value(contextkeys.UserKey).(string)
		userDataService := userData.NewUserDataService(username)

		/*==========================GET=============================*/
		if r.Method == http.MethodGet {

			//Get the user's data
			dishData, err := userDataService.GetDishData()
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			DisplayPage(w, dishData, "web/template/dishData.html")
			w.WriteHeader(http.StatusOK)
			return

			/*==========================POST=============================*/
		} else if r.Method == http.MethodPost {

			//Get the dish data from the request body
			var dishData models.DishData
			err := json.NewDecoder(r.Body).Decode(&dishData)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			valiationService := &validation.DefaultValidationService{}
			err = valiationService.ValidateDishData(dishData)
			if err != nil {
				log.Println(err)
				http.Error(w, "Invalid User Input: "+err.Error(), http.StatusBadRequest)
				return
			}

			//Calculate the dish's nutrition
			dishData.NutritionalValues, err = userDataService.CalculateDishDataNutritionalValues(dishData)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			//Add the dish to the user's data
			err = userDataService.AddDishData(dishData)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			w.WriteHeader(http.StatusOK)
			w.Write([]byte("Dish added!\n"))

			/*==========================PUT=============================*/
		} else if r.Method == http.MethodPut {

			//Get the dish from the request body
			var dishData models.DishData
			err := json.NewDecoder(r.Body).Decode(&dishData)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			//Update the dish in the user's data
			err = userDataService.UpdateDishData(dishData)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			//Display the dish page
			tmpl, err := template.ParseFiles("web/template/dish.html")
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			err = tmpl.Execute(w, dishData)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			//Display a message saying the dish was updated
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("Dish updated!\n"))
			return

			/*==========================DELETE=============================*/
		} else if r.Method == http.MethodDelete {

			//Get the dish from the request body
			var dishData models.DishData
			err := json.NewDecoder(r.Body).Decode(&dishData)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			//Delete the dish from the user's data
			err = userDataService.DeleteDishData(dishData)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			//Display a message saying the dish was deleted
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("Dish deleted!\n"))
			//Display the dish page
			tmpl, err := template.ParseFiles("web/template/dish.html")
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			err = tmpl.Execute(w, dishData)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}

			return
		}

	}

}
