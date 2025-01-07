package routes

import (
	"github.com/Samyog-G/jwt-authentication/controllers"
	"github.com/Samyog-G/jwt-authentication/middleware"
	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.Use(middleware.Authenticate())
	incomingRoutes.GET("/users", controllers.GetUsers())
	incomingRoutes.GET("/users/:user_id", controllers.GetUserById())
}
