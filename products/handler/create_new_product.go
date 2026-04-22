package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tristaamne/flowershopbe-v4/products/model"
)

func (h *ProductHandler) CreateNewProduct() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req model.CreateProductRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Invalid Product body": err.Error()})
			return
		}

		id, err := h.service.CreateANewProduct(c.Request.Context(), req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, gin.H{"id": id.Hex()})
	}
}
