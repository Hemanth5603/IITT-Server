package auth_handlers

import (
	"fmt"

	"github.com/Hemanth5603/IITT-Server/auth_utils"
	"github.com/Hemanth5603/IITT-Server/models"
	"github.com/Hemanth5603/IITT-Server/utils"
	"github.com/gofiber/fiber/v2"
	"gopkg.in/mail.v2"
)

func ResetPasswordEmail(ctx *fiber.Ctx) error {
	var payload models.SendMailRequest

	if err := ctx.BodyParser(&payload); err != nil {
		return ctx.Status(fiber.StatusBadRequest).
			JSON(fiber.Map{"status": "false", "error": err.Error()})
	}
	otp := auth_utils.GenerateOTP()
	from := "outreach@iittnif.com"
	password := "eusfgdljcklenfmi"
	smtpHost := "smtp.gmail.com"
	smtpPort := 587

	_, status, err := utils.DBCheckUserExists(payload.To)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).
			JSON(fiber.Map{"status": "false", "error": err.Error()})
	}

	if status == 404 {
		return ctx.Status(fiber.StatusNotFound).
			JSON(fiber.Map{"status": "false", "error": err, "message": "User Does not exist"})
	}

	// imagePath := "assets/iittnmicps.png"
	// imageData, err := ioutil.ReadFile(imagePath)
	// if err != nil {
	// 	fmt.Println("Error reading image file:", err)
	// 	return ctx.Status(fiber.StatusBadRequest).
	// 		JSON(fiber.Map{"status": "false", "error": err.Error()})
	// }

	//base64Image := base64.StdEncoding.EncodeToString(imageData)

	m := mail.NewMessage()

	m.SetHeader("From", from)
	m.SetHeader("To", payload.To)
	m.SetHeader("Subject", "IITTNiF")
	m.SetBody("text/html", fmt.Sprintf(`
		<html>
			<body style="font-family: Arial, sans-serif; text-align: center;">
				<h2>Hello, Welcome to IITTNiF app</h2>
				<p>Please use this OTP to verify your email address and Reset your password</p>
				<div style="margin: 20px auto; padding: 10px 20px; background-color: #007BFF; color: white; font-size: 24px; font-weight: bold; display: inline-block; border-radius: 5px;">
					%s
				</div>
			</body>
		</html>`, otp))

	d := mail.NewDialer(smtpHost, smtpPort, from, password)
	if err := d.DialAndSend(m); err != nil {
		return err
	}

	err = auth_utils.DBHandleOTP(payload.Token, otp)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).
			JSON(fiber.Map{"status": "false", "error": err.Error()})
	}

	return ctx.Status(fiber.StatusOK).
		JSON(fiber.Map{"status": "true", "message": "OTP Sent Successfully", "otp": otp})

}
