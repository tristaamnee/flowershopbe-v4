package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CheckOwner() gin.HandlerFunc {
	return func(c *gin.Context) {
		userIDFromToken, _ := c.Get("userID")
		userRole, _ := c.Get("userRole")
		reqIDStr := c.Param("id")

		if userIDFromToken.(string) == reqIDStr || userRole.(int) == 3 {
			c.Next()
			return
		}
		c.JSON(http.StatusForbidden, gin.H{
			"error": "You are not allowed to access this resource",
		})
		c.Abort()
		return
	}
}
