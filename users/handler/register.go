package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tristaamne/flowershopbe-v4/common/mailer"
	"github.com/tristaamne/flowershopbe-v4/common/utils"
	"github.com/tristaamne/flowershopbe-v4/users/model"
	"github.com/tristaamne/flowershopbe-v4/users/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func Register(coll *mongo.Collection) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req model.UserRequest

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Dữ liệu không hợp lệ": err.Error()})
			return
		}

		filter := bson.M{
			"email": req.Email,
		}

		er := mailer.EmailValidate(req.Email)
		if er != nil {
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"error": "Email format is invalid",
				"msg":   er,
			})
			return
		}

		existingUser, err := repository.GetUserByCondition(coll, filter, nil)
		if err == nil && existingUser != nil {
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"message": "User already exists",
			})
			return
		}

		hashedPassword, err := utils.PasswordHasher(req.Password)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error when hashing password": err.Error()})
			return
		}

		user := &model.User{
			Name:              req.Name,
			Password:          hashedPassword,
			Birthday:          req.Birthday,
			Email:             req.Email,
			DeliveryAddresses: req.DeliveryAddresses,
			Role:              0,
			ProviderID:        "manual",
			EmailVerified:     false,
		}
		id, err := repository.RegisterUser(coll, user)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error when adding user into database": err.Error()})
			return
		}

		err = mailer.OTPSender(c.Request.Context(), user.Email)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error when sending OTP": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Registration successful, please check your email for verification.	",
			"User has been added with id: ": id.Hex()})
	}
}
