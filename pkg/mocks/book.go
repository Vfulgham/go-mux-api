package mocks

import (
	"github.com/Vfulgham/go-mux-api/pkg/models"
)

var Books = []models.Book{
	{
		Id:     1,
		Title:  "GoLang",
		Author: "Gopher",
		Desc:   "A book for Go",
	},
}