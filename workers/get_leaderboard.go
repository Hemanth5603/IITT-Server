package workers

import (
	"fmt"

	"github.com/Hemanth5603/IITT-Server/models"
	"github.com/Hemanth5603/IITT-Server/utils"
	"github.com/gofiber/fiber/v2"
)

func GetLeaderBoard(ctx *fiber.Ctx) error {

	var payload models.GetLeaderboardRequest

	if err := ctx.BodyParser(&payload); err != nil {
		return ctx.Status(fiber.StatusBadRequest).
			JSON(fiber.Map{"status": "false", "error": err.Error()})
	}
	var userList []models.UserModel = []models.UserModel{}
	fmt.Println("category", payload.Category)

	userList, err := utils.FetchLeaderBoardFromDB(payload.Limit, payload.Category)

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": "false", "err": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "true", "leaderboard": userList,
	})

}
