package adminhandler

import (
	adminutils "github.com/Hemanth5603/IITT-Server/admin_utils"
	"github.com/Hemanth5603/IITT-Server/models"
	"github.com/gofiber/fiber/v2"
)

func FetchAllUnApprovedData(ctx *fiber.Ctx) error {

	var dataList []models.DataModel = []models.DataModel{}

	dataList, err := adminutils.DBFetchUnApprovedData()

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).
			JSON(fiber.Map{"status": "false", "error": err.Error(), "message": "Cannot Fetch data"})
	}
	return ctx.Status(fiber.StatusOK).
		JSON(fiber.Map{
			"status": "true", "data": dataList,
		})
}
