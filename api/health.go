package api

import "net/http"

// Health - Health api ping.
func Health(w http.ResponseWriter, r *http.Request) {
	respJSON(w, http.StatusOK, message{
		Message: "OK",
	})
}
