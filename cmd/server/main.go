package server

import (
	pagelogin "NutritionCalculator/pkg/handlers/registration"
	"fmt"
	"net/http"
	"time"
)

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World! %s", time.Now())
}

func main() {

	http.HandleFunc("/", greet)
	http.HandleFunc("/register", pagelogin.RegisterHandler)
	http.ListenAndServeTLS(":443", "server.crt", "server.key", nil)
}
