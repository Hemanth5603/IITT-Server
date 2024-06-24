package authentication

import (
	"github.com/Hemanth5603/IITT-Server/models"
	"github.com/gofiber/fiber/v2"
)

func SendSms(ctx *fiber.Ctx) error {
	var payload models.OTPData

	if err := ctx.BodyParser(&payload); err != nil {
		return ctx.Status(fiber.StatusBadRequest).
			JSON(fiber.Map{"status": "false", "error": err.Error()})
	}

	newData := models.OTPData{
		PhoneNumber: payload.PhoneNumber,
	}

	// apiKey := "AIzaSyBna1QJkW4u4-cFyX_1cQIeMzZy887ZPtk"
	// phoneNumber := "+917997435603"
	// recaptchaToken := "" // You need to handle Recaptcha if required

	sessionInfo, err := twilioSendOTP(newData.PhoneNumber)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "false", "err": err,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "true", "message": "OTP sent Successfully", "info": sessionInfo,
	})

}
