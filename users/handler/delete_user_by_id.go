package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tristaamne/flowershopbe-v4/users/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DeleteUserById(coll *mongo.Collection) gin.HandlerFunc {
	return func(c *gin.Context) {
		idStr := c.Param("id")
		id, err := primitive.ObjectIDFromHex(idStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "invalid user id",
			})
			return
		}
		filter := bson.M{"_id": id}
		opts := options.Delete()

		err = repository.DeleteAUser(coll, filter, opts)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"Error when delete user: ": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"Message": "User deleted successfully",
			"ID":      id.Hex(),
		})
	}
}
