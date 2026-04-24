package middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/tristaamne/flowershopbe-v4/common/config"
	"github.com/tristaamne/flowershopbe-v4/common/pagination"
	"github.com/tristaamne/flowershopbe-v4/users/model"
	"github.com/tristaamne/flowershopbe-v4/users/service"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

func LoginAuthenticate(svc service.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		email, password, hasAuth := c.Request.BasicAuth()
		if !hasAuth {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			c.Abort()
			return
		}

		filter := bson.M{model.ColEmail: email}
		pg := pagination.PaginationQuery{}

		user, err := svc.GetUserByCondition(c.Request.Context(), filter, pg)

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

func APIAuthentication(cfg *config.Config, roleRequire int64) gin.HandlerFunc {
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
			return []byte(cfg.JWTSecret), nil
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

		rawRole, ok := claims["role"].(float64)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Can't parse token"})
			c.Abort()
			return
		}
		userRole := int64(rawRole)
		// 0 = normal user, 1 = favor guest, 2 = admin, 3 = master
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
