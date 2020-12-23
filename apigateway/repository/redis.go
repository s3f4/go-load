package repository

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

// RedisRepository is using for redis
type RedisRepository interface {
	Set(key, value string, expire time.Duration) error
	Get(key string) (string, error)
	Del(key string) (int64, error)
}

type redisRepository struct {
	client *redis.Client
}

// NewRedisRepository ...
func NewRedisRepository(client *redis.Client) RedisRepository {
	return &redisRepository{
		client: client,
	}
}

func (r *redisRepository) Set(key, value string, expire time.Duration) error {
	return r.client.Set(context.Background(), key, value, expire).Err()
}

func (r *redisRepository) Get(key string) (string, error) {
	return r.client.Get(context.Background(), key).Result()
}

func (r *redisRepository) Del(key string) (int64, error) {
	return r.client.Del(context.Background(), key).Result()
}
