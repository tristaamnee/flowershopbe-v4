package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tristaamne/flowershopbe-v4/users/model"
	"go.mongodb.org/mongo-driver/mongo"
)

func Login(coll *mongo.Collection) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req model.LoginRequest
		if err := c.ShouldBind(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error when receive login form: ": err.Error()})
			return
		}

	}
}
