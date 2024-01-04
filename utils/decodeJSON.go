package utils

import "net/http"

func DecodeFromHTML(r http.Request, v interface{}) error {
	err := r.ParseForm()
	if err != nil {
		return err
	}
	return nil
}
