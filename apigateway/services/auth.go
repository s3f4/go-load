package services

import (
	"errors"
	"time"

	"github.com/s3f4/go-load/apigateway/library/log"
	"github.com/s3f4/go-load/apigateway/models"
	"github.com/s3f4/go-load/apigateway/repository"
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

// NewAuthService creates new AuthService object
func NewAuthService(repository repository.RedisRepository) AuthService {
	return &authService{
		r: repository,
	}
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
	if err != nil {
		log.Debug(err)
		return err
	}

	if deletedRt != 1 {
		log.Debugf("deleted access token: %#v", deletedRt)
		return errors.New("something went wrong")
	}

	return nil
}
