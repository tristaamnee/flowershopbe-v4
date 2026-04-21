package route

import (
	"github.com/gin-gonic/gin"
	"github.com/tristaamne/flowershopbe-v4/common/config"
	"github.com/tristaamne/flowershopbe-v4/common/payment"
	"github.com/tristaamne/flowershopbe-v4/orders/handler"
	"github.com/tristaamne/flowershopbe-v4/orders/repository"
	"github.com/tristaamne/flowershopbe-v4/orders/service"
	"go.mongodb.org/mongo-driver/mongo"
)

func ConfigureOrderRoute(r *gin.Engine, db *mongo.Database, payOsProvider payment.PaymentProvider, cfg *config.Config) {

	orderRepo := repository.NewMongoOrderRepository(db.Collection("orders"))
	orderSvc := service.NewService(orderRepo, payOsProvider, cfg)
	orderHandler := handler.NewOrderHandler(orderSvc)

	orderGroup := r.Group("/orders")
	{
		orderGroup.POST("/checkout", orderHandler.CheckOut())
	}
}

//cai nay refactor xong, con 3 cai nua :))))
