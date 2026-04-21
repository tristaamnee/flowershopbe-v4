package main

import (
	"context"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/tristaamne/flowershopbe-v4/common/config"
	"github.com/tristaamne/flowershopbe-v4/common/db"
	"github.com/tristaamne/flowershopbe-v4/common/payment"
	"github.com/tristaamne/flowershopbe-v4/common/ratelimit"
	orderRoute "github.com/tristaamne/flowershopbe-v4/orders/route"
	productRoute "github.com/tristaamne/flowershopbe-v4/products/route"
	userRoute "github.com/tristaamne/flowershopbe-v4/users/route"
)

func main() {

	cfg := config.LoadConfig()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	db.InitRedis()

	r := gin.Default()

	r.Use(ratelimit.RateLimiter(10, 20))

	client, err := db.ConnectClient(ctx, cfg.MongoDBURI)
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}
	defer db.CloseDatabase(client)

	database := db.ConnectToDatabase(*client)

	payOsProvider := payment.NewPayOSProvider(cfg)

	//production route
	productRoute.ConfigureRoute(r, database)
	userRoute.ConfigureRoute(r, database)
	orderRoute.ConfigureOrderRoute(r, database, payOsProvider, cfg)

	err = r.Run(":8080")
	if err != nil {
		return
	}
}
