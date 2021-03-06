package services

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/s3f4/go-load/apigateway/library"
	"github.com/s3f4/go-load/apigateway/models"
	"github.com/twinj/uuid"
)

// TokenService using for jwt token methods.
type TokenService interface {
	CreateToken(r *http.Request, user *models.User) (*models.AccessToken, *models.RefreshToken, error)
	TokenFromCookie(r *http.Request, key string) string
	TokenFromHeader(r *http.Request) string
	VerifyToken(r *http.Request, key string) (*jwt.Token, error)
	GetDetailsFromToken(r *http.Request, from string) (*models.Details, error)
}

type tokenService struct{}

// NewTokenService returns a token service object
func NewTokenService() TokenService {
	return &tokenService{}
}

func (s *tokenService) CreateToken(r *http.Request, user *models.User) (*models.AccessToken, *models.RefreshToken, error) {
	at := &models.AccessToken{}
	rt := &models.RefreshToken{}

	atExpireMinutes, _ := strconv.Atoi(os.Getenv("AT_EXPIRE_MINUTES"))
	rtExpireMinutes, _ := strconv.Atoi(os.Getenv("RT_EXPIRE_MINUTES"))

	at.Expire = time.Now().Add(time.Minute * time.Duration(atExpireMinutes)).Unix()
	at.UUID = uuid.NewV4().String()

	rt.Expire = time.Now().Add(time.Minute * time.Duration(rtExpireMinutes)).Unix()
	rt.UUID = at.UUID

	var err error

	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["uuid"] = at.UUID
	atClaims["user_id"] = user.ID
	atClaims["remote_addr"] = r.RemoteAddr
	atClaims["user_agent"] = r.Header.Get("User-Agent")
	atClaims["exp"] = at.Expire

	atJWT := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	if at.Token, err = atJWT.SignedString([]byte(os.Getenv("AUTH_ACCESS_SECRET"))); err != nil {
		return nil, nil, err
	}

	rtClaims := jwt.MapClaims{}
	rtClaims["uuid"] = rt.UUID
	rtClaims["user_id"] = user.ID
	rtClaims["exp"] = rt.Expire

	rtJWT := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)
	if rt.Token, err = rtJWT.SignedString([]byte(os.Getenv("AUTH_REFRESH_SECRET"))); err != nil {
		return nil, nil, err
	}

	return at, rt, nil
}

// TokenFromCookie ...
func (s *tokenService) TokenFromCookie(r *http.Request, key string) string {
	vals, err := library.GetCookie(r, key)
	if err != nil {
		return ""
	}

	return vals[key]
}

// TokenFromHeader ...
func (s *tokenService) TokenFromHeader(r *http.Request) string {
	bearer := r.Header.Get("Authorization")
	if len(bearer) > 7 && strings.ToUpper(bearer[0:6]) == "BEARER" {
		return bearer[7:]
	}
	return ""
}

// VerifyToken
func (s *tokenService) VerifyToken(r *http.Request, key string) (*jwt.Token, error) {
	var tokenStr string
	if key == "at" {
		tokenStr = s.TokenFromHeader(r)
	} else {
		tokenStr = s.TokenFromCookie(r, key)
	}

	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method %v", token.Header["alg"])
		}
		return []byte(os.Getenv("AUTH_ACCESS_SECRET")), nil
	})

	if err != nil {
		return nil, err
	}

	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		return nil, err
	}

	return token, nil
}

func (s *tokenService) GetDetailsFromToken(r *http.Request, key string) (*models.Details, error) {
	token, err := s.VerifyToken(r, key)
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		UUID, ok := claims["uuid"].(string)
		if !ok {
			return nil, err
		}

		if key == "at" {
			remoteAddr, ok := claims["remote_addr"].(string)
			if !ok {
				return nil, err
			}

			userAgent, ok := claims["user_agent"].(string)
			if !ok {
				return nil, err
			}

			if userID, err := strconv.ParseUint(fmt.Sprintf("%.f", claims["user_id"]), 10, 64); err == nil {
				return &models.Details{
					UUID:       UUID,
					UserID:     uint(userID),
					RemoteAddr: remoteAddr,
					UserAgent:  userAgent,
				}, nil
			}
		}

		if userID, err := strconv.ParseUint(fmt.Sprintf("%.f", claims["user_id"]), 10, 64); err == nil {
			return &models.Details{
				UUID:   UUID,
				UserID: uint(userID),
			}, nil
		}
	}
	return nil, err
}
