package routes

import (
	"main/internal/main/controllers/maincontrollers"
	"main/internal/main/controllers/usercontrollers"

	"github.com/gin-gonic/gin"

	"main/internal/auth"
)

func init() {
	route := gin.Default()

	MainRoutes(route)
	AuthRoutes(route)
	UserRoutes(route)

	route.Run(":3000")
}

func MainRoutes(route *gin.Engine) {
	route.GET("/", maincontrollers.GetServiceDetail)
}

func AuthRoutes(route *gin.Engine) {
	authV1 := route.Group("/api/v1")
	{
		authV1.POST("/auth", usercontrollers.Login)
	}
}

func UserRoutes(route *gin.Engine) {
	userV1 := route.Group("/api/v1").Use(auth.UserAuthMiddleware())
	{
		userV1.GET("/me", usercontrollers.GetProfile)
	}
}
