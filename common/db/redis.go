package db

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/tristaamne/flowershopbe-v4/common/config"
)

var (
	CacheRdb   *redis.Client
	SessionRdb *redis.Client
	PaymentRdb *redis.Client
)

func InitRedis(cfg *config.Config) {
	addr := cfg.RedisAddr
	password := cfg.RedisPass

	CacheRdb = redis.NewClient(&redis.Options{Addr: addr, Password: password, DB: 0})
	SessionRdb = redis.NewClient(&redis.Options{Addr: addr, Password: password, DB: 0})
	PaymentRdb = redis.NewClient(&redis.Options{Addr: addr, Password: password, DB: 0})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	clients := map[string]*redis.Client{
		"Cache (DB 0)":   CacheRdb,
		"Session (DB 0)": SessionRdb,
		"Payment (DB 0)": PaymentRdb,
	}

	for name, client := range clients {
		if err := client.Ping(ctx).Err(); err != nil {
			panic(fmt.Sprintf("Lỗi kết nối Redis %s: %v", name, err))
		}
	}
	fmt.Println("Redis connected successfully!")
}
