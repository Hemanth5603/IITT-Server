package auth_handlers

import (
	"github.com/Hemanth5603/IITT-Server/auth_utils"
	"github.com/Hemanth5603/IITT-Server/models"
	"github.com/gofiber/fiber/v2"
)

func HandleExpiredOtp(ctx *fiber.Ctx) error {
	var payload models.DeleteExpiredOtpRequest

	if err := ctx.BodyParser(&payload); err != nil {
		return ctx.Status(fiber.StatusBadRequest).
			JSON(fiber.Map{"status": "false", "error": err.Error()})
	}

	err := auth_utils.DBHandleExpiredOtp(payload.Token)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).
			JSON(fiber.Map{"status": "false", "error": err.Error()})
	}

	return ctx.Status(fiber.StatusOK).
		JSON(fiber.Map{"status": "true", "message": "Deleted OTP Succesfully"})
}
