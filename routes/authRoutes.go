package routes

import (
	"github.com/gin-gonic/gin"
)

func RegisterAuthRoutes(router *gin.Engine) {
	auth := router.Group("/vanguard")
	{
		auth.POST("/register", internal.handlers.Register)
		auth.POST("/login", internal.handlers.Login)

		// secured := auth.Group("/")
		// secured.Use(middleware.Authorize())
		// {
		// 	secured.POST("/logout", handlers.Logout)
		// }
	}
}
