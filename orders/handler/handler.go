package handler

import (
	"github.com/tristaamne/flowershopbe-v4/common/payment"
	"github.com/tristaamne/flowershopbe-v4/orders/service"
)

type OrderHandler struct {
	service service.Service
	payment payment.PaymentProvider
}

func NewOrderHandler(service service.Service, payment payment.PaymentProvider) *OrderHandler {
	return &OrderHandler{service: service, payment: payment}
}
