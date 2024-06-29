package handlers

import (
	"fmt"
	"path/filepath"

	"github.com/Hemanth5603/IITT-Server/helpers"
	"github.com/Hemanth5603/IITT-Server/models"
	"github.com/Hemanth5603/IITT-Server/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func uploadData(ctx *fiber.Ctx) error {
	var payload models.DataModelRequest
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

	files := form.File["image"]
	var image string

	for _, file := range files {
		binaryFile, err := helpers.FileToByteArray(file)
		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).
				JSON(fiber.Map{"success": false, "error": err.Error(), "message": "binary conversion"})
		}

		tuuid := uuid.New()
		idFile = "/iitt/" + tuuid.String() + filepath.Ext(file.Filename)

		err = helpers.UploadFile(idFile, binaryFile)
		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).
				JSON(fiber.Map{"success": false, "error": err.Error(), "msg": "uploading file err"})
		}
		image = file.Filename
		fmt.Printf(image)
		fmt.Printf(idFile)

	}
	data := models.DataModelRequest{
		Id:         payload.Id,
		Latitude:   payload.Latitude,
		Longitude:  payload.Longitude,
		Image:      idFile,
		Category:   payload.Category,
		Remarks:    payload.Remarks,
		Address:    payload.Address,
		Date:       payload.Date,
		Time:       payload.Time,
		IsApproved: payload.IsApproved,
	}

	err = utils.InsertData(data)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).
			JSON(fiber.Map{"success": false, "error": err})
	}

	return ctx.Status(fiber.StatusOK).
		JSON(fiber.Map{"status": "true", "message": "Data Insertion Successfull"})
}
