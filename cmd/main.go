package main

import (
	// "fmt"
	//m "github.com/Vfulgham/go-mux-api/math"

	"log"
	"net/http"
	"github.com/Vfulgham/go-mux-api/pkg/handlers"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	// define api endpoints
	// has 2 arguments - path and handler (inline function)
	router.HandleFunc("/books", handlers.GetAllBooks).Methods(http.MethodGet)
	router.HandleFunc("/books", handlers.AddBook).Methods(http.MethodPost)
	router.HandleFunc("/books/{id}", handlers.GetBook).Methods(http.MethodGet)
	router.HandleFunc("/books/{id}", handlers.UpdateBook).Methods(http.MethodPut)
	router.HandleFunc("/books/{id}", handlers.DeleteBook).Methods(http.MethodDelete)

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

