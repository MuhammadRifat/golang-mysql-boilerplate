package user

import "github.com/gin-gonic/gin"

func RoutesHandler(router *gin.RouterGroup) {
	group := router.Group("/user")
	{
		// group.POST("/", UserController.CreateOne)
		group.GET("/", UserController.GetAll)
		group.GET("/:id", UserController.GetOne)
		group.PUT("/:id", UserController.UpdateOne)
		group.DELETE("/:id", UserController.DeleteOne)
	}
}
