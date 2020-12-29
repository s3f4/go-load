package services

import (
	"errors"
	"time"

	"github.com/s3f4/go-load/apigateway/library"
	"github.com/s3f4/go-load/apigateway/library/log"
	"github.com/s3f4/go-load/apigateway/models"
	"github.com/s3f4/go-load/apigateway/repository"
	"golang.org/x/crypto/bcrypt"
)

// AuthService service
type AuthService interface {
	CreateAuthCache(at *models.AccessToken, rt *models.RefreshToken) error
	GetAuthCache(UUID string) (string, error)
	DeleteAuthCache(UUID string) error
	CreatePassword(user *models.User) error
	HashPassword(password, salt string) (string, error)
	CheckPassword(password, salt, hash string) bool
}

type authService struct {
	r repository.CacheRepository
}

// NewAuthService creates new AuthService object
func NewAuthService(repository repository.CacheRepository) AuthService {
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

// CreatePassword creates password for user
func (s *authService) CreatePassword(user *models.User) error {
	salt, err := library.RandomBytes(32)
	if err != nil {
		log.Info(err)
		return err
	}
	user.Salt = string(salt)
	user.Password, err = s.HashPassword(user.Password, user.Salt)

	if err != nil {
		log.Info(err)
		return err
	}
	return nil
}

// HashPassword generate hash from password and random salt string
func (s *authService) HashPassword(password, salt string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password+salt), 14)
	return string(bytes), err
}

// CheckPassword check hash is equal to coming password and db users salt's hash
func (s *authService) CheckPassword(password, salt, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password+salt))
	return err == nil
}
