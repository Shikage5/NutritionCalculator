package handlers

import "net/http"

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	// Handle the homepage, maybe show some general information
	http.ServeFile(w, r, "pkg/templates/index.html")
}
