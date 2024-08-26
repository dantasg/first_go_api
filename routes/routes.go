package routes

import (
	"first_api/controllers"
	"first_api/database"
	"first_api/repositories"
	"first_api/webSocketManager"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"

	"github.com/gofiber/swagger"
)

func AppRoutes(app *fiber.App) {
	db := database.DB

	// Inicializa o ClientManager
	clientManager := webSocketManager.NewClientManager()

	userRepository := repositories.NewUserRepository(db)
	userController := controllers.NewUserController(userRepository, *clientManager)

	v1 := app.Group("/v1")

	// Rota para acessar a documentação do Swagger
	app.Get("/swagger/*", swagger.HandlerDefault)

	// Configuração do WebSocket
	v1.Get("/ws", websocket.New(webSocketManager.WebSocketHandler(clientManager)))

	// Other routes
	v1.Get("/users", userController.FindAll)
	v1.Post("/user", userController.Create)
	v1.Post("/login", userController.Login)
}
