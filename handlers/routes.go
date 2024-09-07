package handlers

import (
	adminhandler "github.com/Hemanth5603/IITT-Server/admin_handler"
	"github.com/Hemanth5603/IITT-Server/auth_handlers"
	davidhandlers "github.com/Hemanth5603/IITT-Server/david_handlers"
	"github.com/Hemanth5603/IITT-Server/portfolio"
	"github.com/gofiber/fiber/v2"
)

func Routes(routes *fiber.App) {
	routes.Post("/uploadImage", uploadData)
	routes.Post("/iitt/register", RegisterUser)
	routes.Post("/iitt/login", Login)
	routes.Get("/iitt/getUserUploads/:user_id", GetUploads)
	routes.Post("/iitt/getLeaderBoard", GetLeaderBoard)
	routes.Post("/iitt/getUser", GetUser)
	routes.Post("/iitt/deleteData", DeleteData)
	routes.Post("/iitt/updateProfile", UpdateProfile)
	routes.Post("/iitt/sendSms", auth_handlers.SendSms)
	routes.Post("/iitt/verifySms", auth_handlers.VerifySms)
	routes.Post("/iitt/sendEmail", auth_handlers.SendEmail)
	routes.Post("/iitt/verifyOtp", auth_handlers.VerifyOtp)
	routes.Delete("/iitt/expiredOtp", auth_handlers.HandleExpiredOtp)
	routes.Get("/iitt-admin/fetch-data", adminhandler.FetchAllUnApprovedData)
	routes.Post("/iitt-admin/approve-data", adminhandler.ApproveData)
	routes.Post("/iitt-admin/reject-data", adminhandler.RejectData)
	routes.Post("/iitt/resetPasswordEmail", auth_handlers.ResetPasswordEmail)
	routes.Post("/iitt/resetPassword", auth_handlers.ResetPassword)
	routes.Post("/iitt/deleteAccount", DeleteAccount)

	routes.Post("/coderxop/reachme", portfolio.SendReachMeEmail)

	//David Routes
	routes.Post("/david/register", davidhandlers.RegisterStudent)
	routes.Post("/david/findface", davidhandlers.FindFace)

}
