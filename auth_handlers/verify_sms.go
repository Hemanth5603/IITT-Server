package auth_handlers

import (
	"github.com/Hemanth5603/IITT-Server/auth_utils"
	"github.com/Hemanth5603/IITT-Server/models"
	"github.com/gofiber/fiber/v2"
)

func VerifySms(ctx *fiber.Ctx) error {
	var payload models.VerifyVonageOTP

	if err := ctx.BodyParser(&payload); err != nil {
		return ctx.Status(fiber.StatusBadRequest).
			JSON(fiber.Map{"status": "false", "error": err.Error()})
	}

	request_id, status, err := auth_utils.VerifyVonageOTP(payload.RequestID, payload.OTP)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).
			JSON(fiber.Map{"status": "false", "request_id": request_id, "status_code": status})
	}

	return ctx.Status(fiber.StatusOK).
		JSON(fiber.Map{"status": "true", "request_id": request_id, "status_code": status})
}
