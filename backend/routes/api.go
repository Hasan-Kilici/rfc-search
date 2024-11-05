package routes

import (
	"rfc/handlers"
	"github.com/gofiber/fiber/v2"
)

func Api(app fiber.Router) {
	app.Get("/search", handlers.Search)
}