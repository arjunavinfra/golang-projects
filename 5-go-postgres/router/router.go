package router

import (
	"postgresql/middleware"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/api/stock/{id}", middleware.GetStocks).Methods("GET")
	router.HandleFunc("/api/stock", middleware.GetAllStocks).Methods("GET")
	router.HandleFunc("/api/newstock", middleware.CreateStock).Methods("POST")
	router.HandleFunc("/api/stock/{id}", middleware.UpdateStocks).Methods("PUT")
	router.HandleFunc("/api/deletestock/{id}", middleware.DeleteStocks).Methods("DELETE")
	return router

}
