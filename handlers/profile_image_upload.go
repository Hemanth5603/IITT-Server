package handlers

import (
	"fmt"
	"path/filepath"

	"github.com/Hemanth5603/IITT-Server/helpers"
	"github.com/Hemanth5603/IITT-Server/models"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func ProfileImageUpload(ctx *fiber.Ctx) error {
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

	for _, file := range files {
		binaryFile, err := helpers.FileToByteArray(file)
		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).
				JSON(fiber.Map{"status": false, "error": err.Error(), "message": "binary conversion"})
		}

		tuuid := uuid.New()
		idFile = "/profiles/" + tuuid.String() + filepath.Ext(file.Filename)

		err = helpers.UploadFile(idFile, binaryFile)

		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).
				JSON(fiber.Map{"status": false, "error": err.Error(), "msg": "uploading file err"})
		}
		image = file.Filename
		fmt.Printf(image)
		fmt.Printf(idFile)

	}

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": false, "error": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": true, "message": "Profile Image Upload Sucessfull",
	})
}
