package davidhandlers

import (
	"bytes"
	"encoding/json"
	"io"
	"mime/multipart"
	"net/http"

	davidmodels "github.com/Hemanth5603/IITT-Server/david_models"
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
	resp, err := http.Post("https://df95-2409-40f0-161-1a8-248b-21dc-1249-6dfe.ngrok-free.app/upload-image", writer.FormDataContentType(), &b)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": "false",
			"error":  "Failed to send request to external server",
		})
	}
	defer resp.Body.Close()
	print(resp.Body)

	if err := json.NewDecoder(resp.Body).Decode(&davidmodels.FindFaceResponse); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": "false",
			"error":  "Failed to decode response from external server",
		})
	}

	rollValue := davidmodels.FindFaceResponse.Roll
	print(rollValue)

	student, err := davidutils.DBFetchStudentRoll(rollValue)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).
			JSON(fiber.Map{"error": err.Error(), "status": "false"})
	}

	return ctx.Status(fiber.StatusOK).
		JSON(fiber.Map{
			"status":   "success",
			"Id":       student.Id,
			"Name":     student.Name,
			"Branch":   student.Branch,
			"Phone":    student.Phone,
			"Roll":     student.Roll,
			"Academic": student.AcademicYear,
			"Semester": student.Semester,
			"Section":  student.Section,
		})

}
