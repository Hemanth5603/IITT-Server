package workers

import (
	"github.com/Hemanth5603/IITT-Server/models"
	"github.com/Hemanth5603/IITT-Server/utils"
	"github.com/gofiber/fiber/v2"
)

func RegisterUser(ctx *fiber.Ctx) error {
	var payload models.SignUpRequest

	if err := ctx.BodyParser(&payload); err != nil {
		return ctx.Status(fiber.StatusBadRequest).
			JSON(fiber.Map{"status": "false", "error": err.Error()})
	}

	hashedPassword, err := utils.HashPassword(payload.Password)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).
			JSON(fiber.Map{"status": "false", "err": err.Error()})
	}

	newUser := models.SignUpRequest{
		Name:     payload.Name,
		Email:    payload.Email,
		Password: hashedPassword,
	}

	id, err := utils.InsertUser(newUser)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).
			JSON(fiber.Map{"status": "false", "error": err.Error()})
	}

	return ctx.Status(fiber.StatusOK).
		JSON(fiber.Map{"status": "true", "id": id, "name": payload.Name})

}
