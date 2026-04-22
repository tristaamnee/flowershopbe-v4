package handler

import (
	"github.com/tristaamne/flowershopbe-v4/users/service"
)

type UserHandler struct {
	service service.Service
}

func NewUserHandler(service service.Service) *UserHandler {
	return &UserHandler{service: service}
}
