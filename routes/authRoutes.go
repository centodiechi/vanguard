package routes

import (
	"github.com/centodiechi/vanguard/internal/handlers"
	"github.com/gin-gonic/gin"
)

func RegisterAuthRoutes(router *gin.Engine) {
	auth := router.Group("/vanguard")
	{
		auth.POST("/createUser", handlers.CreateUser())
		auth.POST("/login", handlers.Login())

		// secured := auth.Group("/")
		// secured.Use(middleware.Authorize())
		// {
		// 	secured.POST("/logout", handlers.Logout)
		// }
	}
}
