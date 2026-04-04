package handler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/tristaamne/flowershopbe-v4/products/model"
	"github.com/tristaamne/flowershopbe-v4/products/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func UpdateProduct(coll *mongo.Collection) gin.HandlerFunc {
	return func(c *gin.Context) {
		idStr := c.Param("id")
		id, err := primitive.ObjectIDFromHex(idStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Invalid product id": err.Error()})
			return
		}

		var req model.CreateProductRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		update := bson.M{}

		if req.Name != nil {
			update["name"] = *req.Name
		}
		if req.Price != nil {
			update["price"] = *req.Price
		}
		if req.Description != nil {
			update["description"] = *req.Description
		}
		if req.Detail != nil {
			update["detail"] = *req.Detail
		}
		if req.Categories != nil {
			update["categories"] = *req.Categories
		}
		if req.Pictures != nil {
			update["pictures"] = *req.Pictures
		}
		update["updated_at"] = time.Now()

		if len(update) == 0 {
			c.JSON(http.StatusNoContent, gin.H{
				"message": "no field to update",
			})
			return
		}
		filter := bson.M{"_id": id}
		er := repository.UpdateAProduct(coll, filter, update)
		if er != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error when update product: ": er.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": id.Hex()})
	}
}
