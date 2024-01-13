package helpers

import (
	"gopkg.in/gomail.v2"
	"log"
)

const CONFIG_SMTP_HOST = "smtp.gmail.com"
const CONFIG_SMTP_PORT = 587
const CONFIG_SENDER_NAME = "No Reply - Sesuai App <sesuai.application@gmail.com>"
const CONFIG_AUTH_EMAIL = "sesuai.application@gmail.com"
const CONFIG_AUTH_PASSWORD = "wpnygwlymohalhrj"

func SendEmail(file, emailReceiver, fileName string) (success bool, err error) {
	mailer := gomail.NewMessage()
	mailer.SetHeader("From", CONFIG_SENDER_NAME)
	mailer.SetHeader("To", "faridoesnt@gmail.com")
	mailer.SetHeader("Subject", fileName)
	mailer.Attach(file)

	dialer := gomail.NewDialer(
		CONFIG_SMTP_HOST,
		CONFIG_SMTP_PORT,
		CONFIG_AUTH_EMAIL,
		CONFIG_AUTH_PASSWORD,
	)

	err = dialer.DialAndSend(mailer)
	if err != nil {
		log.Fatal(err.Error())
		return false, err
	}

	return true, nil
}
