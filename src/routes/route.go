package routes

import (
	"net/http"
	"url-shortner/src/modules/auth"
	"url-shortner/src/modules/user"

	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) *gin.Engine {
	v1 := router.Group("/api/v1/")

	// Register routes
	auth.RoutesHandler(v1)
	user.RoutesHandler(v1)

	router.NoRoute(func(c *gin.Context) {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"Status": http.StatusNotFound, "message": "Route not found"})
	})
	return router
}
