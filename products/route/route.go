package route

import (
	"github.com/gin-gonic/gin"
	"github.com/tristaamne/flowershopbe-v4/common/config"
	"github.com/tristaamne/flowershopbe-v4/common/middleware"
	"github.com/tristaamne/flowershopbe-v4/products/handler"
	"github.com/tristaamne/flowershopbe-v4/products/repository"
	"github.com/tristaamne/flowershopbe-v4/products/service"
	"go.mongodb.org/mongo-driver/mongo"
)

func ConfigureRoute(r *gin.Engine, db *mongo.Database, cfg *config.Config) {
	coll := db.Collection("products")
	prodRepo := repository.NewProductRepository(coll)
	prodSvc := service.NewService(prodRepo, cfg)
	prodHandler := handler.NewProductHandler(prodSvc)

	//admin
	r.POST("/products", middleware.APIAuthentication(3), prodHandler.CreateNewProduct())
	r.DELETE("/products/:id", middleware.APIAuthentication(3), prodHandler.DeleteProductByID())
	r.PATCH("/products/:id", middleware.APIAuthentication(3), prodHandler.UpdateProduct())

	//public
	r.GET("/products", prodHandler.GetProductByCategory())
	r.GET("/products/:id", prodHandler.GetProductByID())
}
