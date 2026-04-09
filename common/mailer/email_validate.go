package mailer

import (
	"errors"
	"net/mail"
)

func EmailValidate(email string) error {
	addr, err := mail.ParseAddress(email)
	if err != nil {
		return err
	}

	if addr.Name != "" && addr.Name != email {
		return errors.New("invalid email format")
	}
	return nil
}
