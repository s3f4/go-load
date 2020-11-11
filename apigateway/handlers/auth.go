package handlers

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/s3f4/go-load/apigateway/library"
	"github.com/s3f4/go-load/apigateway/middlewares"
	"github.com/s3f4/go-load/apigateway/models"
	"github.com/s3f4/go-load/apigateway/repository"
	"github.com/s3f4/go-load/apigateway/services"
	"github.com/s3f4/mu/log"
)

type authHandlerInterface interface {
	Signin(w http.ResponseWriter, r *http.Request)
	Signup(w http.ResponseWriter, r *http.Request)
	Signout(w http.ResponseWriter, r *http.Request)
	RefreshToken(w http.ResponseWriter, r *http.Request)
	CurrentUser(w http.ResponseWriter, r *http.Request)
	ResponseWithCookie(http.ResponseWriter, *http.Request, *models.User, *models.AccessToken, *models.RefreshToken)
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

func (h *authHandler) Signup(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := parse(r, &user); err != nil {
		library.R400(w, r, fmt.Errorf("Bad Request"))
		return
	}

	if err := h.ur.Create(&user); err != nil {
		library.R500(w, r, err)
		return
	}

	at, rt, err := h.ts.CreateToken(r, &user)
	if err != nil {
		library.R401(w, r, fmt.Errorf("Unauthorized"))
		return
	}

	if err := h.as.CreateAuthCache(at, rt); err != nil {
		library.R401(w, r, fmt.Errorf("Unauthorized"))
		return
	}

	h.ResponseWithCookie(w, r, &user, at, rt)
}

func (h *authHandler) Signin(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := parse(r, &user); err != nil {
		log.Debug(err)
		library.R400(w, r, fmt.Errorf("Bad Request"))
		return
	}

	dbUser, err := h.ur.GetByEmailAndPassword(&user)
	if err != nil {
		log.Debug(err)
		library.R404(w, r, "User Not Found")
		return
	}

	at, rt, err := h.ts.CreateToken(r, dbUser)
	if err != nil {
		log.Debug(err)
		library.R401(w, r, fmt.Errorf("Unauthorized"))
		return
	}

	if err := h.as.CreateAuthCache(at, rt); err != nil {
		log.Debug(err)
		library.R401(w, r, fmt.Errorf("Unauthorized"))
		return
	}

	h.ResponseWithCookie(w, r, dbUser, at, rt)
}

func (h *authHandler) Signout(w http.ResponseWriter, r *http.Request) {
	refresh, err := h.ts.GetDetailsFromToken(r, "rt")

	if err != nil {
		log.Debug(err)
		library.R401(w, r, fmt.Errorf("Unauthorized"))
		return
	}

	if err = h.as.DeleteAuthCache(refresh.UUID); err != nil {
		log.Debug(err)
		library.R401(w, r, err.Error())
		return
	}

	http.SetCookie(w, &http.Cookie{
		MaxAge:   -1,
		Name:     "rt",
		Value:    "",
		Path:     "/",
		HttpOnly: true,
		Expires:  time.Now().Add(-100 * time.Hour),
	})

	library.R200(w, r, "Successfully logged out")
}

func (h *authHandler) RefreshToken(w http.ResponseWriter, r *http.Request) {
	refreshToken := h.ts.TokenFromCookie(r, "rt")

	token, err := jwt.Parse(refreshToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("AUTH_REFRESH_SECRET")), nil
	})

	if err != nil {
		library.R401(w, r, "Refresh token expired")
		return
	}

	if _, ok := token.Claims.(jwt.Claims); !ok || !token.Valid {
		log.Debug(err)
		library.R401(w, r, err)
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		refreshUUID, ok := claims["uuid"].(string)
		if !ok {
			log.Debug(err)
			library.R422(w, r, err)
			return
		}
		userID, err := strconv.ParseUint(fmt.Sprintf("%.f", claims["user_id"]), 10, 64)
		if err != nil {
			log.Debug(err)
			library.R422(w, r, err)
			return
		}

		if err := h.as.DeleteAuthCache(refreshUUID); err != nil {
			log.Debug(err)
			library.R401(w, r, fmt.Errorf("Unauthorized"))
			return
		}

		at, rt, err := h.ts.CreateToken(r, &models.User{ID: uint(userID)})
		if err != nil {
			library.R401(w, r, err.Error())
			return
		}

		if err := h.as.CreateAuthCache(at, rt); err != nil {
			log.Debug(err)
			library.R401(w, r, err.Error())
			return
		}

		user, err := h.ur.Get(uint(userID))
		if err != nil {
			log.Debug(err)
			library.R401(w, r, "refresh token is expired")
			return
		}

		h.ResponseWithCookie(w, r, user, at, rt)

	} else {
		library.R401(w, r, "refresh token is expired")
	}
}

func (h *authHandler) CurrentUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userID, ok := ctx.Value(middlewares.UserIDKey).(uint)
	if !ok {
		library.R422(w, r, "unprocessable entity")
		return
	}
	library.R200(w, r, userID)
}

func (h *authHandler) ResponseWithCookie(w http.ResponseWriter, r *http.Request, user *models.User, at *models.AccessToken, rt *models.RefreshToken) {
	rtCookie := http.Cookie{
		HttpOnly: true,
		Name:     "rt",
		Path:     "/",
		Expires:  time.Unix(rt.Expire, 0),
	}

	if os.Getenv("APP_ENV") == "production" {
		rtCookie.Domain = os.Getenv("DOMAIN")
		rtCookie.Secure = true
	}

	if err := library.SetCookie(w, &rtCookie, map[string]string{"rt": rt.Token}); err != nil {
		log.Debug(err)
	}

	library.R200(w, r, map[string]interface{}{
		"token": at.Token,
		"user":  user,
	})
}
