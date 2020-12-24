package repository

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

// CacheRepository is using for redis
type CacheRepository interface {
	Set(key, value string, expire time.Duration) error
	Get(key string) (string, error)
	Del(key string) (int64, error)
}

type cacheRepository struct {
	client *redis.Client
}

// NewRedisRepository ...
func NewRedisRepository(client *redis.Client) CacheRepository {
	return &cacheRepository{
		client: client,
	}
}

func (r *cacheRepository) Set(key, value string, expire time.Duration) error {
	return r.client.Set(context.Background(), key, value, expire).Err()
}

func (r *cacheRepository) Get(key string) (string, error) {
	return r.client.Get(context.Background(), key).Result()
}

func (r *cacheRepository) Del(key string) (int64, error) {
	return r.client.Del(context.Background(), key).Result()
}
