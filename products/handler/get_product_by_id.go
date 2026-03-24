package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tristaamne/flowershopbe-v4/products/model"
	"github.com/tristaamne/flowershopbe-v4/products/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetProductByID(coll *mongo.Collection) gin.HandlerFunc {
	return func(c *gin.Context) {
		filter := bson.M{}
		opts := options.Find().SetLimit(1)
		idStr := c.Param("id")
		id, err := primitive.ObjectIDFromHex(idStr)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"msg": err.Error(),
			})
			return
		}

		filter["_id"] = id

		productData, err := repository.GetByCondition[model.Product](coll, filter, opts)
		if err != nil {
			c.JSON(500, gin.H{
				"Error": err,
			})
			return
		}
		if len(productData) == 0 {
			c.JSON(http.StatusNotFound, gin.H{
				"msg": "Product not found",
			})
			return
		}
		c.JSON(200, gin.H{
			"data": productData,
		})
	}
}
