package database

import (
	"fmt"
	"go-api/utils"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	dbUser := utils.UseEnv("DB_USER", "postgres")
	dbPassword := utils.UseEnv("DB_PASSWORD", "postgres")
	dbHost := utils.UseEnv("DB_HOST", "localhost")
	dbPort := utils.UseEnv("DB_PORT", "5432")
	dbName := utils.UseEnv("DB_NAME", "postgres")
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		dbUser,
		dbPassword,
		dbHost,
		dbPort,
		dbName,
	)
	
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Error to try to connect in database")
	}

	DB = db
	fmt.Println("Success to connect into Database!")
}