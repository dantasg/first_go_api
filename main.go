package main

import (
	"first_api/database"
	"first_api/entities"
	"first_api/routes"
	"first_api/utils"

	_ "first_api/docs"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

func init() {
	utils.LoadVariables()
}

func main() {

	app := fiber.New()

	// Conecta ao banco de dados
	database.Connect()

	// Migração automática do GORM
	database.DB.AutoMigrate(&entities.User{})

	// Configura as rotas
	routes.AppRoutes(app)

	// Adiciona a rota para a documentação Swagger
	app.Get("/swagger/*", swagger.HandlerDefault) // Serve a documentação Swagger

	// Inicia o servidor
	app.Listen(":8080")
}
