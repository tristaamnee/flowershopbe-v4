package handler

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/tristaamne/flowershopbe-v4/products/model"
	"github.com/tristaamne/flowershopbe-v4/products/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetProductByCategory(coll *mongo.Collection) gin.HandlerFunc {
	return func(c *gin.Context) {
		filter := bson.M{}
		opts := options.Find()

		category := c.Query("categories")
		pageStr := c.DefaultQuery("page", "1")
		limitStr := c.DefaultQuery("limit", "10")
		sortField := c.DefaultQuery("sortField", "date_ordered")
		orderStr := c.DefaultQuery("order", "-1")

		//just being reminded that I haven't valid these thing yet :))
		page, err := strconv.Atoi(pageStr)
		if err != nil || page < 1 {
			page = 1
		}
		limit, err := strconv.Atoi(limitStr)
		if err != nil || limit < 1 || limit > 100 {
			limit = 10
		}

		order, err := strconv.Atoi(orderStr)
		if err != nil || (order != 1 && order != -1) {
			order = -1
		}

		allowedSortFields := map[string]bool{
			"create_at": true,
			"price":     true,
			"name":      true,
		}

		if !allowedSortFields[sortField] {
			sortField = "create_at"
		}

		skip := (page - 1) * limit

		opts.SetLimit(int64(limit))
		opts.SetSkip(int64(skip))
		opts.SetSort(bson.D{{sortField, order}})

		if category != "" {
			filter["categories"] = bson.M{
				"$elemMatch": bson.M{
					"$regex":   category,
					"$options": "i"},
			}
		}

		productData, err := repository.GetByCondition[model.Product](coll, filter, opts)
		if err != nil {
			c.JSON(500, gin.H{
				"Error": err,
			})
			return
		}
		c.JSON(200, gin.H{
			"data":  productData,
			"page":  page,
			"limit": limit,
		})
	}
}
