package route

import (
	"github.com/gin-gonic/gin"
	"github.com/tristaamne/flowershopbe-v4/orders/handler"
	"go.mongodb.org/mongo-driver/mongo"
)

func ConfigureRoute(r *gin.Engine, db *mongo.Database) {
	//coll := db.Collection("orders")

	//guest
	r.POST("/orders/checkout", handler.CheckOut())
}
