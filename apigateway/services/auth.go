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
	CreateAuthCache(userID uint, at *models.AccessToken, rt *models.RefreshToken) error
	GetAuthCache(authDetails *models.AccessDetails) (uint, error)
	DeleteAuthCache(at *models.AccessToken, rt *models.RefreshToken) error
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
func (s *authService) CreateAuthCache(
	userID uint,
	at *models.AccessToken,
	rt *models.RefreshToken,
) error {
	atExpire := time.Unix(at.Expire, 0)
	rtExpire := time.Unix(rt.Expire, 0)

	now := time.Now()
	userIDstr := strconv.Itoa(int(userID))

	if err := s.r.Set(at.UUID, userIDstr, atExpire.Sub(now)); err != nil {
		return err
	}

	if err := s.r.Set(rt.UUID, userIDstr, rtExpire.Sub(now)); err != nil {
		return err
	}

	return nil
}

// GetAuthCache gets auth object from cache
func (s *authService) GetAuthCache(authDetails *models.AccessDetails) (uint, error) {
	userid, err := s.r.Get(authDetails.AccessUUID)
	if err != nil {
		return 0, err
	}
	userID, _ := strconv.ParseUint(userid, 10, 64)
	return uint(userID), nil
}

// DeleteAuthCache clears auth objects on cache database.
func (s *authService) DeleteAuthCache(at *models.AccessToken, rt *models.RefreshToken) error {
	deletedAt, err := s.r.Del(at.UUID)
	if err != nil {
		return err
	}

	deletedRt, err := s.r.Del(rt.UUID)
	if err != nil {
		return err
	}

	if deletedAt != 1 || deletedRt != 1 {
		return errors.New("something went wrong")
	}

	return nil
}
