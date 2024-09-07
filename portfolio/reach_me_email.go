package portfolio

import (
	"crypto/tls"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"gopkg.in/mail.v2"
)

func SendReachMeEmail(ctx *fiber.Ctx) error {
	var payload ReachMeRequest

	if err := ctx.BodyParser(&payload); err != nil {
		return ctx.Status(fiber.StatusBadRequest).
			JSON(fiber.Map{"status": "false", "error": err.Error()})
	}
	from := "shemanth5603@gmail.com"
	password := "wlrnybpgebxnqvgl"
	smtpHost := "smtp.gmail.com"
	smtpPort := 587

	m := mail.NewMessage()
	m.SetHeader("From", from)
	m.SetHeader("To", "shemanth.kgp@gmail.com")
	m.SetHeader("Subject", fmt.Sprintf("CoderXOP Reach Me - Name - %s", payload.Name))
	m.SetBody("text/html", fmt.Sprintf(`
		<html>
			<body>
				<h3>%s</h3>
				<h4>%s</h4>
			</body>
		</html>`, payload.Email, payload.Message))

	d := mail.NewDialer(smtpHost, smtpPort, from, password)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	if err := d.DialAndSend(m); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).
			JSON(fiber.Map{"status": "false", "error": "Failed to send email", "details": err.Error()})
	}

	return ctx.Status(fiber.StatusOK).
		JSON(fiber.Map{"status": "true", "message": "CoderXOP reached Successfully"})

}
