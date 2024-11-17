package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type book struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Quantity int    `json:"quantity"`
}

var books = []book{
	{ID: "1", Title: "The Alchemist", Author: "Paulo Coelho", Quantity: 10},
	{ID: "2", Title: "Ikigai", Author: "Unknown", Quantity: 3},
	{ID: "3", Title: "Rich Dad Poor Dad", Author: "Someone", Quantity: 5},
}

// Code to get all the books
func getBook(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}

func bookID(c *gin.Context) {
	id := c.Param("id")
	book, err := getBookByID(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found"})
		return
	}
	c.IndentedJSON(http.StatusOK, book)
}

// Code to get a book by its id
func getBookByID(id string) (*book, error) {
	for i, b := range books {
		if b.ID == id {
		}
		return &books[i], nil
	}
	return nil, errors.New("Book not found")
}

// Code or checking out a book
func checkoutBook(c *gin.Context) {
	id, ok := c.GetQuery("id")

	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Missing id parameter"})
		return
	}
	book, err := getBookByID(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found"})
		return
	}

	if book.Quantity <= 1 {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found"})
		return
	}
	book.Quantity -= 1
	c.IndentedJSON(http.StatusOK, book)
}

// code to return the book
func returnBook(c *gin.Context) {
	id, ok := c.GetQuery("id")

	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Missing id parameter"})
		return
	}
	book, err := getBookByID(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found"})
		return
	}

	if book.Quantity <= 1 {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found"})
		return
	}
	book.Quantity += 1
	c.IndentedJSON(http.StatusOK, book)
}

// Code to add a new book to the list
func createBook(c *gin.Context) {
	var newBook book

	if err := c.BindJSON(&newBook); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	books = append(books, newBook)
	c.IndentedJSON(http.StatusCreated, newBook)

}

// Router
func main() {
	router := gin.Default()
	// router.SetTrustedProxies([]string{"127.0.0.1"})
	router.GET("/books", getBook)
	router.POST("/books", createBook)
	router.PATCH("/checkout", checkoutBook)

	router.GET("/books/:id", bookID)
	router.PATCH("/return", returnBook)

	router.Run("localhost:8081")
}
