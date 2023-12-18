package router

import (
	"github.com/gorilla/mux"
	"github.com/joeytatu/go-postgres/middleware"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	idUrl := "/api/stocks/stock/{id}"
	idUrlS := idUrl + "/"

	router.HandleFunc("/api/stocks", middleware.GetAllStocks).Methods("GET")
	router.HandleFunc("/api/stocks/stock", middleware.CreateStock).Methods("POST")
	router.HandleFunc(idUrl, middleware.GetStockById).Methods("GET")
	router.HandleFunc(idUrl, middleware.UpdateStock).Methods("PUT", "PATCH")
	router.HandleFunc(idUrl, middleware.DeleteStock).Methods("DELETE")

	// Route path to allow "/" at the end of URL
	router.HandleFunc("/api/stocks/", middleware.GetAllStocks).Methods("GET")
	router.HandleFunc("/api/stocks/stock/", middleware.CreateStock).Methods("POST")
	router.HandleFunc(idUrlS, middleware.GetStockById).Methods("GET")
	router.HandleFunc(idUrlS, middleware.UpdateStock).Methods("PUT", "PATCH")
	router.HandleFunc(idUrlS, middleware.DeleteStock).Methods("DELETE")

	return router
}
