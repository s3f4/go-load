package services

import (
	"context"
	"errors"
	"os"
	"strconv"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/s3f4/go-load/apigateway/models"
)

// AuthService service
type AuthService interface {
	CreateAuthCache(userID uint, tokenDetails *models.TokenDetails) error
	FetchAuth(authDetails *models.AccessDetails) (uint, error)
}

type authService struct {
	client *redis.Client
}

var authServiceObject *authService
var redisClient *redis.Client

// NewAuthService creates new AuthService object
func NewAuthService() AuthService {
	if authServiceObject == nil {
		return &authService{
			client: ConnectRedis(),
		}
	}
	return authServiceObject
}

// ConnectRedis connect redis
func ConnectRedis() *redis.Client {
	dsn := os.Getenv("REDIS_DSN")
	if len(dsn) == 0 {
		dsn = "redis:6379"
	}

	client := redis.NewClient(&redis.Options{
		Addr:     dsn,
		Password: os.Getenv("REDIS_SERVER_PASSWORD"),
		DB:       0,
	})

	if _, err := client.Ping(context.Background()).Result(); err != nil {
		panic(err)
	}
	return client
}

func (s *authService) CreateAuthCache(userID uint, tokenDetails *models.TokenDetails) error {
	at := time.Unix(tokenDetails.AccessTokenExpires, 0)
	rt := time.Unix(tokenDetails.RefreshTokenExpires, 0)

	now := time.Now()
	ctx := context.Background()

	if errAccess := s.client.Set(ctx, tokenDetails.AccessUUID, strconv.Itoa(int(userID)), at.Sub(now)).Err(); errAccess != nil {
		return errAccess
	}

	if errRefresh := s.client.Set(ctx, tokenDetails.RefreshUUID, strconv.Itoa(int(userID)), rt.Sub(now)).Err(); errRefresh != nil {
		return errRefresh
	}

	return nil
}

func (s *authService) FetchAuth(authDetails *models.AccessDetails) (uint, error) {
	userid, err := s.client.Get(context.Background(), authDetails.AccessUUID).Result()
	if err != nil {
		return 0, err
	}
	userID, _ := strconv.ParseUint(userid, 10, 64)
	return uint(userID), nil
}

func (s *authService) Clear(authDetails *models.TokenDetails) error {
	ctx := context.Background()
	deletedAt, err := s.client.Del(ctx, authDetails.AccessUUID).Result()
	if err != nil {
		return err
	}

	deletedRt, err := s.client.Del(ctx, authDetails.RefreshUUID).Result()
	if err != nil {
		return err
	}

	if deletedAt != 1 || deletedRt != 1 {
		return errors.New("something went wrong")
	}

	return nil
}

// Delete redis entry with key
func (s *authService) Del(key string) (int64, error) {
	deleted, err := s.client.Del(context.Background(), key).Result()
	if err != nil {
		return 0, err
	}
	return deleted, nil
}
