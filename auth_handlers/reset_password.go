package auth_handlers

import (
	"github.com/Hemanth5603/IITT-Server/auth_utils"
	"github.com/Hemanth5603/IITT-Server/helpers"
	"github.com/Hemanth5603/IITT-Server/models"
	"github.com/Hemanth5603/IITT-Server/utils"
	"github.com/gofiber/fiber/v2"
)

func ResetPassword(ctx *fiber.Ctx) error {
	var payload models.LoginInRequest

	if err := ctx.BodyParser(&payload); err != nil {
		return ctx.Status(fiber.StatusBadRequest).
			JSON(fiber.Map{"status": "false", "error": err.Error()})
	}
	println(payload.Password)
	println(payload.Email)
	id, _, err := utils.DBCheckUserExists(payload.Email)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).
			JSON(fiber.Map{"status": "false", "error": err.Error()})
	}

	hashedPassword, err := helpers.HashPassword(payload.Password)

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).
			JSON(fiber.Map{"status": "false", "err": err.Error()})
	}
	println(id)
	println(hashedPassword)
	err = auth_utils.DBResetPassword(id, hashedPassword)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).
			JSON(fiber.Map{"status": "false", "error": err.Error()})
	}

	return ctx.Status(fiber.StatusOK).
		JSON(fiber.Map{"status": "true", "message": "Password Reset Succesfull"})

}
