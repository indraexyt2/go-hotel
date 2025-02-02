package helpers

import (
	"gopkg.in/gomail.v2"
	"os"
	"strconv"
)

type Email struct {
	To      string
	Subject string
	Body    string
}

func (e *Email) SendEmailNotification() error {
	message := gomail.NewMessage()
	message.SetHeader("From", "hello@demomailtrap.com")
	message.SetHeader("To", e.To)
	message.SetHeader("Subject", e.Subject)
	message.SetBody("text/html", e.Body)

	portStr := os.Getenv("MAILTRAP_PORT")
	portInt, _ := strconv.Atoi(portStr)
	dialer := gomail.Dialer{
		Host:     os.Getenv("MAILTRAP_HOST"),
		Port:     portInt,
		Username: os.Getenv("MAILTRAP_USERNAME"),
		Password: os.Getenv("MAILTRAP_PASSWORD"),
	}

	err := dialer.DialAndSend(message)
	if err != nil {
		return err
	}

	return nil
}
