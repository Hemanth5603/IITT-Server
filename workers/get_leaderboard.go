package workers

import (
	"github.com/Hemanth5603/IITT-Server/models"
	"github.com/Hemanth5603/IITT-Server/utils"
	"github.com/gofiber/fiber/v2"
)

func GetLeaderBoard(ctx *fiber.Ctx) error {

	var userList []models.UserModel = []models.UserModel{}

	userList, err := utils.FetchLeaderBoardFromDB()

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": "false", "err": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "true", "leaderboard": userList,
	})

}
