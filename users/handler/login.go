package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tristaamne/flowershopbe-v4/token/handler"
	"github.com/tristaamne/flowershopbe-v4/users/model"
)

func Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		req, exists := c.Get("user_obj")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Login context missing"})
			return
		}

		user, ok := req.([]model.User)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Login context missing"})
		}

		tokenString, err := handler.CreateToken(&user[0])
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error when create user token": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"token": tokenString})
	}
}
