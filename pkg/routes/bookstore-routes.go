package routes

import (
	"github.com/Aditya7feb/bookstore-golang/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterBookStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/api/books", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/api/books", controllers.GetBooks).Methods("GET")
	router.HandleFunc("/api/books/{bookId}", controllers.GetBookById).Methods("GET")
	router.HandleFunc("/api/books/{bookId}", controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/api/books/{bookId}", controllers.DeleteBook).Methods("DELETE")
}
