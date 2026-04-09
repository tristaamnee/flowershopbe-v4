package route

import (
	"github.com/gin-gonic/gin"
	"github.com/tristaamne/flowershopbe-v4/common/middleware"
	"github.com/tristaamne/flowershopbe-v4/products/handler"
	"go.mongodb.org/mongo-driver/mongo"
)

func ConfigureRoute(r *gin.Engine, db *mongo.Database) {
	coll := db.Collection("products")
	//admin
	r.POST("/products", middleware.APIAuthentication(3), handler.CreateNewProduct(coll))
	r.PATCH("/products/:id", middleware.APIAuthentication(3), handler.UpdateProduct(coll))
	r.DELETE("/products/:id", middleware.APIAuthentication(3), handler.DeleteProductByID(coll))
	r.GET("/products/:id", middleware.APIAuthentication(3), handler.GetProductByID(coll))
	//public
	r.GET("/products", handler.GetProductByCategory(coll))
	r.GET("/products/:id", handler.GetProductByID(coll))
}
