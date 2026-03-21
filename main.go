package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/tristaamne/flowershopbe-v4/common/db"
	"github.com/tristaamne/flowershopbe-v4/common/ratelimit"
	productRoute "github.com/tristaamne/flowershopbe-v4/products/route"
)

func main() {
	var uri string
	if uri = os.Getenv("MONGODB_URI"); uri == "" {
		log.Fatal("MONGODB_URI environment variable not set")
	}

	r := gin.Default()

	r.Use(ratelimit.RateLimiter(10, 20))

	client, err := db.ConnectClient(uri)
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}
	defer db.CloseDatabase(client)

	database := db.ConnectToDatabase(*client)

	//production route
	productRoute.ConfigureRoute(r, database)
	err = r.Run(":8080")
	if err != nil {
		return
	}
}
