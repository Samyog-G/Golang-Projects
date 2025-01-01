package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

var secretKey = []byte("your-secret-key")

type loginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type loginResponse struct {
	Token string `json:"token"`
}

func generateJWT(email, password string) (string, error) {
	claims := jwt.MapClaims{
		"email": email,
	}
	tokens := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokens.SignedString(secretKey)
	if err != nil {
		return "", err
	}
	return token, nil
}

func AuthenticateJWT(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")

	if tokenString == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Missing Token"})
		c.Abort()
		return
	}

	tokenString = strings.TrimPrefix(tokenString, "Bearer ")

	_, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("invalid signing method")
		}
		return secretKey, nil
	})

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid token"})
		c.Abort()
		return
	}

	c.Next()
}

func main() {
	r := gin.Default()
	r.POST("/login", func(c *gin.Context) {
		var req loginRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid Request"})
			return
		}
		token, err := generateJWT(req.Email, req.Password)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Error generating token"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"token": token})
	})

	r.GET("/protected", AuthenticateJWT, func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "You have access to the protected route"})
	})

	r.Run(":4545")
}
