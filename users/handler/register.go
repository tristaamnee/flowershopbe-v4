package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tristaamne/flowershopbe-v4/common/utils"
	"github.com/tristaamne/flowershopbe-v4/users/model"
	"github.com/tristaamne/flowershopbe-v4/users/repository"
	"go.mongodb.org/mongo-driver/mongo"
)

func Register(coll *mongo.Collection) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req model.UserRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
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
			PhoneNumber:       req.PhoneNumber,
			Role:              0,
		}
		id, err := repository.RegisterUser(coll, user)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error when adding user into database": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"user has been added with id: ": id.Hex()})
	}
}
