package handlers

import (
	contextkeys "NutritionCalculator/pkg/contextKeys"
	"NutritionCalculator/pkg/services/userData"
	"NutritionCalculator/utils"
	"net/http"
)

func ExportHandler(userDataPath string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// Get user data based on username from context
		username := r.Context().Value(contextkeys.UserKey).(string)
		userDataService := userData.NewUserDataService(username, userDataPath)

		/*==========================GET=============================*/
		if r.Method == http.MethodGet {
			// Get the data
			userData, err := userDataService.GetUserData()
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			// Set the content type header to the MIME type of the file.
			// This tells the client what the file is, so it knows how to handle it.
			w.Header().Set("Content-Type", "text/csv")

			// Set the Content-Disposition header to "attachment", along with the filename.
			// This tells the client that it should download the file instead of displaying it.
			w.Header().Set("Content-Disposition", "attachment; filename=days.csv")

			// Write the data to CSV
			err = utils.WriteDaysToCSV(w, userData.Days)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}
	}
}
