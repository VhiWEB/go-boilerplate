package auth

import (
	"net/http"

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
		if header == "" {
			c.Next()
			return
		}

		// Validate jwt token
		tokenStr := header
		email, err := jwt.ParseToken(tokenStr)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err})
			return
		}

		// check if user exists in db
		user := users.User{Email: email}
		id, err := users.GetUserIdByEmail(email)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err})
			return
		}
		user.ID = id
		c.Set("user", user)

		c.Next()
	}
}
