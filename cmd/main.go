package main

import (
	"go-api/initializers"
	"go-api/utils"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
}

func main() {
	server:= gin.Default()

	server.GET("ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H {
			"message": "pong",
		})
	})

	port := utils.UseEnv("PORT", "8080")

	server.Run(":" + port)
}