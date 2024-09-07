package davidhandlers

import (
	"bytes"
	"encoding/json"
	"io"
	"mime/multipart"
	"net/http"

	davidutils "github.com/Hemanth5603/IITT-Server/david_utils"
	"github.com/gofiber/fiber/v2"
)

func FindFace(ctx *fiber.Ctx) error {
	// Parse the multipart form data
	form, err := ctx.MultipartForm()
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).
			JSON(fiber.Map{"status": "false", "error": err.Error()})
	}

	// Extract image file from form data
	imageFiles := form.File["image"]
	if len(imageFiles) == 0 {
		return ctx.Status(fiber.StatusBadRequest).
			JSON(fiber.Map{"status": "false", "error": "Image file is required"})
	}

	// Read the image file
	imageFile, err := imageFiles[0].Open()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).
			JSON(fiber.Map{"status": "false", "error": err.Error()})
	}
	defer imageFile.Close()

	// Prepare form data for the external request
	var b bytes.Buffer
	writer := multipart.NewWriter(&b)
	part, err := writer.CreateFormFile("image", imageFiles[0].Filename)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).
			JSON(fiber.Map{"status": "false", "error": err.Error()})
	}

	if _, err = io.Copy(part, imageFile); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).
			JSON(fiber.Map{"status": "false", "error": err.Error()})
	}

	// Add other form fields if necessary
	if err := writer.WriteField("msg", "Image Data Sent!!"); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).
			JSON(fiber.Map{"status": "false", "error": err.Error()})
	}

	// Close the writer to set the terminating boundary
	writer.Close()

	// Send the POST request to the external server
	resp, err := http.Post("https://6906-2409-40f0-1f-68b2-e524-ee8-5046-b55c.ngrok-free.app/upload-image", writer.FormDataContentType(), &b)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": "false",
			"error":  "Failed to send request to external server",
		})
	}
	defer resp.Body.Close()

	// Decode the response from the external server
	var response map[string]inerface{}
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": "false",
			"error":  "Failed to decode response from external server",
		})
	}

	//return ctx.JSON(response)

	student, err := davidutils.DBFetchStudentRoll(rollNumber)
}
