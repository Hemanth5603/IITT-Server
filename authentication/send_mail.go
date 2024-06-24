package authentication

import (
	"fmt"

	"github.com/Hemanth5603/IITT-Server/models"
	"github.com/gofiber/fiber/v2"
	"gopkg.in/mail.v2"
)

func SendEmail(ctx *fiber.Ctx) error {
	var payload models.SendMailRequest

	if err := ctx.BodyParser(&payload); err != nil {
		return ctx.Status(fiber.StatusBadRequest).
			JSON(fiber.Map{"status": "false", "error": err.Error()})
	}
	otp := generateOTP()
	from := "shemanth.kgp@gmail.com"
	password := "cpcrwjzjuyndskol"
	smtpHost := "smtp.gmail.com"
	smtpPort := 587

	m := mail.NewMessage()
	m.SetHeader("From", from)
	m.SetHeader("To", payload.To)
	m.SetHeader("Subject", "IITTNiF")
	m.SetBody("text/plain", fmt.Sprintf("Here is the OTP Verification Code : %s", otp))

	d := mail.NewDialer(smtpHost, smtpPort, from, password)
	if err := d.DialAndSend(m); err != nil {
		return err
	}

	err := DBHandleOTP(payload.Token, otp)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).
			JSON(fiber.Map{"status": "false", "error": err.Error()})
	}

	return ctx.Status(fiber.StatusOK).
		JSON(fiber.Map{"status": "true", "message": "OTP Sent Successfully"})
}
