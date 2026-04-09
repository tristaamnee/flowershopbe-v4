package mailer

import (
	"context"
	"crypto/rand"
	"errors"
	"io"
	"time"

	"github.com/tristaamne/flowershopbe-v4/common/db"
)

func GenerateOTP(max int) (string, error) {
	var table = [...]byte{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0'}
	b := make([]byte, max)
	n, err := io.ReadAtLeast(rand.Reader, b, max)
	if n != max || err != nil {
		return "123456", err
	}
	for i := 0; i < len(table); i++ {
		b[i] = table[int(b[i])%len(table)]
	}
	return string(b), nil
}

func SaveOTP(ctx context.Context, email string, otp string) error {
	err := db.Rdb.Set(ctx, email, otp, 5*time.Minute).Err()
	return err
}

func VerifyOTP(ctx context.Context, email string, inputOTP string) (bool, error) {
	storedOTP, err := db.Rdb.Get(ctx, email).Result()
	if err != nil {
		return false, errors.New("OTP is expired / not found")
	}

	return storedOTP == inputOTP, nil
}
