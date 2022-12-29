package main

import (
	// "fmt"
	//m "github.com/Vfulgham/go-mux-api/math"

	"log"
	"net/http"

	"github.com/Vfulgham/go-mux-api/pkg/db"
	"github.com/Vfulgham/go-mux-api/pkg/handlers"
	"github.com/gorilla/mux"
)

func main() {
	DB := db.Init()       // initialize database
	h := handlers.New(DB) // pass to handler so handlers can access database object
	router := mux.NewRouter()

	// define api endpoints
	// has 2 arguments - path and handler (inline function)
	router.HandleFunc("/books", h.GetAllBooks).Methods(http.MethodGet)
	router.HandleFunc("/books", h.AddBook).Methods(http.MethodPost)
	router.HandleFunc("/books/{id}", h.GetBook).Methods(http.MethodGet)
	router.HandleFunc("/books/{id}", h.UpdateBook).Methods(http.MethodPut) // add to mock example (not added to db)
	router.HandleFunc("/books/{id}", h.DeleteBook).Methods(http.MethodDelete)

	log.Println("API is running")
	http.ListenAndServe(":4000", router)

	practice()
}
