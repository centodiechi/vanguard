package handlers

import (
	syncprovider "github.com/centodiechi/vanguard/sync-provider"
	"github.com/gin-gonic/gin"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

func CreateUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(400, gin.H{"error": "Invalid request"})
			return
		}

		err := syncprovider.ServiceProvider.IdtClient
		if err != nil {
			c.JSON(500, gin.H{"error": "Failed to create user"})
			return
		}

		c.JSON(200, gin.H{"message": "User created successfully"})
	}
}

func Login() gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}
