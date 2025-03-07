package main

import (
	"fmt"
	"go-api/database"
	"go-api/database/migrations"
	"go-api/initializers"
	"go-api/utils"
	"os"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	database.ConnectDB()
}

func main() {
	if len(os.Args) > 1 && os.Args[1] == "migrate" {
		if len(os.Args) < 3 {
			fmt.Println("Use: go run main.go migrate [create|up|down] name_of_migration")
			os.Exit(1)
		}
		migrations.HandleMigration(os.Args[2:])
		return
	}

	server := gin.Default()

	server.GET("ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	port := utils.UseEnv("PORT", "8080")

	server.Run(":" + port)
}