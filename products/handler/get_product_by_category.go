package handler

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/tristaamne/flowershopbe-v4/products/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetProductByCategory(coll *mongo.Collection) gin.HandlerFunc {
	return func(c *gin.Context) {
		filter := bson.M{}
		opts := options.Find()

		category := c.Query("category")
		pageStr := c.DefaultQuery("page", "1")
		limitStr := c.DefaultQuery("limit", "10")
		sortField := c.DefaultQuery("sortField", "date_ordered")
		orderStr := c.DefaultQuery("order", "-1")

		page, _ := strconv.ParseInt(pageStr, 10, 64)
		limit, _ := strconv.ParseInt(limitStr, 10, 64)
		order, _ := strconv.Atoi(orderStr)

		skip := (page - 1) * limit

		opts.SetLimit(limit)
		opts.SetSkip(skip)
		opts.SetSort(bson.D{{sortField, order}})

		if category != "" {
			filter["category"] = bson.M{
				"$regex":   category,
				"$options": "i",
			}
		}

		productData, err := repository.GetByCategory(coll, filter, opts)
		if err != nil {
			c.JSON(400, gin.H{
				"Error": err,
			})
			return
		}
		c.JSON(200, gin.H{
			"Product": productData,
		})
	}
}
