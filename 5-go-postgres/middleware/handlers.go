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

func createConnection() *sql.DB {
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
	allStocks, err := getAllStock()
	if err != nil {
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
	var stock models.Stock
	err = json.NewDecoder(r.Body).Decode(&stock)
	if err != nil {
		log.Fatal("Unable to decode the request body %v", err)
	}
	updateRow := updateStock(int64(id), stock)
	msg := fmt.Sprintf("Stock successfully updated row affected %v ", updateRow)

	res := response{
		ID:      int64(id),
		Message: msg,
	}
	json.NewEncoder(w).Encode(res)
}

func DeleteStocks(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Fatal("Unable to convert into integer")
	}
	deleteRow := deleteStock(int64(id))
	msg := fmt.Sprintf("stock deleted successfully %s row affected", &deleteRow)
	res := response{
		ID:      int64(id),
		Message: msg,
	}
	json.NewEncoder(w).Encode(res)
}

func insertStock(stock models.Stock) int64 {
	db := createConnection()
	defer db.Close()
	sqlStatement := `INSERT INTO stock(name,price,company) value($1,$2,$3) RETURNING stockid`
	var id int64
	err := db.QueryRow(sqlStatement, stock.Name, stock.Price, stock.Company).Scan(&id)
	if err != nil {
		log.Fatal("Unable to decode the request body %v", err)
	}
	fmt.Printf("Inserted single record %v", id)
	return id
}

func getStock(id int64) (models.Stock, error) {
	db := createConnection()
	defer db.Close()
	var stock models.Stock
	sqlStatement := `SELECT * FROM stocks WHERE stock id = $1`
	row := db.QueryRow(sqlStatement, id)
	err := row.Scan(&stock.StockID, &stock.Name, &stock.Price, &stock.Company)
	switch err {
	case sql.ErrNoRows:
		fmt.Printf("No rows where returned")
	case nil:
		return stock, nil
	default:
		log.Fatal("Unable to scan the row %v", err)

	}
	return stock, err

}

func getAllStock() ([]models.Stock, error) {
	db := createConnection()
	defer db.Close()
	var stocks []models.Stock //slice as there is multiple data
	sqlStatement := `SELECT * FROM stocks`
	rows, err := db.Query(sqlStatement)
	if err != nil {
		log.Fatal("Unable to execute the query %v", err)
	}
	defer rows.Close()
	for rows.Next() {
		var stock models.Stock
		err := rows.Scan(&stock.StockID, &stock.Name, &stock.Price, &stock.Company)
		if err != nil {
			log.Fatal("Unable to execute the query %v", err)
		}
		stocks = append(stocks, stock)
	}
	return stocks, err
}

func updateStock(id int64, stock models.Stock) int64 {
	db := createConnection()
	defer db.Close()
	sqlStatement := `UPDATE * FROM stocks SET name=$2 , price=$3 ,company=$4 WHERE stockid=$1 `
	res, err := db.Exec(sqlStatement, id, stock.Name, stock.Price, stock.Company)
	if err != nil {
		log.Fatal("Unable to execute the query %v", err)
	}
	rowAffected, err := res.RowsAffected()
	if err != nil {
		log.Fatal("error while checking the affected row %v", err)
	}
	fmt.Println("Totoal affected row", rowAffected)
	return rowAffected
}

func deleteStock(id int64) int64 {
	db := createConnection()
	defer db.Close()
	sqlStatement := `DELETE FROM stocks where stockid=$1`
	res, err := db.Exec(sqlStatement, id)
	if err != nil {
		log.Fatal("Unable to execute the query %v", err)
	}
	rowAffected, err := res.RowsAffected()
	if err != nil {
		log.Fatal("error while checking the affected row %v", err)
	}
	fmt.Printf("Total rows affected")
	return rowAffected
}
