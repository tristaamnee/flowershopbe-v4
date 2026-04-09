package jwt

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/tristaamne/flowershopbe-v4/users/model"
)

func GenerateToken(user *model.User) (string, error) {

	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		return "", errors.New("JWT_SECRET environment variable is not set")
	}

	claims := jwt.MapClaims{
		"sub":  user.Email,
		"role": user.Role,
		"id":   user.ID.Hex(),
		"iat":  time.Now().Unix(),
		"exp":  time.Now().Add(time.Hour * 72).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
