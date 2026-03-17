package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v10"
	"github.com/tristaamne/flowershopbe-v4/common/db"
	"github.com/tristaamne/flowershopbe-v4/common/ratelimit"
	"github.com/tristaamne/flowershopbe-v4/common/utils"
	productModel "github.com/tristaamne/flowershopbe-v4/products/model"
	productRoute "github.com/tristaamne/flowershopbe-v4/products/route"
)

func main() {
	r := gin.Default()

	conf, err := utils.Get()

	r.Use(ratelimit.RateLimiter(10, 20))

	database := db.ConnectDatabase(conf.(db.DatabaseConfiguration))
	defer database.Close()

	createDBTables(err, database)

	productRoute.ConfigureRoute(r, database)
	err = r.Run(":8080")
	if err != nil {
		return
	}
}

func createDBTables(err error, database *pg.DB) {
	er := db.CreateTable(database, (*productModel.Product)(nil))
	if er != nil {
		log.Fatalf("Error creating table: %v", er)
	}
}
