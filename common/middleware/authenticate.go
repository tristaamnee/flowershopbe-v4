package middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/tristaamne/flowershopbe-v4/users/model"
	"github.com/tristaamne/flowershopbe-v4/users/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
)

func LoginAuthenticate(coll *mongo.Collection) gin.HandlerFunc {
	return func(c *gin.Context) {
		email, password, hasAuth := c.Request.BasicAuth()
		if !hasAuth {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			c.Abort()
			return
		}

		filter := bson.M{model.ColEmail: email}
		opts := options.Find()

		user, err := repository.GetUserByCondition(coll, filter, opts)

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Wrong email or password"})
			c.Abort()
			return
		}

		if bcrypt.CompareHashAndPassword([]byte(user[0].Password), []byte(password)) != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			c.Abort()
			return
		}
		c.Set("user_obj", &user[0])
		c.Next()
	}
}

func APIAuthentication(roleRequire int64) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			c.Abort()
			return
		}
		token, er := jwt.Parse(authHeader, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("wrong signing method: %v", token.Header["alg"])
			}
			return []byte("secret"), nil
		})
		if er != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token invalid"})
			c.Abort()
			return
		}
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Can't parse token"})
			c.Abort()
			return
		}

		userRole := claims["role"].(int64)
		if userRole < roleRequire {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Not enough permission"})
			c.Abort()
			return
		}
		c.Set("userID", claims["id"].(string))
		c.Set("userRole", claims["role"].(int64))
		c.Next()
	}
}
