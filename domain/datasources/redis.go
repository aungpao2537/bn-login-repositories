package datasources

import (
	"context"
	"os"

	"github.com/go-redis/redis/v8"
)

type RedisConnection struct {
	Context context.Context
	Redis   *redis.Client
}

func NewRedisConnection() *RedisConnection {

	opt, _ := redis.ParseURL(os.Getenv("REDIS_URL"))
	rdb := redis.NewClient(opt)

	return &RedisConnection{
		Context: context.Background(),
		Redis:   rdb,
	}
}
