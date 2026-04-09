package route

import (
	"github.com/gin-gonic/gin"
	"github.com/tristaamne/flowershopbe-v4/common/middleware"
	"github.com/tristaamne/flowershopbe-v4/users/handler"
	"go.mongodb.org/mongo-driver/mongo"
)

func ConfigureRoute(r *gin.Engine, db *mongo.Database) {
	coll := db.Collection("users")
	//public
	r.POST("/register", handler.Register(coll))
	r.POST("/verify-otp", handler.EmailVerify(coll))
	r.POST("/login", middleware.LoginAuthenticate(coll), handler.Login())

	//user
	r.PATCH("/user/:id", middleware.APIAuthentication(0), middleware.CheckOwner(), handler.UpdateUser(coll))

	//admin
	r.GET("/user", middleware.APIAuthentication(3), handler.GetUsersByCondition(coll))
	r.DELETE("/user/:id", middleware.APIAuthentication(3), handler.DeleteUserById(coll))
}
