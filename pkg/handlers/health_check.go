package handlers

import (
	"net/http"
)

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	// fmt.Println(r.URL.Query().Get("param2"))
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("hello World"))
}
