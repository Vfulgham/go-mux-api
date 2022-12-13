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
	DB := db.Init() // initialize database
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
 
	// xs := []float64{}
	// avg, _ := m.Average(&xs)
	// fmt.Println(avg)

	// var nil_slice []int
    // var empty_slice = []int{}
    // var emptySlice1 = make([]int, 0)

    // fmt.Println(nil_slice == nil, len(nil_slice), cap(nil_slice))
    // fmt.Println(empty_slice == nil, len(empty_slice), cap(empty_slice))
    // fmt.Println(emptySlice1 == nil, len(empty_slice), cap(empty_slice))


}

