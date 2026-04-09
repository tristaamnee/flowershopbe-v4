package mailer

import (
	"context"
	"fmt"
)

func OTPSender(ctx context.Context, userEmail string) error {
	otpCode, err := GenerateOTP(6)
	if err != nil {
		return fmt.Errorf("error when generating OTP code: %v", err)
	}
	err = SaveOTP(ctx, userEmail, otpCode)
	if err != nil {
		return fmt.Errorf("error when saving OTP: %v", err)
	}
	go func() {
		subject := "Mã xác thực tài khoản trang Bé Bán Hoa"
		content := "Mã OTP của bạn là: " + otpCode
		_ = EmailSender(userEmail, subject, content)
	}()
	return nil
}
