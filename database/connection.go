package database

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	// DSN (Data Source Name) para a conexão com o banco de dados
	dsn := "host=localhost user=" + dbUser + " password=" + dbPassword + " dbname=" + dbName + " port=5432 sslmode=disable"

	// Conecta ao banco de dados usando o driver do PostgreSQL
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Armazena a conexão na variável global DB
	DB = database
}
