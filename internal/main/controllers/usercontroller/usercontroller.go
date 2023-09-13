package usercontroller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"main/internal/main/models/users"

	"main/pkg/jwt"
)

func GetAllUsers(c *gin.Context) {
	var users = users.Get()

	c.JSON(http.StatusOK, gin.H{"data": users})
}

func Register(c *gin.Context) {
	var requestBody struct {
		Username string `json:"username" form:"username"`
		Email    string `json:"email" form:"email"`
		Password string `json:"password" form:"password"`
	}

	if err := c.Bind(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	var user users.User
	user.Email = requestBody.Email
	user.Password = requestBody.Password

	_, err := user.Create()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": &users.RegisterFailed{},
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": gin.H{
			"success": true,
			"message": "Your account is successfully created!",
		},
	})
}

func Login(c *gin.Context) {
	var requestBody struct {
		Email    string `json:"email" form:"email"`
		Password string `json:"password" form:"password"`
	}

	if err := c.Bind(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	correct, user, _ := users.Authenticate(requestBody.Email, requestBody.Password)
	if !correct {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"message": &users.WrongUsernameOrPasswordError{},
		})
		return
	}

	token, err := jwt.GenerateToken(user.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": gin.H{
			"access_token": token,
			"user":         user,
		},
	})
}

func GetProfile(c *gin.Context) {
	user, exists := c.Get("user")

	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "No user found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": user,
	})
}
