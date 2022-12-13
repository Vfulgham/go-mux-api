package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
    "fmt"

	"github.com/Vfulgham/go-mux-api/pkg/models"
	"github.com/gorilla/mux"
)

func (h handler) DeleteBook(w http.ResponseWriter, r *http.Request) {
	// Read id parameter
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

    var book models.Book

    // if id matches id on any book in db
	if result := h.DB.First(&book, id); result.Error != nil {
		fmt.Println(result.Error)
	}
    // delete book
    h.DB.Delete(&book)

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Deleted")

}
