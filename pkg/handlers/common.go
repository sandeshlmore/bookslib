package handlers

import (
	"encoding/json"
	"net/http"
)

func makeError(w http.ResponseWriter, errorcode int, message string) {

	response, err := json.Marshal(map[string]string{"error": message})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(errorcode)
	w.Write([]byte(response))
}
