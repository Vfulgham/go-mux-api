package handlers

import (
	"log"
	"net/http"
	"encoding/json"
	"io/ioutil"
    "fmt"

	"github.com/Vfulgham/go-mux-api/pkg/models"
)

func (h handler) AddBook(w http.ResponseWriter, r *http.Request) {
    
    defer r.Body.Close()
    body, err := ioutil.ReadAll(r.Body)

    if err != nil {
        log.Fatalln(err)
    }

    var book models.Book
    json.Unmarshal(body, &book) // map json to struct and put into variable book

    // add book to db
    if result := h.DB.Create(&book); result.Error != nil {
            fmt.Println(result.Error)
    }

    // Send a 201 created response
    w.Header().Add("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode("Created")
}