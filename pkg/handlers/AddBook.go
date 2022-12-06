package handlers

import (
	"log"
	"net/http"
	"encoding/json"
	"math/rand"
	"io/ioutil"

	"github.com/Vfulgham/go-mux-api/pkg/models"
	"github.com/Vfulgham/go-mux-api/pkg/mocks"
)

func AddBook(w http.ResponseWriter, r *http.Request) {
    
    defer r.Body.Close()
    body, err := ioutil.ReadAll(r.Body)

    if err != nil {
        log.Fatalln(err)
    }

    var book models.Book
    json.Unmarshal(body, &book) // map json to struct and put into variable book

    // Append to the Book mocks
    book.Id = rand.Intn(100)
    mocks.Books = append(mocks.Books, book)

    // Send a 201 created response
    w.Header().Add("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode("Created")
}