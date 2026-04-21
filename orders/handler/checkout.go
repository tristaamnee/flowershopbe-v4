package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tristaamne/flowershopbe-v4/orders/model"
	"github.com/tristaamne/flowershopbe-v4/orders/repository"
	"github.com/tristaamne/flowershopbe-v4/orders/service"
)

type OrderHandler struct {
	repo    repository.OrderRepository
	service service.Service
}

func NewOrderHandler(service service.Service) *OrderHandler {
	return &OrderHandler{service: service}
}

func (h *OrderHandler) CheckOut() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req model.OrderRequest
		if err := c.ShouldBind(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Invalid Order body": err.Error()})
			return
		}

		result, err := h.service.Checkout(c.Request.Context(), req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"status":  "success",
			"payment": result["data"],
		})

	}
}
