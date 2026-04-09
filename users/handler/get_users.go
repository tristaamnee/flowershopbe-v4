package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tristaamne/flowershopbe-v4/users/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetUsersByCondition(coll *mongo.Collection) gin.HandlerFunc {
	return func(c *gin.Context) {
		filter := bson.M{}
		opts := options.Find()

		userData, err := repository.GetUserByCondition(coll, filter, opts)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"msg": err.Error(),
			})
			return
		}
		c.JSON(200, gin.H{
			"msg":  "success",
			"data": userData,
		})
	}
}
