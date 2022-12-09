package handlers

import (
    "encoding/json"
    "net/http"
    "strconv"

    "github.com/gorilla/mux"
	"github.com/Vfulgham/go-mux-api/pkg/mocks"
)

func (h handler) GetBook(w http.ResponseWriter, r *http.Request) {
    // Read id parameter
    vars := mux.Vars(r)
    id, _ := strconv.Atoi(vars["id"]) // convert string to int

    // Iterate over all the mock books
    for _, book := range mocks.Books {
        if book.Id == id {
            // If ids are equal send book as a response
            w.Header().Add("Content-Type", "application/json")
            w.WriteHeader(http.StatusOK)

            json.NewEncoder(w).Encode(book)
            break
        }
    }
}