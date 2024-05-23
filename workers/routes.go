package workers

import "github.com/gofiber/fiber/v2"

func Routes(routes *fiber.App) {
	routes.Post("/uploadImage", uploadImage)
	routes.Post("/iitt/register", RegisterUser)
	routes.Post("/iitt/login", Login)

}
