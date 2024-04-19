package models

import (
	"C:/Users/earls/Documents/Learning-Go/Go_by_example/go-bookstore/pkg/config"

	"github.com/jihnzu/gorm"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `gorm:""json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

// to connect to the database
func init() {
	config.Connect()
	db = config.GetDB()
	db.Automigrate(&Book{})
}

// Logic
// User interacts with the routes. Routes with the controllers
// Controllers perform operations with the database

// different functions for the controllers
func (b *Book) CreateBook() *Book {
	db.NewRecord(b)
	db.Create(&b)
	return b
}

func GetAllBooks() []Book {
	var Books []Book
	db.Find(&Books)
	return Books
}

func GetBookById(Id int64) (*Book, *gorm.DB) {
	var getBook Book
	db := db.Where("ID=?", Id).Find(&getBook)
	return &getBook, db
}

func DeleteBook(ID, int64) Book { //notice difference between ID and Id here 
	var book Book
	db.Where("ID=?", ID).Delete(book)
	return book
}

//Update will not be included
//Combination of three functions to implement update