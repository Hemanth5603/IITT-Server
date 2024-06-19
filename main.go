package main

import (
	"github.com/Hemanth5603/IITT-Server/handlers"
	"github.com/Hemanth5603/IITT-Server/infrastructure"

	"github.com/gofiber/fiber/v2"
)

func main() {

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Post("/iitt/post", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).
			JSON(fiber.Map{"status": "true", "post": "Accepted"})
	})

	handlers.Routes(app)

	infrastructure.InitializePostgresSQL()
	infrastructure.InitializeSpaces()

	app.Listen(":8081")

}
