package registrationHandlers

import (
	"net/http"
)

// RegisterHandler is the intermediate handler that dispatches the request to the appropriate specialized handler.
func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		RegisterGETHandler(w, r)
	case http.MethodPost:
		RegisterPOSTHandler(w, r)
	default:
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}
