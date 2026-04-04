package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tristaamne/flowershopbe-v4/users/model"
	"github.com/tristaamne/flowershopbe-v4/users/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
)

func Authenticate(coll *mongo.Collection) gin.HandlerFunc {
	return func(c *gin.Context) {
		username, password, hasAuth := c.Request.BasicAuth()
		if !hasAuth {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			c.Abort()
			return
		}

		filter := bson.M{}
		opts := options.Find()

		filter[model.ColName] = bson.M{
			"$elemMatch": bson.M{
				"$regex": username,
			},
		}
		user, err := repository.GetUserByCondition(coll, filter, opts)
		if err != nil || bcrypt.CompareHashAndPassword([]byte(user[0].Password), []byte(password)) != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			c.Abort()
			return
		}
		c.Set("username", user[0].Name)
		c.Next()
	}
}
