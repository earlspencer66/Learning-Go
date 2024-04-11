package main

import (
	"github.com/gorilla/mux"
)

type Book struct {
	ID     string  `json:"id"`
	Isbn   string  `json:"isbn"`
	Title  string  `json: "title"`
	Author *Author `json:author"`
}

func main() {
	r := mux.NewRouter()

	books = append(books)
}
