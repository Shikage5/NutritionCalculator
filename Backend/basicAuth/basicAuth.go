package basicAuth

import (
	"fmt"
	"net/http"
)

// Authentication Check here
func checkUserValid(user, pswd string) bool {
	fmt.Println(user, ":", pswd)
	return true
}
func Wrapper(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user, pswd, ok := r.BasicAuth()
		if ok && Authenticator.Authenticate(user, pswd) {
			handler(w, r)
		} else {
			w.Header().Set("WWW-Authenticate",
				"Basic realm=\"My Simple Server\"")
			http.Error(w,
				http.StatusText(http.StatusUnauthorized),
				http.StatusUnauthorized)
		}
	}
}
