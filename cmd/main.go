package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sandeshlmore/bookslib/pkg/routes"
)

func main() {

	router := mux.NewRouter()

	routes.RegisterRoutes(router)

	if err := http.ListenAndServe("0.0.0.0:8080", router); err != nil && err != http.ErrServerClosed {
		log.Fatal("Error: ", err)
	}

}
