package handlers

import "github.com/gofiber/fiber/v2"

func Routes(routes *fiber.App) {
	routes.Post("/uploadImage", uploadData)
	routes.Post("/iitt/register", RegisterUser)
	routes.Post("/iitt/login", Login)
	routes.Get("/iitt/getUserUploads/:user_id", GetUploads)
	routes.Post("/iitt/getLeaderBoard", GetLeaderBoard)
	routes.Post("/iitt/getUser", GetUser)
	routes.Post("/iitt/updateProfile", UpdateProfile)

}
