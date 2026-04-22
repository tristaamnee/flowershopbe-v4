package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *UserHandler) Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		req, exists := c.Get("user_obj")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Login context missing"})
			return
		}
		tokenString, err := h.service.Login(c.Request.Context(), req)
		if err != nil || tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Login failed", "reason": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"token": tokenString})
	}
}
