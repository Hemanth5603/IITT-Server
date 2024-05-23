package main

import (
	"github.com/Hemanth5603/IITT-Server/infrastructure"
	"github.com/Hemanth5603/IITT-Server/workers"
	"github.com/gofiber/fiber/v2"
)

func main() {

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	workers.Routes(app)

	infrastructure.InitializePostgresSQL()
	infrastructure.InitializeSpaces()

	app.Listen(":3000")

}
