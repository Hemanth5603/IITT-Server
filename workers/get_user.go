package workers

import (
	"database/sql"
	"errors"

	"github.com/Hemanth5603/IITT-Server/models"
	"github.com/Hemanth5603/IITT-Server/utils"
	"github.com/gofiber/fiber/v2"
)

func GetUser(ctx *fiber.Ctx) error {
	var payload models.GetUserRequest

	if err := ctx.BodyParser(&payload); err != nil {
		return ctx.Status(fiber.StatusBadRequest).
			JSON(fiber.Map{"status": "false", "error": err.Error()})
	}

	findUser, rank, err := utils.FindUserById(payload.Id)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return ctx.Status(fiber.StatusNotFound).
				JSON(fiber.Map{"status": "false", "error": "User not found"})
		}
		return ctx.Status(fiber.StatusInternalServerError).
			JSON(fiber.Map{"status": "false", "error": err.Error()})
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
			"rank":          rank,
			"location":      findUser.Location,
			"profile_image": findUser.ProfileImage,
		})

}
