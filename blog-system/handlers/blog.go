package handlers

import (
	"net/http"

	"github.com/Samyog-G/Golang-Projects/blog-system/database"
	"github.com/Samyog-G/Golang-Projects/blog-system/models"
	"github.com/gin-gonic/gin"
)

func GetAllBlogPost(c *gin.Context) {
	var posts []models.BlogPost
	if err := database.DB.Find(&posts).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to fetch the blog posts"})
		return
	}
	c.JSON(http.StatusOK, posts)
}

func CreateBlogPost(c *gin.Context) {
	var input models.BlogPost
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid Input"})
		return
	}
	if err := database.DB.Create(&input).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to create the database"})
		return
	}
	c.JSON(http.StatusOK, input)
}

func UpdateBlogPost(c *gin.Context) {
	id := c.Param("id")
	var post models.BlogPost
	if err := database.DB.First(&post, id).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Post not found"})
		return
	}
	var input models.BlogPost
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid Input"})
		return
	}
	post.Title = input.Title
	post.Content = input.Content

	if err := database.DB.Save(&post).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to update"})
	}
	c.JSON(http.StatusOK, post)
}

func DeleteBlogPost(c *gin.Context) {
	id := c.Param("id")
	var post models.BlogPost
	if err := database.DB.First(&post, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Post not found"})
		return
	}
	if err := database.DB.Delete(&post).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Unable to delete"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Post deleted successfully"})
}
