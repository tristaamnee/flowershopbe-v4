package db

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/redis/go-redis/v9"
)

var (
	CacheRdb   *redis.Client
	SessionRdb *redis.Client
	PaymentRdb *redis.Client
)

func InitRedis() {
	addr := os.Getenv("REDIS_ADDR")
	password := os.Getenv("REDIS_PASSWORD")

	CacheRdb = redis.NewClient(&redis.Options{Addr: addr, Password: password, DB: 0})
	SessionRdb = redis.NewClient(&redis.Options{Addr: addr, Password: password, DB: 1})
	PaymentRdb = redis.NewClient(&redis.Options{Addr: addr, Password: password, DB: 2})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	clients := map[string]*redis.Client{
		"Cache (DB 0)":   CacheRdb,
		"Session (DB 1)": SessionRdb,
		"Payment (DB 2)": PaymentRdb,
	}

	for name, client := range clients {
		if err := client.Ping(ctx).Err(); err != nil {
			panic(fmt.Sprintf("Lỗi kết nối Redis %s: %v", name, err))
		}
	}
	fmt.Println("Redis connected to DB 0, 1, 2 successfully!")
}
