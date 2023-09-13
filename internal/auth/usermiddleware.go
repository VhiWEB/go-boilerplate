package auth

import (
	"context"
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
		bearerToken := strings.Split(header, "Bearer ")

		if len(bearerToken) != 2 {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "No Bearer token found"})
		}

		id, err := jwt.ParseToken(bearerToken[1])
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err})
			return
		}

		// Check if user exists in db
		user, tx := users.GetById(id)
		if tx.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err})
			return
		}

		c.Set("user", user)

		c.Next()
	}
}

// ForContext finds the user from the context. REQUIRES Middleware to have run.
func ForContext(ctx context.Context) *users.User {
	raw, _ := ctx.Value(userCtxKey).(*users.User)
	return raw
}
