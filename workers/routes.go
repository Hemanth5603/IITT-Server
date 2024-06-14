package workers

import "github.com/gofiber/fiber/v2"

func Routes(routes *fiber.App) {
	routes.Post("/uploadImage", uploadImage)
	routes.Post("/iitt/register", RegisterUser)
	routes.Post("/iitt/login", Login)
	routes.Get("/iitt/getUserUploads/:user_id", GetUploads)
	routes.Get("/iitt/getLeaderBoard", GetLeaderBoard)
	routes.Post("/iitt/getUser", GetUser)
	routes.Post("/iitt/uploadProfile", ProfileImageUpload)

}
