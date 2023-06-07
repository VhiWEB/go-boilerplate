package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func init() {
	router := gin.Default()

	router.GET("/", getServiceDetail)

	router.Run("localhost:3000")
}

func getServiceDetail(c *gin.Context) {
	c.Data(http.StatusOK, "text/html; charset=utf-8", []byte("<html>Go Boilerplate</html>"))
}
