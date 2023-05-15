package routes

import (
	"bookstore/pkg/controller"

	"github.com/gorilla/mux"
)

var BookRoute = func(r mux.Router) {

	r.HandleFunc("/book/", controller.CreateBook).Methods("POST")
	r.HandleFunc("/book/", controller.GetBook).Methods("GET")
	r.HandleFunc("/book/{bookId}", controller.GetBookById).Methods("GET")
	r.HandleFunc("/book/{bookId}", controller.UpdateBook).Methods("PUT")
	r.HandleFunc("/book/{bookId}", controller.DeleteBook).Methods("DELETE")

}
