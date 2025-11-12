package http

import "github.com/gofiber/fiber/v2"

// 🧠 Define las rutas del servidor y asigna los handlers.

func SetupRoutes(app *fiber.App, handler *UserHandler) {
	api := app.Group("/api")
	api.Post("/users", handler.CreateUser)
}
