package auth_handlers

import (
	"fmt"

	"github.com/Hemanth5603/IITT-Server/auth_utils"
	"github.com/Hemanth5603/IITT-Server/models"
	"github.com/gofiber/fiber/v2"
)

func VerifyOtp(ctx *fiber.Ctx) error {
	var payload models.VerifyOtpRequest
	fmt.Println("Called Verify OTP method")
	if err := ctx.BodyParser(&payload); err != nil {
		return ctx.Status(fiber.StatusBadRequest).
			JSON(fiber.Map{"status": "false", "error": err.Error()})
	}

	status, err := auth_utils.DBHandleVerifyOtp(payload.Token, payload.Otp)

	if err != nil {
		return ctx.Status(fiber.StatusNotFound).
			JSON(fiber.Map{"status": "false", "error": err.Error(), "message": "OTP Expired"})
	}

	if status == 400 {
		return ctx.Status(fiber.StatusBadRequest).
			JSON(fiber.Map{"status": "false", "message": "Incorrect OTP"})
	}

	return ctx.Status(fiber.StatusOK).
		JSON(fiber.Map{"status": "true", "message": "OTP Verification Successfull"})

}
