package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tristaamne/flowershopbe-v4/common/payment"
)

func (h *OrderHandler) PayOSWebHook() gin.HandlerFunc {
	return func(c *gin.Context) {
		var body payment.PayOSWebhookBody
		if err := c.ShouldBind(&body); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error when receive body from PayOS": "invalid body"})
			return
		}

		err := h.payment.CheckWebhookSignature(c.Request.Context(), body)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error when check signature": err,
			})
			return
		}
		err = h.service.UpdateOrderStatus(c.Request.Context(), body.Data.OrderCode, "paid")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error when update order status": err,
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"status": "payment succeeded",
		})
		return
	}
}
