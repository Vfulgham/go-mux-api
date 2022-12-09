package handlers

import (
    "encoding/json"
    "io/ioutil"
    "log"
    "net/http"
    "strconv"

    "github.com/gorilla/mux"
    "github.com/Vfulgham/go-mux-api/pkg/mocks"
    "github.com/Vfulgham/go-mux-api/pkg/models"
)

func (h handler) UpdateBook(w http.ResponseWriter, r *http.Request) {
    // Read id parameter
    vars := mux.Vars(r)
    id, _ := strconv.Atoi(vars["id"])

    // Read body from request
    defer r.Body.Close()
    body, err := ioutil.ReadAll(r.Body)

    if err != nil {
        log.Fatalln(err)
    }

    var updatedBook models.Book
    json.Unmarshal(body, &updatedBook) // map json to struct

    // Iterate over all the mock Books and update
    for index, book := range mocks.Books {
        if book.Id == id {
            // Update and send a response when book Id matches dynamic Id
            book.Title = updatedBook.Title
            book.Author = updatedBook.Author
            book.Desc = updatedBook.Desc

            mocks.Books[index] = book
            w.Header().Add("Content-Type", "application/json")
            w.WriteHeader(http.StatusOK)

            json.NewEncoder(w).Encode("Updated")
            break
        }
    }
}