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
	VerifyToken(r *http.Request, from ...string) (*jwt.Token, error)
	IsTokenValid(r *http.Request) error
	GetDetailsFromToken(r *http.Request, from string) (*models.Details, error)
}

type tokenService struct{}

var tokenServiceObject *tokenService

// NewTokenService returns a token service object
func NewTokenService() TokenService {
	if tokenServiceObject == nil {
		tokenServiceObject = &tokenService{}
	}
	return tokenServiceObject
}

func (s *tokenService) CreateToken(r *http.Request, user *models.User) (*models.AccessToken, *models.RefreshToken, error) {
	at := &models.AccessToken{}
	rt := &models.RefreshToken{}

	at.Expire = time.Now().Add(time.Minute * 15).Unix()
	at.UUID = uuid.NewV4().String()

	rt.Expire = time.Now().Add(time.Hour * 24 * 7).Unix()
	rt.UUID = uuid.NewV4().String()

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
	return library.GetCookie(r, key)["rt"]
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
func (s *tokenService) VerifyToken(r *http.Request, from ...string) (*jwt.Token, error) {
	var tokenStr string

	if from == nil {
		tokenStr = s.TokenFromHeader(r)
	} else {
		for _, t := range from {
			switch t {
			case "cookie":
				tokenStr = s.TokenFromCookie(r, "rt")
			case "header":
				tokenStr = s.TokenFromHeader(r)
			}
		}
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

	return token, nil
}

func (s *tokenService) IsTokenValid(r *http.Request) error {
	token, err := s.VerifyToken(r)
	if err != nil {
		return err
	}

	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		return err
	}

	return nil
}

func (s *tokenService) GetDetailsFromToken(r *http.Request, from string) (*models.Details, error) {
	token, err := s.VerifyToken(r, from)
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		UUID, ok := claims["uuid"].(string)
		if !ok {
			return nil, err
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
