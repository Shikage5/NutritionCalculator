package utils

import (
	"log"
	"net/http"
	"text/template"
)

func DisplayPage(w http.ResponseWriter, data interface{}, filePath string) {
	tmpl, err := template.ParseFiles(filePath)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = tmpl.Execute(w, data)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
