package route

import (
	"github.com/gin-gonic/gin"
	"github.com/tristaamne/flowershopbe-v4/products/handler"
	"go.mongodb.org/mongo-driver/mongo"
)

func ConfigureRoute(r *gin.Engine, db *mongo.Database) {
	coll := db.Collection("products")
	r.POST("/products", handler.CreateNewProduct(coll))
	r.PATCH("/products/:id", handler.UpdateProduct(coll))
	r.GET("/products", handler.GetProductByCategory(coll))
	r.GET("/products/:id", handler.GetProductByID(coll))
	r.DELETE("/products/:id", handler.DeleteProductByID(coll))
}
