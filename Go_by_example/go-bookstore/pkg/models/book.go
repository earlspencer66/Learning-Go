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
