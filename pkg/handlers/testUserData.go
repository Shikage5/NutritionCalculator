package handlers

import (
	contextkeys "NutritionCalculator/pkg/contextKeys"
	"NutritionCalculator/pkg/services/userData"
	"encoding/json"
	"net/http"
)

// ignore this
func TestUserData(UserDataService userData.UserDataService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			username := r.Context().Value(contextkeys.UserKey).(string)
			userData, err := UserDataService.GetUserData(username)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			// Marshal the userData object into JSON
			jsonData, err := json.Marshal(userData)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			// Set the Content-Type header to application/json
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)

			// Write the JSON data to the response
			w.Write(jsonData)
			return
		}
	}

}
