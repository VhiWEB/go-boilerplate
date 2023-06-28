package maincontrollers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	_ "main/internal/auth"
)

func getServiceDetail(c *gin.Context) {
	c.Data(http.StatusOK, "text/html; charset=utf-8", []byte("<html>Go Boilerplate</html>"))
}
