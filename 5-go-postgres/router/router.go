package router

import (
	"postgresql/middleware"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/api/stock/{id}", middleware.GetStock).Methods("GET")
	router.HandleFunc("/api/stock", middleware.GetAllStock).Methods("GET")
	router.HandleFunc("/api/newstock", middleware.CreateStock).Methods("POST")
	router.HandleFunc("/api/stock/{id}", middleware.UpdateStock).Methods("PUT")
	router.HandleFunc("/api/deletestock/{id}", middleware.DeleteStock).Methods("DELETE")

}
