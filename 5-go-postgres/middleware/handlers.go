package middleware

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"postgresql/models"
	"strconv"

	"github.com/gorilla/mux"
)

type response struct {
	ID      int64  `json:"id,omitempty"`
	Message string `json:"message,omitempty"`
}

func CreateConnection() *sql.DB {
	const sqlConnectionString = "postgresql://28631a8e-b665-4ebd-bedb-9189f5f9a10e-user:pw-1df59d3d-e515-4dee-81b0-25c5962c4e2f@postgres-free-tier-v2020.gigalixir.com:5432/28631a8e-b665-4ebd-bedb-9189f5f9a10e"
	db, err := sql.Open("postgres", sqlConnectionString)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Successsfully connected to postgresql")

	return db
}

func CreateStock(w http.ResponseWriter, r *http.Request) {
	var stock models.Stock
	err := json.NewDecoder(r.Body).Decode(&stock)
	if err != nil {
		log.Fatal("Unable to decode the request body %v", err)
	}
	insertID := insertStock(stock)
	res := response{
		ID:      insertID,
		Message: "stock created successfully",
	}
	json.NewEncoder(w).Encode(res)
}

func GetStocks(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Fatal("Unable to convert into integer")
	}
	stock, err := getStock(int64(id))
	if err != nil {
		log.Fatal("Unable to get stock")
	}
	json.NewEncoder(w).Encode(stock)
}

func GetAllStocks(w http.ResponseWriter, r *http.Request) {
	allStocks, err := selectAllStocks()
	if err := nil {
		log.Fatal("Error in getting stocks")
	}
	json.NewEncoder(w).Encode(allStocks)
}

func UpdateStocks(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Fatal("Unable to convert into integer")
	}
	var stock model.Stock
	err = json.NewDecoder(r.Body).Decode(&stock)
	if err != nil {
		log.Fatal("Unable to decode the request body %v", err)
	}
	updateRow := UpdateStock(int64(id),stock)
	msg := fmt.Sprintf("Stock successfully updated row affected %v ",updateRow)

	res := response{
		ID:      int64(id),
		Message: "stock updated successfully",
	}
	json.NewEncoder(w).Encode(res)
}

func DeleteStocks(w http.ResponseWriter, r *http.Request) {

}
