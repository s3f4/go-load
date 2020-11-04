package services

import (
	"errors"
	"strconv"
	"time"

	"github.com/s3f4/go-load/apigateway/models"
	"github.com/s3f4/go-load/apigateway/repository"
)

// AuthService service
type AuthService interface {
	CreateAuthCache(userID uint, tokenDetails *models.TokenDetails) error
	FetchAuth(authDetails *models.AccessDetails) (uint, error)
}

type authService struct {
	r repository.RedisRepository
}

var authServiceObject *authService

// NewAuthService creates new AuthService object
func NewAuthService() AuthService {
	if authServiceObject == nil {
		return &authService{
			r: repository.NewRedisRepository(),
		}
	}
	return authServiceObject
}

// CreateAuthCache creates auth object on cache database.
func (s *authService) CreateAuthCache(userID uint, tokenDetails *models.TokenDetails) error {
	at := time.Unix(tokenDetails.AccessTokenExpires, 0)
	rt := time.Unix(tokenDetails.RefreshTokenExpires, 0)

	now := time.Now()
	userIDstr := strconv.Itoa(int(userID))

	if err := s.r.Set(tokenDetails.AccessUUID, userIDstr, at.Sub(now)); err != nil {
		return err
	}

	if err := s.r.Set(tokenDetails.RefreshUUID, userIDstr, rt.Sub(now)); err != nil {
		return err
	}

	return nil
}

// FetchAuth gets auth object from cache
func (s *authService) FetchAuth(authDetails *models.AccessDetails) (uint, error) {
	userid, err := s.r.Get(authDetails.AccessUUID)
	if err != nil {
		return 0, err
	}
	userID, _ := strconv.ParseUint(userid, 10, 64)
	return uint(userID), nil
}

// Clear clears auth objects on cache database.
func (s *authService) Clear(authDetails *models.TokenDetails) error {
	deletedAt, err := s.r.Del(authDetails.AccessUUID)
	if err != nil {
		return err
	}

	deletedRt, err := s.r.Del(authDetails.RefreshUUID)
	if err != nil {
		return err
	}

	if deletedAt != 1 || deletedRt != 1 {
		return errors.New("something went wrong")
	}

	return nil
}
