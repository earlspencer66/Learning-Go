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

	books = append(books, Book{ID: "1", Isbn: "1563254", Title: "The Desire of Ages", Author: &Author{Firstname: "Ellen", Lastname: "White", Year: 1905}})
}
