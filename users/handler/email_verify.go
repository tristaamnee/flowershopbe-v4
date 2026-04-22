package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *UserHandler) EmailVerify() gin.HandlerFunc {
	return func(c *gin.Context) {
		var VerifyEmailReq struct {
			Email string `form:"email" binding:"required"`
			OTP   string `form:"otp" binding:"required"`
		}

		if err := c.ShouldBindQuery(&VerifyEmailReq); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Invalid data: ": err.Error()})
			return
		}

		err := h.service.EmailVerify(c.Request.Context(), VerifyEmailReq.Email, VerifyEmailReq.OTP)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"Error when trying to send email verify mail": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"Success": "Email verify success"})
	}
}
