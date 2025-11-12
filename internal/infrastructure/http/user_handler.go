package http

import (
	"templateApiRestGo/internal/application"
	"templateApiRestGo/internal/domain"

	"github.com/gofiber/fiber/v2"
)

// 🧠 Handler HTTP: capa más externa, interactúa con el cliente.
// Usa el servicio (application layer) para ejecutar la lógica.

type UserHandler struct {
	service *application.UserService
}

// Constructor con inyección de dependencia
func NewUserHandler(service *application.UserService) *UserHandler {
	return &UserHandler{service: service}
}

// Endpoint POST: crea un usuario
func (h *UserHandler) CreateUser(c *fiber.Ctx) error {
	var user domain.User
	// Parsear el JSON recibido
	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "JSON inválido"})
	}

	// Llamar al servicio
	if err := h.service.CreateUser(&user); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(201).JSON(fiber.Map{
		"message": "Usuario creado correctamente",
		"user":    user,
	})
}
