package repositories

import (
	"context"

	. "github.com/aungpao2537/bn-login-repositories/domain/datasources"

	"github.com/go-redis/redis/v8"
)

type redisConnectionRepository struct {
	Context context.Context
	Redis   *redis.Client
}

type IRedisConnectionRepository interface {
}

func NewRedisRepository(redis *RedisConnection) IRedisConnectionRepository {
	return &redisConnectionRepository{
		Context: redis.Context,
		Redis:   redis.Redis,
	}
}
