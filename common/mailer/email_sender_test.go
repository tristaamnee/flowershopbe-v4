package mailer

import (
	"os"
	"testing"
)

func TestEmailSender(t *testing.T) {
	// Cài đặt biến môi trường tạm thời để test (nếu bạn chưa set trong hệ thống)
	// Lưu ý: Thay thông tin thật của bạn vào đây
	os.Setenv("GMAIL_EMAIL", "yuiqwevn@gmail.com")
	os.Setenv("GMAIL_EMAIL_PASSWORD", "qrpjntlgoqlfpwoy")

	receiver := "tritam20082001@gmail.com" // Email nhận bài test
	subject := "Test Golang Email"
	content := "<h1>Chào bạn!</h1><p>Đây là nội dung gửi từ hàm test trong Go.</p>"

	err := EmailSender(receiver, subject, content)

	if err != nil {
		t.Errorf("EmailSender failed: %v", err)
	} else {
		t.Log("Email sent successfully!")
	}
}
