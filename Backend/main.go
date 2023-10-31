package main

import (
	"NutritionCalculator/basicAuth"
	"log"
	"net/http"
)

func mainHandler(w http.ResponseWriter, r *http.Request) {

	responseString := "<html><body>Hello WORLD</body></html>"
	w.Write([]byte(responseString)) // unbedingt Templates verwenden!
}

func main() {
	http.HandleFunc("/", basicAuth.Wrapper(mainHandler))

	log.Fatalln(http.ListenAndServe(":8080", nil))
}
