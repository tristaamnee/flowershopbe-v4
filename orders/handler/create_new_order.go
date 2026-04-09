package handler

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreateNewOrder(coll *mongo.Collection) gin.HandlerFunc {
	return func(c *gin.Context) {
	}
}
