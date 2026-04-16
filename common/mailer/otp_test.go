package mailer

import (
	"context"
	"testing"
	"time"

	"github.com/tristaamne/flowershopbe-v4/common/db"
)

func TestOTPSender(t *testing.T) {
	// Khởi tạo context với timeout để tránh test chạy vô hạn
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	db.InitRedis()
	// Giả lập dữ liệu đầu vào
	userEmail := "tritam20082001@gmail.com"

	// Chạy hàm test
	err := OTPSender(ctx, userEmail)

	// 1. Kiểm tra lỗi trả về
	if err != nil {
		t.Fatalf("OTPSender() error = %v, muốn không có lỗi", err)
	}

	// 2. Vì EmailSender chạy trong goroutine (go func),
	// nên mình đợi một chút để nó chạy xong (nếu cần log hoặc kiểm tra)
	time.Sleep(100 * time.Millisecond)

	t.Logf("Đã chạy OTPSender thành công cho email: %s", userEmail)
}
