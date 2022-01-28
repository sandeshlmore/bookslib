package routes

import (
	"github.com/gorilla/mux"
	"github.com/sandeshlmore/bookslib/pkg/handlers"
)

func RegisterRoutes(router *mux.Router) {

	router.HandleFunc("/health-check", handlers.HealthCheck).Methods("GET")
	router.HandleFunc("/books", handlers.GetBooks).Methods("GET")
	router.HandleFunc("/books", handlers.AddBook).Methods("POST")

}
