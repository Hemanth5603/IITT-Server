package authentication

import (
	"github.com/Hemanth5603/IITT-Server/models"
	"github.com/gofiber/fiber/v2"
)

func VerifySms(ctx *fiber.Ctx) error {
	var payload models.VerifyData

	if err := ctx.BodyParser(&payload); err != nil {
		return ctx.Status(fiber.StatusBadRequest).
			JSON(fiber.Map{"status": "false", "error": err.Error()})
	}

	// newData := models.VerifyData{
	// 	User: payload.User,
	// 	Code: payload.Code,
	// }

	// err := twilioVerifyOTP(newData.User.PhoneNumber, newData.Code)
	// if err != nil {
	// 	return ctx.Status(fiber.StatusBadRequest).
	// 		JSON(fiber.Map{"status": "false", "error": err.Error()})
	// }

	return ctx.Status(fiber.StatusOK).
		JSON(fiber.Map{"status": "true", "messsage": "OTP Verified"})
}
