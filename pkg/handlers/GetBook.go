package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Vfulgham/go-mux-api/pkg/models"
	"github.com/gorilla/mux"
)

func (h handler) GetBook(w http.ResponseWriter, r *http.Request) {
    // Read id parameter
    vars := mux.Vars(r)
    id, _ := strconv.Atoi(vars["id"]) // convert string to int

    var book models.Book

    // get book from db
    result := h.DB.First(&book, id)

    if result.Error != nil{
        fmt.Println(result.Error)
    }

    w.Header().Add("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)

    json.NewEncoder(w).Encode(book)
}