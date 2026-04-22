package mailer

import (
	"context"

	"github.com/tristaamne/flowershopbe-v4/common/config"
	"github.com/tristaamne/flowershopbe-v4/common/security/otp"
)

type mailer struct {
	cfg *config.Config
	otp otp.OTP
}

type Mailer interface {
	OTPSender(ctx context.Context, userEmail string) error
	EmailSender(receiver string, subject string, content string) error
	EmailValidate(email string) error
}

func NewMailer(cfg *config.Config, otp otp.OTP) Mailer {
	return &mailer{
		cfg: cfg,
		otp: otp,
	}
}
