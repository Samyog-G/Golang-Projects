package main

import (
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

func getBook(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}

func createBook(c *gin.Context) {
	var newBook book

	if err := c.BindJSON(&newBook); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	books = append(books, newBook)
	c.IndentedJSON(http.StatusCreated, newBook)

}
func main() {
	router := gin.Default()
	router.SetTrustedProxies([]string{"127.0.0.1"})
	router.GET("/books", getBook)
	router.POST("/newbooks", createBook)
	router.Run("localhost:8081")
	// router.GET()
}
