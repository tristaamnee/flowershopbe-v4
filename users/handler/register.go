package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tristaamne/flowershopbe-v4/users/model"
	"go.mongodb.org/mongo-driver/mongo"
)

func (h *UserHandler) Register(coll *mongo.Collection) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req model.UserRequest

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Dữ liệu không hợp lệ": err.Error()})
			return
		}
		id, err := h.service.Register(c.Request.Context(), req)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"Error when register": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Registration successful, please check your email for verification.	",
			"User has been added with id: ": id.Hex()})
	}
}
