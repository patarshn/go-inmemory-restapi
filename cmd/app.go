package main

import (
	"example/internal/users"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	userRepo := users.NewUserRepository()
	userService := users.NewUserService(userRepo)
	userHandler := users.NewUserHandler(userService)

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	users.Router(r, userHandler)
	r.Run()
}
