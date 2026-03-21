package route

import (
	"github.com/gin-gonic/gin"
	"github.com/tristaamne/flowershopbe-v4/products/handler"
	"go.mongodb.org/mongo-driver/mongo"
)

func ConfigureRoute(r *gin.Engine, db *mongo.Database) {
	coll := db.Collection("products")
	r.POST("/products", handler.CreateNewProduct(coll))
	r.GET("/products/:category", handler.GetProductByCategory(coll))
}
