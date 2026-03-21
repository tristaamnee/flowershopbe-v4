package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/tristaamne/flowershopbe-v4/products/model"
	"github.com/tristaamne/flowershopbe-v4/products/repository"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreateNewProduct(coll *mongo.Collection) gin.HandlerFunc {
	return func(c *gin.Context) {
		data := model.Product{}
		data.Name = c.Query("name")
		data.Categories = c.QueryArray("categories")
		data.Detail = c.Query("detail")
		data.Description = c.Query("description")
		data.Pictures = c.QueryArray("pictures")
		//
		priceStr := c.Query("price")
		price, _ := strconv.ParseInt(priceStr, 10, 64)
		data.Price = price
		//

		result, err := repository.CreateAProduct(coll, data)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusCreated, gin.H{
			"product": result,
		})
	}
}
