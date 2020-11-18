package api

import (
	"encoding/json"
	"net/http"
)

type message struct {
	Message string `json:"message"`
}

func respJSON(w http.ResponseWriter, code int, entity interface{}) {
	response, _ := json.Marshal(entity)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
