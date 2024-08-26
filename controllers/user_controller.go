package controllers

import (
	"encoding/json"
	"first_api/entities"
	"first_api/repositories"
	"first_api/webSocketManager"
	"log"

	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

type UserController struct {
	repository    repositories.UserRepository
	clientManager webSocketManager.ClientManager
}

func NewUserController(repo repositories.UserRepository, cm webSocketManager.ClientManager) *UserController {
	return &UserController{repository: repo, clientManager: cm}
}

func (u *UserController) Login(c *fiber.Ctx) error {
	var userLogin entities.UserLogin

	if err := c.BodyParser(&userLogin); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	var user entities.User

	user, err := u.repository.FindOneByEmail(userLogin.Email)

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	fmt.Println("Login efetuado - ", user)

	if user.Password == userLogin.Password {
		user.Password = ""
		return c.Status(http.StatusOK).JSON(user)
	} else {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Email or password invalid!"})
	}
}

func (u *UserController) FindAll(c *fiber.Ctx) error {
	users, err := u.repository.FindAll()

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(http.StatusOK).JSON(users)
}

func (u *UserController) Create(c *fiber.Ctx) error {
	var user entities.User

	if err := c.BodyParser(&user); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	// Criptografa a senha
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to encrypt password"})
	}

	user.Password = string(hashedPassword)

	createdUser, err := u.repository.Create(user)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	createdUser.Password = ""

	eventMessage := webSocketManager.Message{
		Event: "connection",
		Data:  "Hello, World!",
	}

	messageBytes, err := json.Marshal(eventMessage)
	if err != nil {
		log.Println("Erro ao serializar a mensagem:", err)
	}

	u.clientManager.Broadcast([]byte(messageBytes))

	return c.Status(http.StatusCreated).JSON(createdUser)
}
