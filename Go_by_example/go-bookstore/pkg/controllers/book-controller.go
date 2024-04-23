package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/pkg/models"

	//import utils

	//import models
	"C:/Users/earls/Documents/Learning-Go/Go_by_example/go-bookstore/pkg/models"
)

var NewBook models.Book

//getBooks function first

func GetBook(w http.ResponseWriter, r http.Request) {
	newBooks := models.GetAllBooks()
	res, _ := json.Marshal(newBooks)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
