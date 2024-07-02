package handlers

import (
	"database/sql"
	"errors"

	"github.com/Hemanth5603/IITT-Server/helpers"
	"github.com/Hemanth5603/IITT-Server/models"
	"github.com/Hemanth5603/IITT-Server/utils"
	"github.com/gofiber/fiber/v2"
)

func Login(ctx *fiber.Ctx) error {
	var payload models.LoginInRequest

	if err := ctx.BodyParser(&payload); err != nil {
		return ctx.Status(fiber.StatusBadRequest).
			JSON(fiber.Map{"status": "false", "error": err.Error()})
	}

	findUser, err := utils.FindUserByEmail(payload.Email)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return ctx.Status(fiber.StatusNotFound).
				JSON(fiber.Map{"status": "false", "error": "User not found"})
		}
		return ctx.Status(fiber.StatusInternalServerError).
			JSON(fiber.Map{"status": "false", "error": err.Error()})
	}

	match := helpers.DecryptPassword(findUser.Password, payload.Password)

	if !match {
		return ctx.Status(fiber.StatusConflict).
			JSON(fiber.Map{"status": "false", "error": "Incorrect Password"})
	}

	return ctx.Status(fiber.StatusOK).
		JSON(fiber.Map{
			"status":        "True",
			"id":            findUser.Id,
			"name":          findUser.Name,
			"email":         findUser.Email,
			"phone":         findUser.Phone,
			"dob":           findUser.Dob,
			"contributions": findUser.Contributions,
			"rank":          findUser.Rank,
			"location":      findUser.Location,
		})

}
