package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tristaamne/flowershopbe-v4/common/pagination"
)

func (h *UserHandler) GetUsersByCondition() gin.HandlerFunc {
	return func(c *gin.Context) {
		rawQuery := make(map[string]interface{})
		if err := c.ShouldBindQuery(&rawQuery); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid query parameters"})
			return
		}
		var pg pagination.PaginationQuery
		if err := c.ShouldBindQuery(&pg); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid pagination settings"})
			return
		}

		userData, err := h.service.GetUserByCondition(c.Request.Context(), rawQuery, pg)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error when get user from database": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"msg":  "success",
			"data": userData,
		})
	}
}
