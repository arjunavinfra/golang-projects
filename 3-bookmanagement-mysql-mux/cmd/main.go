package main

import (
	"bookstore/pkg/routes"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	r := mux.NewRouter()
	routes.BookRoute(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8000", r))

}
