package mailer

import (
	"context"
	"fmt"
)

func (m *mailer) OTPSender(ctx context.Context, userEmail string) error {
	otpCode, err := m.otp.GenerateOTP(6)
	if err != nil {
		return fmt.Errorf("error when generating OTP code: %v", err)
	}
	err = m.otp.SaveOTP(ctx, userEmail, otpCode)
	if err != nil {
		return fmt.Errorf("error when saving OTP: %v", err)
	}
	go func() {
		subject := "Mã xác thực tài khoản trang Bé Bán Hoa"
		content := "Mã OTP của bạn là: " + otpCode + " sẽ khả dụng trong 2 phút."
		err = m.EmailSender(userEmail, subject, content)
		if err != nil {
			fmt.Printf("error when sending email: %v", err)
			return
		}
	}()
	return nil
}
