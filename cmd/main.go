package main

import (
	pagelogin "NutritionCalculator/pkg/handlers/pageLogin"
	"fmt"
	"net/http"
	"time"
)

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World! %s", time.Now())
}

func main() {

	http.HandleFunc("/", greet)
	http.HandleFunc("/register", pagelogin.RegisterHandle)
	http.ListenAndServeTLS(":443", "server.crt", "server.key", nil)
}
