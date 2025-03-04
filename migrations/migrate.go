package migrations

import (
	"fmt"
	"go-api/utils"
	"log"
	"os"
	"path/filepath"

	"github.com/golang-migrate/migrate/v4"
	migratepostgres "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	gormpostgres "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const migrationDir = "db/migrations"

func HandleMigration(args []string) {
	command := args[0]

	switch command {
	case "create":
		if len(args) < 2 {
			log.Fatal("The name of migrations is mandatory!")
		}
		createMigration(args[1])
	case "up":
		migrateUp()
	case "down":
		migrateDown()
	default:
		fmt.Println("Invalid command. Use: create, up ou down")
	}
}

func createMigration(name string) {
	if _, err := os.Stat(migrationDir); os.IsNotExist(err) {
		os.MkdirAll(migrationDir, os.ModePerm)
	}

	timestamp := fmt.Sprintf("%d", os.Getpid())
	upFile := filepath.Join(migrationDir, fmt.Sprintf("%s_%s.up.sql", timestamp, name))
	downFile := filepath.Join(migrationDir, fmt.Sprintf("%s_%s.down.sql", timestamp, name))

	if err := os.WriteFile(upFile, []byte("-- UP MIGRATION\n"), 0644); err != nil {
		log.Fatal(err)
	}
	if err := os.WriteFile(downFile, []byte("-- DOWN MIGRATION\n"), 0644); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Migração criada: %s e %s\n", upFile, downFile)
}

func getMigrateInstance() *migrate.Migrate {
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
	db, err := gorm.Open(gormpostgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error to connect in database:", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("Error to obtain database:", err)
	}

	driver, err := migratepostgres.WithInstance(sqlDB, &migratepostgres.Config{})
	if err != nil {
		log.Fatal("Error to configure driver:", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://"+migrationDir,
		"postgres", driver)
	if err != nil {
		log.Fatal("Error to create migration:", err)
	}

	return m
}

func migrateUp() {
	m := getMigrateInstance()
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal("Error to apply migration:", err)
	}
	fmt.Println("Success to apply migrations!")
}

func migrateDown() {
	m := getMigrateInstance()
	if err := m.Down(); err != nil && err != migrate.ErrNoChange {
		log.Fatal("Error to revert migration:", err)
	}
	fmt.Println("Success to revert migrations!")
}