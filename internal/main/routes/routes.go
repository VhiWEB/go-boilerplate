package routes

import (
	"main/internal/main/controllers/maincontroller"
	"main/internal/main/controllers/usercontroller"

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
	route.GET("/", maincontroller.GetServiceDetail)
}

func AuthRoutes(route *gin.Engine) {
	authV1 := route.Group("/api/v1")
	{
		authV1.POST("/register", usercontroller.Register)
		authV1.POST("/auth", usercontroller.Login)
	}
}

func UserRoutes(route *gin.Engine) {
	userV1 := route.Group("/api/v1").Use(auth.UserAuthMiddleware())
	{
		userV1.GET("/me", usercontroller.GetProfile)
	}
}
