package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

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

func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
	return
}

func getBook(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range books {
		if item.ID == params["id"]{
			json.NewEncoder(w).Encode(item)
			return
		}
}

func createBook(w http.http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.ID = strconv.Ito(rand.Intn(100000000))
	books = append(books, book)
	json.NewEncoder(w).Encode(book)
	return
}

func deleteBook (w http.http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range books {
		if item.ID == params["id"]{
			books = append(books[:index], books[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(books)//return the remaining books
}

func updateBook (w http.http.ResponseWriter, r *http.Request) {
	//basically delete and add book...to the same id
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range books{
		if item.ID == params ["id"]{
			books = append(books[:index], books[index+1:]...)
			var book Book
			_ = json.NewDecoder(r.Body).Decode(&book)
			book.ID = params["id"]
			books = append(books, book)
			json.NewEncoder(w).Encode(book)
			return
		}
	}
}

func main() {
	r := mux.NewRouter()

	books = append(books, Book{ID: "1", Isbn: "1563254", Title: "The Desire of Ages", Author: &Author{Firstname: "Ellen", Lastname: "White", Year: 1905}})

	//need five different routes - see the flowchart
	//"GET ALL", "GET BY ID", "CREATE", "UPDATE", "DELETE"
	//with their corresponding methods

	r.HandleFunc("/books", getBooks).Methods("GET")
	r.HandleFunc("/books/id", getBook).Methods("GET")
	r.HandleFunc("/books", createBook).Methods("POST")
	r.HandleFunc("/books/id", deleteBook).Methods("DELETE")
	r.HandleFunc("/books/id", updateBook).Methods("PUT")

	fmt.Printf("Starting server at port 8000\n")
	log.Fatal(http.ListenAndServe(":8000", r))
}
