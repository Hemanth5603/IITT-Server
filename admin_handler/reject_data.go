package adminhandler

import (
	adminmodels "github.com/Hemanth5603/IITT-Server/admin_models"
	adminutils "github.com/Hemanth5603/IITT-Server/admin_utils"
	"github.com/gofiber/fiber/v2"
)

func RejectData(ctx *fiber.Ctx) error {
	var payload adminmodels.ApproveDataRequest

	if err := ctx.BodyParser(&payload); err != nil {
		return ctx.Status(fiber.StatusBadRequest).
			JSON(fiber.Map{"status": "false", "error": err.Error()})
	}

	err := adminutils.DBRejectData(payload.DataId)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).
			JSON(fiber.Map{"status": false, "error": err})
	}

	return ctx.Status(fiber.StatusOK).
		JSON(fiber.Map{"status": true, "message": "Data Rejected"})

}
