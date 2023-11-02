package main

import (
	"NutritionCalculator/pkg/handlers"
	"log"
	"net/http"
)

func mainHandler(w http.ResponseWriter, r *http.Request) {

	responseString := "<html><body>Hello WORLD</body></html>"
	w.Write([]byte(responseString)) // unbedingt Templates verwenden!
}

func main() {

	http.HandleFunc("/", handlers.IndexHandler)
	http.HandleFunc("/login", handlers.LoginHandler)
	http.HandleFunc("/register", handlers.RegisterHandler)

	log.Println("Server is running on :8080...")
	log.Fatalln(http.ListenAndServe(":8080", nil))
}
