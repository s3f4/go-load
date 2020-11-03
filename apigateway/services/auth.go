package services

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-redis/redis/v8"
	"github.com/s3f4/go-load/apigateway/models"
	"github.com/twinj/uuid"
)

// AuthService service
type AuthService interface {
	CreateToken(userID uint) (*models.TokenDetails, error)
	CreateAuthCache(userID uint, tokenDetails *models.TokenDetails) error
	ExtractToken(r *http.Request) string
	VerifyToken(r *http.Request) (*jwt.Token, error)
	IsTokenValid(r *http.Request) error
	ExtractTokenMetadata(r *http.Request) (*models.AccessDetails, error)
	FetchAuth(authDetails *models.AccessDetails) (uint, error)
	DeleteAuth(givenUUID string) (int64, error)
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

func (s *authService) CreateToken(userID uint) (*models.TokenDetails, error) {
	td := &models.TokenDetails{}

	td.AccessTokenExpires = time.Now().Add(time.Minute * 15).Unix()
	td.AccessUUID = uuid.NewV4().String()

	td.RefreshTokenExpires = time.Now().Add(time.Hour * 24 * 7).Unix()
	td.RefreshUUID = uuid.NewV4().String()

	var err error

	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["refresh_uuid"] = td.AccessUUID
	atClaims["user_id"] = userID
	atClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)

	td.AccessToken, err = at.SignedString([]byte(os.Getenv("AUTH_ACCESS_SECRET")))
	if err != nil {
		return nil, err
	}

	rtClaims := jwt.MapClaims{}
	rtClaims["refresh_uuid"] = td.RefreshUUID
	rtClaims["user_id"] = userID
	rtClaims["exp"] = td.RefreshTokenExpires

	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)

	td.RefreshToken, err = rt.SignedString([]byte(os.Getenv("AUTH_REFRESH_SECRET")))

	if err != nil {
		return nil, err
	}

	return td, nil
}

func (s *authService) CreateAuthCache(userID uint, tokenDetails *models.TokenDetails) error {
	at := time.Unix(tokenDetails.AccessTokenExpires, 0)
	rt := time.Unix(tokenDetails.RefreshTokenExpires, 0)

	now := time.Now()

	ctx := context.Background()
	errAccess := s.client.Set(ctx, tokenDetails.AccessUUID, strconv.Itoa(int(userID)), at.Sub(now)).Err()
	if errAccess != nil {
		return errAccess
	}

	errRefresh := s.client.Set(ctx, tokenDetails.RefreshUUID, strconv.Itoa(int(userID)), rt.Sub(now)).Err()
	if errRefresh != nil {
		return errRefresh
	}
	return nil
}

func (s *authService) ExtractToken(r *http.Request) string {
	bearToken := r.Header.Get("Authorization")

	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}

	return ""
}

func (s *authService) VerifyToken(r *http.Request) (*jwt.Token, error) {
	tokenString := s.ExtractToken(r)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Make sure that the token method conform to "SigningMethodMAC"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method %v", token.Header["alg"])
		}
		return []byte(os.Getenv("AUTH_ACCESS_SECRET")), nil
	})

	if err != nil {
		return nil, err
	}

	return token, nil
}

func (s *authService) IsTokenValid(r *http.Request) error {
	token, err := s.VerifyToken(r)
	if err != nil {
		return err
	}
	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		return err
	}
	return nil
}

func (s *authService) ExtractTokenMetadata(r *http.Request) (*models.AccessDetails, error) {
	token, err := s.VerifyToken(r)
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		accessUUID, ok := claims["access_uuid"].(string)
		if !ok {
			return nil, err
		}

		userID, err := strconv.ParseUint(fmt.Sprintf("%.f", claims["user_id"]), 10, 64)
		if err != nil {
			return nil, err
		}
		return &models.AccessDetails{
			AccessUUID: accessUUID,
			UserID:     uint(userID),
		}, nil
	}
	return nil, err
}

func (s *authService) FetchAuth(authDetails *models.AccessDetails) (uint, error) {
	userid, err := s.client.Get(context.Background(), authDetails.AccessUUID).Result()
	if err != nil {
		return 0, err
	}
	userID, _ := strconv.ParseUint(userid, 10, 64)
	return uint(userID), nil
}

func (s *authService) DeleteAuth(givenUUID string) (int64, error) {
	deleted, err := s.client.Del(context.Background(), givenUUID).Result()
	if err != nil {
		return 0, err
	}
	return deleted, nil
}
