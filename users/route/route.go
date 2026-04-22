package route

import (
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"github.com/tristaamne/flowershopbe-v4/common/config"
	"github.com/tristaamne/flowershopbe-v4/common/mailer"
	"github.com/tristaamne/flowershopbe-v4/common/middleware"
	jwt2 "github.com/tristaamne/flowershopbe-v4/common/security/jwt"
	"github.com/tristaamne/flowershopbe-v4/common/security/otp"
	"github.com/tristaamne/flowershopbe-v4/users/handler"
	"github.com/tristaamne/flowershopbe-v4/users/repository"
	"github.com/tristaamne/flowershopbe-v4/users/service"
	"go.mongodb.org/mongo-driver/mongo"
)

func ConfigureRoute(r *gin.Engine, db *mongo.Database, cfg *config.Config, rdb *redis.Client) {
	coll := db.Collection("users")

	repo := repository.NewUserRepository(coll)
	otpMaker := otp.NewOTP(rdb)
	jwt := jwt2.NewJwtSrv(cfg)
	mailSender := mailer.NewMailer(cfg, otpMaker)
	userSvc := service.NewUserService(repo, cfg, rdb, mailSender, jwt, otpMaker)
	userHandler := handler.NewUserHandler(userSvc)

	//public
	r.POST("/register", userHandler.Register(coll))
	r.POST("/verify-otp", userHandler.EmailVerify())
	r.POST("/login", middleware.LoginAuthenticate(userSvc), userHandler.Login())

	//user
	r.PATCH("/user/:id", middleware.APIAuthentication(cfg, 0), middleware.CheckOwner(), userHandler.UpdateUser())

	//admin
	r.GET("/user", middleware.APIAuthentication(cfg, 3), userHandler.GetUsersByCondition())
	r.DELETE("/user/:id", middleware.APIAuthentication(cfg, 3), userHandler.DeleteUserById())
}
