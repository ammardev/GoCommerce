package connections

import (
	"github.com/ammardev/gocommerce/internal"
	"github.com/redis/go-redis/v9"
)

var Redis *redis.Client

func NewRedisConnection() {
    Redis = redis.NewClient(&redis.Options{
        Addr:     internal.GetEnv("REDIS_HOST", "127.0.0.1") + ":" + internal.GetEnv("REDIS_PORT", "6379"),
        Password: internal.GetEnv("REDIS_PASSWORD", ""),
    })
}
