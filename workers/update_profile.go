package workers

import (
	"fmt"
	"path/filepath"

	"github.com/Hemanth5603/IITT-Server/models"
	"github.com/Hemanth5603/IITT-Server/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func UpdateProfile(ctx *fiber.Ctx) error {
	var payload models.ProfileUpdateRequest

	if err := ctx.BodyParser(&payload); err != nil {
		return ctx.Status(fiber.StatusBadRequest).
			JSON(fiber.Map{"status": "false", "error": err.Error()})
	}

	form, err := ctx.MultipartForm()
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).
			JSON(fiber.Map{"status": "false", "error": err.Error()})
	}

	var idFile string

	files := form.File["profileimage"]
	var image string

	if len(files) != 0 {
		for _, file := range files {
			binaryFile, err := utils.FileToByteArray(file)
			if err != nil {
				return ctx.Status(fiber.StatusBadRequest).
					JSON(fiber.Map{"status": false, "error": err.Error(), "message": "binary conversion"})
			}

			tuuid := uuid.New()
			idFile = "/profiles/" + tuuid.String() + filepath.Ext(file.Filename)

			err = utils.UploadFile(idFile, binaryFile)

			if err != nil {
				return ctx.Status(fiber.StatusBadRequest).
					JSON(fiber.Map{"status": false, "error": err.Error(), "msg": "uploading file err"})
			}
			image = file.Filename
			fmt.Printf(image)
			fmt.Printf(idFile)

		}
	} else {
		idFile = "Default"
	}

	err = utils.DBUpdateProfile(payload, idFile)

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": false, "error": err.Error(), "message": "Broke here",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": true, "message": "Profile Updated Successfully",
	})
}
