package middleware

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/joeytatu/go-postgres/models"
	"github.com/joho/godotenv"
)

type response struct {
	Id      int64  `json:"id,omitempty"`
	Message string `json:"message,omitempty"`
}

func createConnection() (db *sql.DB) {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file. Err:", err)
	}

	db, err = sql.Open("postgres", os.Getenv("POSTGRES_URL"))
	if err != nil {
		log.Fatal("Cannot open database. Err:", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("Cannot ping database. Err:", err)
	}

	fmt.Println("Successfully connected to PostgreSQL")
	return db
}

func GetAllStocks(w http.ResponseWriter, r *http.Request) {
	db := createConnection()
	defer db.Close()

	var stocks []models.Stock

	sqlQuery := "SELECT * FROM stocks"

	rows, err := db.Query(sqlQuery)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var stock models.Stock
		err = rows.Scan(&stock.StockId, &stock.Name, &stock.Price, &stock.Company)
		if err != nil {
			log.Fatal("Unable to obtain data from row. Err:", err)
		}
		stocks = append(stocks, stock)
	}

	json.NewEncoder(w).Encode(stocks)
}

func CreateStock(w http.ResponseWriter, r *http.Request) {
	var stock models.Stock

	err := json.NewDecoder(r.Body).Decode(&stock)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	db := createConnection()
	defer db.Close()

	sqlQuery := `
		INSERT INTO stocks(name, price, company)
		VALUES ($1, $2, $3)
		RETURNING stock_id
	`
	var id int64
	err = db.QueryRow(sqlQuery, stock.Name, stock.Price, stock.Company).Scan(&id)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	fmt.Printf("Created stock with stock ID: %v", id)

	res := response{
		Id:      id,
		Message: "stock created successfully",
	}

	json.NewEncoder(w).Encode(res)
}

func GetStockById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	db := createConnection()
	defer db.Close()

	var stock models.Stock

	sqlQuery := `
		SELECT * FROM stocks
		WHERE stock_id = $1
	`

	err = db.QueryRow(sqlQuery, id).Scan(&stock.StockId, &stock.Name, &stock.Price, &stock.Company)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "No stock found", http.StatusNotFound)
			return
		}
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(stock)
}

func UpdateStock(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	var stock models.Stock

	err = json.NewDecoder(r.Body).Decode(&stock)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	db := createConnection()
	defer db.Close()

	sqlQuery := `
		UPDATE stocks
		SET name = $2, price = $3, company = $4
		WHERE stock_id = $1
	`
	query, err := db.Exec(sqlQuery, id, stock.Name, stock.Price, stock.Company)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	rowsAffected, err := query.RowsAffected()
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	fmt.Printf("Total rows affected: %v", rowsAffected)

	msg := fmt.Sprintf("Stock updated successfully. Total rows affected: %v", rowsAffected)

	res := response{
		Id:      int64(id),
		Message: msg,
	}

	json.NewEncoder(w).Encode(res)
}

func DeleteStock(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	db := createConnection()
	defer db.Close()

	sqlQuery := `
		DELETE FROM stocks
		WHERE stock_id = $1
	`

	query, err := db.Exec(sqlQuery, id)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	deletedRows, err := query.RowsAffected()
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	fmt.Printf("Total rows affected: %v", deletedRows)

	msg := fmt.Sprintf("Stock deleted successfully. Total rows affected: %v", deletedRows)

	res := response{
		Id:      int64(id),
		Message: msg,
	}

	json.NewEncoder(w).Encode(res)
}
