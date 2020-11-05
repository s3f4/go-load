package repository

import (
	"context"
	"os"
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

var redisRepositoryObject *redisRepository

// NewRedisRepository ...
func NewRedisRepository() RedisRepository {
	if redisRepositoryObject == nil {
		redisRepositoryObject = &redisRepository{
			client: ConnectRedis(),
		}
	}
	return redisRepositoryObject
}

// ConnectRedis connect redis
func ConnectRedis() *redis.Client {
	dsn := os.Getenv("REDIS_DSN")
	if len(dsn) == 0 {
		dsn = "redis:6379"
	}

	client := redis.NewClient(&redis.Options{
		Addr: dsn,
		DB:   0,
	})

	if _, err := client.Ping(context.Background()).Result(); err != nil {
		panic(err)
	}
	return client
}

func (r *redisRepository) Set(key, value string, expire time.Duration) error {
	return r.client.Set(context.Background(), key, value, expire).Err()
}

func (r *redisRepository) Get(key string) (string, error) {
	return r.client.Get(context.Background(), key).Result()
}

func (r *redisRepository) Del(key string) (int64, error) {
	deleted, err := r.client.Del(context.Background(), key).Result()
	if err != nil {
		return 0, err
	}
	return deleted, nil
}
