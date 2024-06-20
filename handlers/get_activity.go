package handlers

import (
	"github.com/Hemanth5603/IITT-Server/models"
	"github.com/Hemanth5603/IITT-Server/utils"
	"github.com/gofiber/fiber/v2"
)

func GetUploads(ctx *fiber.Ctx) error {
	id := ctx.Params("user_id")

	var dataList []models.DataModel = []models.DataModel{}

	dataList, err := utils.GetUploadsByUserID(id)

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).
			JSON(fiber.Map{"status": "false", "error": err.Error(), "message": "Cannot Fetch data"})
	}
	return ctx.Status(fiber.StatusOK).
		JSON(fiber.Map{
			"status": "true", "data": dataList,
		})
}
