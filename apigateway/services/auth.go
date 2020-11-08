package services

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/s3f4/go-load/apigateway/models"
	"github.com/s3f4/go-load/apigateway/repository"
)

// AuthService service
type AuthService interface {
	CreateAuthCache(userID uint, at *models.AccessToken, rt *models.RefreshToken) error
	GetAuthCache(UUID string) (uint, error)
	DeleteAuthCache(atUUID, rtUUID string) error
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
func (s *authService) GetAuthCache(UUID string) (uint, error) {
	userid, err := s.r.Get(UUID)
	if err != nil {
		return 0, err
	}
	userID, _ := strconv.ParseUint(userid, 10, 64)
	return uint(userID), nil
}

// DeleteAuthCache clears auth objects on cache database.
func (s *authService) DeleteAuthCache(atUUID, rtUUID string) error {

	// If this call comes from refresh token
	// there may not be an atUUID at cache server
	if len(atUUID) > 0 {
		deletedAt, err := s.r.Del(atUUID)
		if err != nil {
			return err
		}

		if deletedAt != 1 {
			fmt.Printf("deleted access token: %#v", deletedAt)
			return errors.New("something went wrong")
		}
	}

	deletedRt, err := s.r.Del(rtUUID)
	if err != nil {
		return err
	}

	fmt.Printf("deleted refresh token: %#v", deletedRt)
	return nil
}
