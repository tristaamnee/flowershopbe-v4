package handler

import (
	"github.com/tristaamne/flowershopbe-v4/common/security/jwt"
	userModel "github.com/tristaamne/flowershopbe-v4/users/model"
)

func CreateToken(u *userModel.User) (string, error) {
	token, err := jwt.GenerateToken(u)
	if err != nil {
		return "", err
	}

	return token, nil
}
