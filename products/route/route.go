package route

import (
	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v10"
	"github.com/tristaamne/flowershopbe-v4/products/handler"
)

func ConfigureRoute(r *gin.Engine, db *pg.DB) {
	r.GET("/products/:category", handler.GetProductByCategory(db))
}
