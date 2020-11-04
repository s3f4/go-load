package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/s3f4/go-load/apigateway/models"
	"github.com/s3f4/go-load/apigateway/repository"
	"github.com/s3f4/go-load/apigateway/services"
	. "github.com/s3f4/mu"
)

type authHandlerInterface interface {
	Signin(w http.ResponseWriter, r *http.Request)
	Signout(w http.ResponseWriter, r *http.Request)
	RefreshToken(w http.ResponseWriter, r *http.Request)
}

type authHandler struct {
	ur repository.UserRepository
	as services.AuthService
	ts services.TokenService
}

var (
	// AuthHandler is used for authentication
	AuthHandler authHandlerInterface = &authHandler{
		ur: repository.NewUserRepository(),
		as: services.NewAuthService(),
		ts: services.NewTokenService(),
	}
)

func (h *authHandler) Signin(w http.ResponseWriter, r *http.Request) {
	var userRequest models.User
	if err := json.NewDecoder(r.Body).Decode(&userRequest); err != nil {
		R400(w, "Bad Request")
		return
	}

	user, err := h.ur.GetByEmailAndPassword(&userRequest)
	if err != nil {
		R404(w, "User Not Found")
		return
	}

	tokenDetails, err := h.ts.CreateToken(r, user)
	if err != nil {
		R401(w, "unauthorized")
		return
	}

	if err := h.as.CreateAuthCache(user.ID, tokenDetails); err != nil {
		R401(w, "unauthorized")
		return
	}

	cookie := http.Cookie{
		HttpOnly: true,
		Name:     "rt",
		Value:    tokenDetails.RefreshToken,
		Expires:  time.Unix(tokenDetails.RefreshTokenExpires, 0),
	}

	if os.Getenv("APP_ENV") == "production" {
		cookie.Domain = os.Getenv("DOMAIN")
	}

	http.SetCookie(w, &cookie)

}

func (h *authHandler) Signout(w http.ResponseWriter, r *http.Request) {
	metadata, err := h.ts.ExtractTokenMetadata(r)
	if err != nil {
		R401(w, "unauthorized")
		return
	}
	delErr := h.as.Clear(metadata)
	if delErr != nil {
		R401(w, delErr.Error())
		return
	}
	R200(w, "Successfully logged out")
}

func (h *authHandler) RefreshToken(w http.ResponseWriter, r *http.Request) {
	// Read refresh token
	var refreshToken string

	//verify the token
	token, err := jwt.Parse(refreshToken, func(token *jwt.Token) (interface{}, error) {
		//Make sure that the token method conform to "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("AUTH_REFRESH_SECRET")), nil
	})

	if err != nil {
		R401(w, "Refresh token expired")
		return
	}

	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		R401(w, err)
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		refreshUUID, ok := claims["refresh_uuid"].(string)
		if !ok {
			R422(w, err)
			return
		}
		userID, err := strconv.ParseUint(fmt.Sprintf("%.f", claims["user_id"]), 10, 64)
		if err != nil {
			R422(w, err)
			return
		}

		deleted, delErr := h.as.Clean(refreshUUID)
		if delErr != nil || deleted == 0 {
			R401(w, "unauthorized")
			return
		}

		ts, createErr := h.as.CreateToken(uint(userID))
		if createErr != nil {
			R403(w, createErr.Error())
			return
		}

		saveErr := h.as.CreateAuthCache(uint(userID), ts)
		if saveErr != nil {
			R403(w, saveErr.Error())
			return
		}
		tokens := map[string]string{
			"access_token":  ts.AccessToken,
			"refresh_token": ts.RefreshToken,
		}

		R200(w, tokens)
	} else {
		R401(w, "refresh expired")
	}
}
