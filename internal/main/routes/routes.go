package routes

import (
	"main/internal/main/controllers/maincontrollers"

	"github.com/gin-gonic/gin"

	"main/internal/auth"
	_ "main/internal/auth"
)

func init() {
	route := gin.Default()

	route.Use(auth.UserAuthMiddleware())

	route.GET("/", maincontrollers.getServiceDetail)

	route.Run("localhost:3000")
}
