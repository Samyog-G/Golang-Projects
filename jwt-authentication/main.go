package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.Use(gin.Logger())
	routes.AuthRoutes(r)
	routes.UserRoutes(r)

	r.GET("/api-1", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"success": "Successfully accessed api-1"})
	})

	r.GET("/api-2", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Successfully accessed api-2"})
	})

	r.Run(":8888")

}
