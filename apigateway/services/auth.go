package services

import (
	"errors"
	"fmt"
	"time"

	"github.com/s3f4/go-load/apigateway/models"
	"github.com/s3f4/go-load/apigateway/repository"
	"github.com/s3f4/mu/log"
)

// AuthService service
type AuthService interface {
	CreateAuthCache(at *models.AccessToken, rt *models.RefreshToken) error
	GetAuthCache(UUID string) (string, error)
	DeleteAuthCache(UUID string) error
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
func (s *authService) CreateAuthCache(at *models.AccessToken, rt *models.RefreshToken) error {
	rtExpire := time.Unix(rt.Expire, 0)

	now := time.Now()

	if err := s.r.Set(rt.UUID, at.Token, rtExpire.Sub(now)); err != nil {
		return err
	}

	return nil
}

// GetAuthCache gets auth object from cache
func (s *authService) GetAuthCache(UUID string) (string, error) {
	token, err := s.r.Get(UUID)
	if err != nil {
		return "", err
	}
	return token, nil
}

// DeleteAuthCache clears auth objects on cache database.
func (s *authService) DeleteAuthCache(rtUUID string) error {
	deletedRt, err := s.r.Del(rtUUID)
	log.Debug(deletedRt)
	log.Debug(err)
	if err != nil {
		return err
	}

	if deletedRt != 1 {
		fmt.Printf("deleted access token: %#v", deletedRt)
		return errors.New("something went wrong")
	}

	return nil
}
