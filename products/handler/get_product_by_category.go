package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tristaamne/flowershopbe-v4/products/model"
)

func (h *ProductHandler) GetProductByCategory() gin.HandlerFunc {
	return func(c *gin.Context) {
		category := c.Query(model.ColCategory)
		pageStr := c.DefaultQuery("page", "1")
		limitStr := c.DefaultQuery("limit", "10")
		sortField := c.DefaultQuery("sortField", "date_ordered")
		orderStr := c.DefaultQuery("order", "-1")

		productData, err := h.service.GetProductByCategory(c.Request.Context(), category, pageStr, limitStr, sortField, orderStr)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error when get product using category": err.Error(),
			})
		}
		c.JSON(200, gin.H{
			"data":  productData,
			"page":  pageStr,
			"limit": limitStr,
		})
	}
}
