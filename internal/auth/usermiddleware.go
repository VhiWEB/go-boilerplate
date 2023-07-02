package auth

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"main/internal/main/models/users"
	"main/pkg/jwt"
)

var userCtxKey = &contextKey{"user"}

type contextKey struct {
	name string
}

func UserAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.Request.Header.Get("Authorization")

		// Allow unauthenticated users in
		// if header == "" {
		// 	c.Next()
		// 	return
		// }

		// Validate jwt token
		tokenStr := strings.Split(header, "Bearer ")
		email, err := jwt.ParseToken(tokenStr[1])
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err})
			return
		}

		// Check if user exists in db
		user, err := users.GetUserByEmail(email)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err})
			return
		}
		c.Set("user", user)

		c.Next()
	}
}
