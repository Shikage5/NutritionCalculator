package handlers

import (
	contextkeys "NutritionCalculator/pkg/contextKeys"
	"NutritionCalculator/pkg/services/userData"
	"NutritionCalculator/utils"
	"log"
	"net/http"
)

type ExportHandler struct {
	UserDataService userData.DefaultUserDataService
}

func NewExportHandler(userDataService userData.DefaultUserDataService) *ExportHandler {
	return &ExportHandler{UserDataService: userDataService}
}

func (h *ExportHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	username := r.Context().Value(contextkeys.UserKey).(string)

	switch r.Method {
	case http.MethodGet:
		h.handleGet(w, username)
	default:
		http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
	}
}

func (h *ExportHandler) handleGet(w http.ResponseWriter, username string) {
	userData, err := h.UserDataService.GetUserData(username)
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
		handleError(w, err)
		return
	}
}

func handleError(w http.ResponseWriter, err error) {
	log.Println(err)
	http.Error(w, err.Error(), http.StatusInternalServerError)
}
