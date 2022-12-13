package handlers

import (
	"encoding/json"
	"net/http"
	"fmt"

	"github.com/Vfulgham/go-mux-api/pkg/models"
)

func (h handler) GetAllBooks(w http.ResponseWriter, r *http.Request){
	var books []models.Book

	// finds model books in db and fills slice
	result := h.DB.Find(&books);

	if result.Error != nil {
		fmt.Println(result.Error)
	}

	w.Header().Add("Content-Type","application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(books)
	
}