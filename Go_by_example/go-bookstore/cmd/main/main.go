package main

import (
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/jihnzu/gorm/dialects/mysql"
	"github.com/pkg/routes"

	//need to import the routes from the pkg folder
	"github.com/earlspencer66/Learning-Go/tree/main/Go_by_example/go-bookstore/pkg/routes"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	log.fatal(http.ListenAndServe("localhost:9010", r))
}

//create functions for connectiong to database
