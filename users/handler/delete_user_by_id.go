package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (h *UserHandler) DeleteUserById() gin.HandlerFunc {
	return func(c *gin.Context) {
		idStr := c.Param("id")
		id, err := primitive.ObjectIDFromHex(idStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "invalid user id",
			})
			return
		}
		err = h.service.DeleteProductByID(c.Request.Context(), id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error when process delete product in db": err.Error(),
			})
		}

		c.JSON(http.StatusOK, gin.H{
			"Message": "User deleted successfully",
			"ID":      id.Hex(),
		})
	}
}
