package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tristaamne/flowershopbe-v4/users/model"
)

func (h *UserHandler) UpdateUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		userIDStr, exist := c.Get("userID")
		if !exist {
			c.JSON(http.StatusBadRequest, gin.H{"Invalid id": userIDStr})
			return
		}
		var req model.UserRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Invalid data": req})
			return
		}
		userID, err := h.service.UpdateUserById(c.Request.Context(), &req, userIDStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Error when update user": err})
		}
		c.JSON(http.StatusOK, gin.H{
			"user has been updated": userID.Hex(),
		})
	}
}
