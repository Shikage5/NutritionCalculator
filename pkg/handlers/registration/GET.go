package registrationHandlers

import (
	"net/http"
)

func RegisterGETHandler(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("This is a registration form"))
}
