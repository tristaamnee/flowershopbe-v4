package otp

import (
	"context"
	"crypto/rand"
	"errors"
	"io"
	"time"

	"github.com/redis/go-redis/v9"
)

type otp struct {
	rdb *redis.Client
}

type OTP interface {
	GenerateOTP(max int) (string, error)
	SaveOTP(ctx context.Context, email string, otp string) error
	VerifyOTP(ctx context.Context, email string, inputOTP string) (bool, error)
}

func NewOTP(rdb *redis.Client) OTP {
	return &otp{
		rdb: rdb,
	}
}

func (o *otp) GenerateOTP(max int) (string, error) {
	var table = [...]byte{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0'}
	b := make([]byte, max)
	n, err := io.ReadAtLeast(rand.Reader, b, max)
	if n != max || err != nil {
		return "", err
	}
	for i := 0; i < len(b); i++ {
		b[i] = table[int(b[i])%len(table)]
	}
	return string(b), nil
}

func (o *otp) SaveOTP(ctx context.Context, email string, otp string) error {
	err := o.rdb.Set(ctx, email, otp, 2*time.Minute).Err()
	return err
}

func (o *otp) VerifyOTP(ctx context.Context, email string, inputOTP string) (bool, error) {
	storedOTP, err := o.rdb.Get(ctx, email).Result()
	if err != nil {
		return false, errors.New("OTP is expired / not found")
	}

	return storedOTP == inputOTP, nil
}
