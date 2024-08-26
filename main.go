package main

import (
	"first_api/database"
	"first_api/entities"
	"first_api/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	// Conecta ao banco de dados
	database.Connect()

	// Migração automática do GORM
	database.DB.AutoMigrate(&entities.User{})

	// Configura as rotas
	routes.AppRoutes(app)

	// Inicia o servidor
	app.Listen(":8080")
}
