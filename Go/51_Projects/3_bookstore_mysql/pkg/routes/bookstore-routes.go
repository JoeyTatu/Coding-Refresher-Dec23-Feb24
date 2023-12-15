package routes

import (
	"github.com/gorilla/mux"
	"github.com/joeytatu/go-bookstore/pkg/controllers"
)

var RegisterBookStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/store/book", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/store", controllers.GetAllBooks).Methods("GET")
	router.HandleFunc("/store/book/{id}", controllers.GetBookById).Methods("GET")
	router.HandleFunc("/store/book/{id}", controllers.UpdateBook).Methods("PUT", "PATCH")
	router.HandleFunc("/store/book/{id}", controllers.DeleteBook).Methods("DELETE")
}
