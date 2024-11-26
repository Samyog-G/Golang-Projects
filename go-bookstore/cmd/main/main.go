package main

import (
	"log"
	"net/http"

	"github.com/Samyog-G/go-bookstore/pkg/routes"
	"github.com/gorilla/mux"
)

//user interacts with the routes and routes sent controls to the controllers where we have all our logic

func main() {

	//	books = append(books, Book{ID: "1", Name: "Kite Runner", Author: "Khalid", Quantity: 10})
	//	books = append(books, Book{ID: "2", Name: "The Alchemist", Author: "Paulo Coelho", Quantity: 15})

	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:9010", r))

}
