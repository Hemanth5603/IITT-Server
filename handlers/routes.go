package handlers

import (
	adminhandler "github.com/Hemanth5603/IITT-Server/admin_handler"
	"github.com/Hemanth5603/IITT-Server/auth_handlers"
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
	routes.Post("/iitt/sendSms", auth_handlers.SendSms)
	routes.Post("/iitt/verifySms", auth_handlers.VerifySms)
	routes.Post("/iitt/sendEmail", auth_handlers.SendEmail)
	routes.Post("/iitt/verifyOtp", auth_handlers.VerifyOtp)
	routes.Get("/iitt-admin/fetch-data", adminhandler.FetchAllUnApprovedData)
	routes.Post("/iitt-admin/approve-data", adminhandler.ApproveData)
	routes.Post("/iitt-admin/reject-data", adminhandler.RejectData)

}
