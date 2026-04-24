package route

import (
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"github.com/tristaamne/flowershopbe-v4/common/config"
	"github.com/tristaamne/flowershopbe-v4/common/middleware"
	"github.com/tristaamne/flowershopbe-v4/common/payment"
	"github.com/tristaamne/flowershopbe-v4/orders/handler"
	"github.com/tristaamne/flowershopbe-v4/orders/repository"
	"github.com/tristaamne/flowershopbe-v4/orders/service"
	prodRepository "github.com/tristaamne/flowershopbe-v4/products/repository"
	prodService "github.com/tristaamne/flowershopbe-v4/products/service"
	"go.mongodb.org/mongo-driver/mongo"
)

func ConfigureOrderRoute(r *gin.Engine, db *mongo.Database, payment payment.PaymentProvider, cfg *config.Config, rdb *redis.Client) {
	coll := db.Collection("orders")

	prodColl := db.Collection("products")
	prodRepo := prodRepository.NewProductRepository(prodColl)
	prodSvc := prodService.NewService(prodRepo, cfg)

	orderRepo := repository.NewMongoOrderRepository(coll)
	orderSvc := service.NewService(orderRepo, prodSvc, payment, cfg, rdb)
	orderHandler := handler.NewOrderHandler(orderSvc, payment)

	orderGroup := r.Group("/orders")
	{
		//checkout for members
		orderGroup.POST("/checkout/member", middleware.APIAuthentication(cfg, 0), orderHandler.MemberCheckOut())
		//checkout for guests
		orderGroup.POST("/checkout/guest", orderHandler.GuestCheckout())
	}

	//payment
	r.POST("/payment/payos-webhook", orderHandler.PayOSWebHook())

}
