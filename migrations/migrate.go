package migrations

import (
	"fmt"
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
			log.Fatal("Nome da migração é obrigatório.")
		}
		createMigration(args[1])
	case "up":
		migrateUp()
	case "down":
		migrateDown()
	default:
		fmt.Println("Comando inválido. Use: create, up ou down")
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
	dsn := "postgres://user:password@localhost:5432/meubanco?sslmode=disable"
	db, err := gorm.Open(gormpostgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Erro ao conectar ao banco:", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("Erro ao obter DB:", err)
	}

	driver, err := migratepostgres.WithInstance(sqlDB, &migratepostgres.Config{})
	if err != nil {
		log.Fatal("Erro ao configurar driver:", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://"+migrationDir,
		"postgres", driver)
	if err != nil {
		log.Fatal("Erro ao criar migração:", err)
	}

	return m
}

func migrateUp() {
	m := getMigrateInstance()
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal("Erro ao aplicar migrações:", err)
	}
	fmt.Println("Migrações aplicadas com sucesso!")
}

func migrateDown() {
	m := getMigrateInstance()
	if err := m.Down(); err != nil && err != migrate.ErrNoChange {
		log.Fatal("Erro ao reverter migrações:", err)
	}
	fmt.Println("Migrações revertidas com sucesso!")
}