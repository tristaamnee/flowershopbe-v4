package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tristaamne/flowershopbe-v4/common/db"
	"github.com/tristaamne/flowershopbe-v4/common/mailer"
	"github.com/tristaamne/flowershopbe-v4/users/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func EmailVerify(coll *mongo.Collection) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req struct {
			Email string `form:"email" binding:"required"`
			OTP   string `form:"otp" binding:"required"`
		}

		if err := c.ShouldBindQuery(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Invalid data: ": err.Error()})
			return
		}

		isValid, err := mailer.VerifyOTP(c.Request.Context(), req.Email, req.OTP)
		if err != nil || !isValid {
			c.JSON(http.StatusInternalServerError, gin.H{"Error": "OTP not match"})
			return
		}

		filter := bson.M{"email": req.Email}
		update := bson.M{"$set": bson.M{"email_verified": true}}

		err = repository.UpdateAUser(coll, filter, update)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"Error when update user email verify status": err.Error()})
			return
		}

		db.SessionRdb.Del(c.Request.Context(), req.Email)

		c.JSON(http.StatusOK, gin.H{"Success": "Email verify success"})
	}
}
