package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tristaamne/flowershopbe-v4/products/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (h *ProductHandler) UpdateProduct() gin.HandlerFunc {
	return func(c *gin.Context) {
		idStr := c.Param("id")
		id, err := primitive.ObjectIDFromHex(idStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Invalid product id": err.Error()})
			return
		}

		var req model.CreateProductRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		id, er := h.service.UpdateAProduct(c.Request.Context(), req, id)
		if er != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"id": id,
				"error when update product: ": er.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"product has been updated": id.Hex()})
	}
}
