package main

import (
	"github.com/Samyog-G/Golang-Projects/blog-system/database"
	"github.com/Samyog-G/Golang-Projects/blog-system/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	database.InitDB()

	r := gin.Default()

	r.POST("/login", handlers.LoginHandler)
	r.POST("/register", handlers.RegisterUser)
	r.GET("/posts", handlers.GetAllBlogPost)
	r.POST("/posts", handlers.CreateBlogPost)
	r.PUT("/posts/:id", handlers.UpdateBlogPost)
	r.DELETE("/posts/:id", handlers.DeleteBlogPost)

	r.Run(":7777")
}
