package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (h *ProductHandler) GetProductByID() gin.HandlerFunc {
	return func(c *gin.Context) {

		idStr := c.Param("id")
		id, err := primitive.ObjectIDFromHex(idStr)

		ids := append([]primitive.ObjectID(nil), id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"msg": err.Error(),
			})
			return
		}
		productData, err := h.service.GetProductByID(c.Request.Context(), ids)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"msg": err.Error(),
			})
			return
		}
		if len(productData) == 0 {
			c.JSON(http.StatusBadRequest, gin.H{
				"msg": "Product not found",
			})
			return
		}
		c.JSON(200, gin.H{
			"data": productData,
		})
	}
}
