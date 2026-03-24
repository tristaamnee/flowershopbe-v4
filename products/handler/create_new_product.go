package handler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/tristaamne/flowershopbe-v4/products/model"
	"github.com/tristaamne/flowershopbe-v4/products/repository"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreateNewProduct(coll *mongo.Collection) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req model.CreateProductRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		product := &model.Product{
			Name:        req.Name,
			Pictures:    req.Pictures,
			Description: req.Description,
			Price:       req.Price,
			Detail:      req.Detail,
			Categories:  req.Categories,
			UpdatedAt:   time.Now(),
			CreatedAt:   time.Now(),
		}

		id, err := repository.CreateAProduct(coll, product)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, gin.H{"id": id.Hex()})
	}
}
