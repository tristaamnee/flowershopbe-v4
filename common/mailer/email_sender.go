package mailer

import (
	"errors"
	"net/smtp"
)

func (m *mailer) EmailSender(receiver string, subject string, content string) error {
	shopEmailUsername := m.cfg.GmailEmail
	shopEmailPassword := m.cfg.GmailEmailPass
	if shopEmailUsername == "" {
		return errors.New("shop Gmail variable not set")
	}

	smtpHost := m.cfg.SMTPHOST
	smtpPort := m.cfg.SMTPPORT

	auth := smtp.PlainAuth("", shopEmailUsername, shopEmailPassword, smtpHost)

	msg := []byte("To: " + receiver + "\r\n" + "Subject: " + subject + "\r\n" + content)
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, shopEmailUsername, []string{receiver}, msg)
	if err != nil {
		return err
	}
	return nil
}
