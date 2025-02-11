package auth

import (
	"github.com/gin-gonic/gin"
)

func RoutesHandler(router *gin.RouterGroup) {
	auth := router.Group("/auth")

	auth.POST("/login", AuthController.LoginHandler)
	auth.POST("/register", AuthController.RegisterHandler)
}
