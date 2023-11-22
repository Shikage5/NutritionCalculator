package main

import (
	registrationHandlers "NutritionCalculator/pkg/handlers/registration"
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"time"
)

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World! %s", time.Now())
}

func main() {
	http.HandleFunc("/", greet)
	http.HandleFunc("/register", registrationHandlers.RegisterHandler)

	// Get the absolute path to the certificate file
	certFile, err := filepath.Abs("server.crt")
	if err != nil {
		log.Fatal("Failed to get absolute path for server.crt: ", err)
	}

	// Get the absolute path to the key file
	keyFile, err := filepath.Abs("server.key")
	if err != nil {
		log.Fatal("Failed to get absolute path for server.key: ", err)
	}

	err = http.ListenAndServeTLS(":443", certFile, keyFile, nil)
	if err != nil {
		log.Fatal("ListenAndServeTLS: ", err)
	}
}
