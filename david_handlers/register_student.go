package davidhandlers

import (
	davidmodels "github.com/Hemanth5603/IITT-Server/david_models"
	davidutils "github.com/Hemanth5603/IITT-Server/david_utils"
	"github.com/gofiber/fiber/v2"
)

func RegisterStudent(ctx *fiber.Ctx) error {
	var payload davidmodels.RegisterStudentRequest

	if err := ctx.BodyParser(&payload); err != nil {
		return ctx.Status(fiber.StatusBadRequest).
			JSON(fiber.Map{"status": "false", "error": err.Error()})
	}

	newStudent := davidmodels.RegisterStudentRequest{
		Name:         payload.Name,
		Roll:         payload.Roll,
		Phone:        payload.Phone,
		Branch:       payload.Branch,
		Section:      payload.Section,
		AcademicYear: payload.AcademicYear,
		Semester:     payload.Semester,
	}

	id, err := davidutils.DBInsertStudent(newStudent)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).
			JSON(fiber.Map{"status": "false", "error": err.Error()})
	}

	return ctx.Status(fiber.StatusOK).
		JSON(fiber.Map{"status": "true", "id": id})

}
