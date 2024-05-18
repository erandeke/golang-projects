package routes

import (
	"go-bookstore/pkg/controllers"

	"github.com/gorilla/mux"
)

// register the routes
func RegisterRoute(router *mux.Router) {
	router.HandleFunc("/books/", controllers.CreateBooks).Methods("POST")
	router.HandleFunc("/books/", controllers.GetBooks).Methods("GET")
	router.HandleFunc("/book/{id}", controllers.GetBookById).Methods("GET")
	//router.HandleFunc("/book/{id}", controllers.updateBook).Methods("PUT")
	//router.HandleFunc("/book/{id}", controllers.deleteBook).Methods("DELETE")
}
