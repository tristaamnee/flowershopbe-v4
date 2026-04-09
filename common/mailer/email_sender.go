package mailer

import (
	"errors"
	"net/smtp"
	"os"
)

func EmailSender(receiver string, subject string, content string) error {
	shopEmailUsername := os.Getenv("GMAIL_EMAIL")
	shopEmailPassword := os.Getenv("GMAIL_EMAIL_PASSWORD")
	if shopEmailUsername == "" {
		return errors.New("Shop gmail variable not set")
	}

	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	auth := smtp.PlainAuth("", shopEmailUsername, shopEmailPassword, smtpHost)

	msg := []byte("To: " + receiver + "\r\n" + "Subject: " + subject + "\r\n" + content)
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, shopEmailUsername, []string{receiver}, msg)
	if err != nil {
		return err
	}
	return nil
}
