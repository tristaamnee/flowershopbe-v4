package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v10"
	"github.com/tristaamne/flowershopbe-v4/products/model"
	"github.com/tristaamne/flowershopbe-v4/products/repository"
)

func GetProductByCategory(db *pg.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		productCategory := c.Param(model.ColCategory)
		productData, err := repository.GetByCategory(db, productCategory)
		if err != nil {
			c.JSON(400, gin.H{
				"Error": "Category not found",
			})
			return
		}
		c.JSON(200, gin.H{
			"Product": productData,
		})
	}
}
