package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tristaamne/flowershopbe-v4/orders/model"
)

func (h *OrderHandler) MemberCheckOut() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req model.MemberOrderRequest
		if err := c.ShouldBind(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Invalid Order body": err.Error()})
			return
		}

		userId, ok := c.Get("userId")
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"UserId not found when process checkout": userId})
			return
		}
		userIdStr := userId.(string)

		result, err := h.service.MemberCheckout(c.Request.Context(), req, userIdStr)
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

func (h *OrderHandler) GuestCheckout() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req model.GuestOrderRequest
		if err := c.ShouldBind(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Invalid Order body": err.Error()})
			return
		}

		result, err := h.service.GuestCheckout(c.Request.Context(), req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"Error when checking for guest": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"status":  "success",
			"payment": result["data"],
		})
	}
}
