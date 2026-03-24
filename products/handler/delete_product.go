package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tristaamne/flowershopbe-v4/products/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DeleteProductByID(coll *mongo.Collection) gin.HandlerFunc {
	return func(c *gin.Context) {
		idStr := c.Param("id")
		id, err := primitive.ObjectIDFromHex(idStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid product id"})
		}
		filter := bson.M{"_id": id}
		opts := options.Delete()

		err = repository.DeleteAProduct(coll, filter, opts)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"Error when delete product: ": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"Message": "product deleted successfully",
			"ID":      id.Hex(),
		})
	}
}
