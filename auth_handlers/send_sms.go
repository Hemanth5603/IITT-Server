package auth_handlers

import (
	"github.com/Hemanth5603/IITT-Server/auth_utils"
	"github.com/Hemanth5603/IITT-Server/models"
	"github.com/gofiber/fiber/v2"
)

func SendSms(ctx *fiber.Ctx) error {
	var payload models.SentVonageOTPRequest

	if err := ctx.BodyParser(&payload); err != nil {
		return ctx.Status(fiber.StatusBadRequest).
			JSON(fiber.Map{"status": "false", "error": err.Error()})
	}

	requestId, status, err := auth_utils.SendVonageOTP(payload.Phone)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "false", "request_id": requestId, "status_code": status,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "true", "request_id": requestId, "status_code": status,
	})

}
