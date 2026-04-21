package handler

import (
	"github.com/tristaamne/flowershopbe-v4/products/service"
)

type ProductHandler struct {
	service service.Service
}

func NewProductHandler(service service.Service) *ProductHandler {
	return &ProductHandler{service: service}
}
