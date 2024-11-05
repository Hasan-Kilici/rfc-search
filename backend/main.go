package main

import (
	"rfc/middlewares"
	"rfc/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/goccy/go-json"
)

func main() {
	app := fiber.New(fiber.Config{
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	})

	app.Use(middlewares.Cors)
	app.Use(middlewares.Logger)
	app.Use(middlewares.Compress)
	app.Use(middlewares.Security)
	app.Use(middlewares.RateLimit)

	api := app.Group("/api")
	routes.Api(api)

	app.Use(middlewares.NotFound)

	err := app.Listen("127.0.0.1:4000")
	if err != nil {
		panic(err)
	}
}