package routes

import (
	controller "github.com/pi-prakhar/jwt-auth-go/controllers"
	middleware "github.com/pi-prakhar/jwt-auth-go/middleware"

	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.Use(middleware.Authenticate())
	incomingRoutes.GET("/users", controller.GetUsers())
	incomingRoutes.GET("/users/:user_id", controller.GetUser())
}
