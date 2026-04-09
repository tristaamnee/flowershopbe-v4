package handler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/tristaamne/flowershopbe-v4/common/mailer"
	"github.com/tristaamne/flowershopbe-v4/common/utils"
	"github.com/tristaamne/flowershopbe-v4/users/model"
	"github.com/tristaamne/flowershopbe-v4/users/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func UpdateUser(coll *mongo.Collection) gin.HandlerFunc {
	return func(c *gin.Context) {
		userIDStr, exist := c.Get("userID")
		if !exist {
			c.JSON(http.StatusBadRequest, gin.H{"Invalid id": userIDStr})
			return
		}

		userID, err := primitive.ObjectIDFromHex(userIDStr.(string))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "ID không hợp lệ"})
			return
		}

		filter := bson.M{"_id": userID}

		var req model.UserRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Invalid data": req})
			return
		}
		setData := bson.M{}
		if req.Name != "" {
			setData["name"] = req.Name
		}
		if req.Password != "" {
			hashedPassword, err := utils.PasswordHasher(req.Password)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"Invalid password": err.Error()})
				return
			}
			setData["password"] = hashedPassword
		}
		if !req.Birthday.IsZero() {
			setData["birthday"] = req.Birthday
		}
		if req.Email != "" {
			if err := mailer.EmailValidate(req.Email); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Wrong email format"})
				return
			}

			if err := mailer.OTPSender(c.Request.Context(), req.Email); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"Error when sending OTP": err.Error()})
				return
			}
			setData["email"] = req.Email
			setData["email_verified"] = false
		}

		if len(req.DeliveryAddresses) > 0 {
			setData["delivery_addresses"] = req.DeliveryAddresses
		}
		if req.ProfilePicture != "" {
			setData["profile_picture"] = req.ProfilePicture
		}
		if len(setData) == 0 {
			c.JSON(http.StatusNoContent, gin.H{
				"message": "no field to update",
			})
			return
		}
		setData["updated_at"] = time.Now()

		update := bson.M{"$set": setData}
		er := repository.UpdateAUser(coll, filter, update)
		if er != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"Error when update a user: ": er,
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"user has been updated": userID.Hex(),
		})
	}
}
