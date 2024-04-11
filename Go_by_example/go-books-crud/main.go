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

type Author struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Year      int    `json:"year"`
}

var books []Book // slice of type books

func main() {
	r := mux.NewRouter()

	books = append(books)
}
