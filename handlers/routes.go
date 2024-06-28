package handlers

import (
	"github.com/Hemanth5603/IITT-Server/admin_handler"
	"github.com/Hemanth5603/IITT-Server/authentication"
	"github.com/gofiber/fiber/v2"
)

func Routes(routes *fiber.App) {
	routes.Post("/uploadImage", uploadData)
	routes.Post("/iitt/register", RegisterUser)
	routes.Post("/iitt/login", Login)
	routes.Get("/iitt/getUserUploads/:user_id", GetUploads)
	routes.Post("/iitt/getLeaderBoard", GetLeaderBoard)
	routes.Post("/iitt/getUser", GetUser)
	routes.Post("/iitt/updateProfile", UpdateProfile)
	//routes.Post("/iitt/otp", authentication.SendSms)
	routes.Post("/iitt/sendEmail", authentication.SendEmail)
	routes.Post("/iitt/verifyOtp", authentication.VerifyOtp)
	routes.Post("/iitt-admin/fetch-data", admin_handler.FetchAllUnApprovedData)

}
