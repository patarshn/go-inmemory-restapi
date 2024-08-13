package users

import "github.com/gin-gonic/gin"

func Router(r *gin.Engine, handler IUserHandler) {
	users := r.Group("users")
	{
		users.GET("/:id", handler.ReadHandler)
		users.GET("/", handler.ReadAllHandler)
		users.POST("/", handler.CreateHandler)
		users.PUT("/", handler.UpdateHandler)
		users.DELETE("/:id", handler.DeleteHandler)
	}
}
