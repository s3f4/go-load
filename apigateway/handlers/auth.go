package handlers

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/s3f4/go-load/apigateway/library"
	"github.com/s3f4/go-load/apigateway/library/log"
	res "github.com/s3f4/go-load/apigateway/library/response"
	"github.com/s3f4/go-load/apigateway/middlewares"
	"github.com/s3f4/go-load/apigateway/models"
	"github.com/s3f4/go-load/apigateway/repository"
	"github.com/s3f4/go-load/apigateway/services"
)

// AuthHandler interface
type AuthHandler interface {
	Signin(w http.ResponseWriter, r *http.Request)
	Signup(w http.ResponseWriter, r *http.Request)
	Signout(w http.ResponseWriter, r *http.Request)
	RefreshToken(w http.ResponseWriter, r *http.Request)
	CurrentUser(w http.ResponseWriter, r *http.Request)
	ResponseWithCookie(
		http.ResponseWriter,
		*http.Request,
		*models.User,
		*models.AccessToken,
		*models.RefreshToken,
	)
}

type authHandler struct {
	ur repository.UserRepository
	sr repository.SettingsRepository
	as services.AuthService
	ts services.TokenService
}

// NewAuthHandler creates handler
func NewAuthHandler(
	ur repository.UserRepository,
	sr repository.SettingsRepository,
	as services.AuthService,
	ts services.TokenService,
) AuthHandler {
	return &authHandler{
		ur: ur,
		sr: sr,
		as: as,
		ts: ts,
	}
}

func (h *authHandler) Signup(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := parse(r, &user); err != nil {
		log.Debug(err)
		res.R400(w, r, library.ErrBadRequest)
		return
	}

	if h.forbidden() {
		res.R403(w, r, library.ErrForbidden)
		return
	}

	if err := h.as.CreatePassword(&user); err != nil {
		log.Debug(err)
		res.R500(w, r, library.ErrInternalServerError)
		return
	}

	if err := h.ur.Create(&user); err != nil {
		log.Debug(err)
		res.R500(w, r, library.ErrInternalServerError)
		return
	}

	at, rt, err := h.ts.CreateToken(r, &user)
	if err != nil {
		log.Debug(err)
		res.R401(w, r, library.ErrUnauthorized)
		return
	}

	if err := h.as.CreateAuthCache(at, rt); err != nil {
		log.Debug(err)
		res.R401(w, r, library.ErrUnauthorized)
		return
	}

	h.ResponseWithCookie(w, r, &user, at, rt)
}

func (h *authHandler) Signin(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := parse(r, &user); err != nil {
		log.Debug(err)
		res.R400(w, r, library.ErrBadRequest)
		return
	}

	dbUser, err := h.ur.GetByEmail(&user)
	if err != nil {
		log.Debug(err)
		res.R404(w, r, library.ErrNotFound)
		return
	}

	if !h.as.CheckPassword(user.Password, dbUser.Salt, dbUser.Password) {
		res.R401(w, r, library.ErrUnauthorized)
		return
	}

	at, rt, err := h.ts.CreateToken(r, dbUser)
	if err != nil {
		log.Debug(err)
		res.R401(w, r, library.ErrUnauthorized)
		return
	}

	if err := h.as.CreateAuthCache(at, rt); err != nil {
		log.Debug(err)
		res.R401(w, r, library.ErrUnauthorized)
		return
	}

	h.ResponseWithCookie(w, r, dbUser, at, rt)
}

func (h *authHandler) Signout(w http.ResponseWriter, r *http.Request) {
	refresh, err := h.ts.GetDetailsFromToken(r, "rt")

	if err != nil {
		log.Debug(err)
		res.R401(w, r, library.ErrUnauthorized)
		return
	}

	if err = h.as.DeleteAuthCache(refresh.UUID); err != nil {
		log.Debug(err)
		res.R401(w, r, library.ErrUnauthorized)
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

	res.R200(w, r, "Successfully logged out")
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
		res.R401(w, r, "Refresh token expired")
		return
	}

	if _, ok := token.Claims.(jwt.Claims); !ok || !token.Valid {
		log.Debug(err)
		res.R401(w, r, err)
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		refreshUUID, ok := claims["uuid"].(string)
		if !ok {
			log.Debug(err)
			res.R422(w, r, err)
			return
		}
		userID, err := strconv.ParseUint(fmt.Sprintf("%.f", claims["user_id"]), 10, 64)
		if err != nil {
			log.Debug(err)
			res.R422(w, r, err)
			return
		}

		if err := h.as.DeleteAuthCache(refreshUUID); err != nil {
			log.Debug(err)
			res.R401(w, r, library.ErrUnauthorized)
			return
		}

		at, rt, err := h.ts.CreateToken(r, &models.User{ID: uint(userID)})
		if err != nil {
			res.R401(w, r, err.Error())
			return
		}

		if err := h.as.CreateAuthCache(at, rt); err != nil {
			log.Debug(err)
			res.R401(w, r, err.Error())
			return
		}

		user, err := h.ur.Get(uint(userID))
		if err != nil {
			log.Debug(err)
			res.R401(w, r, library.ErrRefreshTokenExpire)
			return
		}

		h.ResponseWithCookie(w, r, user, at, rt)

	} else {
		res.R401(w, r, library.ErrRefreshTokenExpire)
	}
}

func (h *authHandler) CurrentUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userID, ok := ctx.Value(middlewares.UserIDCtxKey).(uint)
	if !ok {
		res.R422(w, r, library.ErrUnprocessableEntity)
		return
	}
	res.R200(w, r, userID)
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

	// Don't write password
	user.Password = "-"

	res.R200(w, r, map[string]interface{}{
		"token": at.Token,
		"user":  user,
	})
}

func (h *authHandler) forbidden() bool {
	settings, _ := h.sr.Get(models.SIGNUP)
	if settings == nil {
		settings = &models.Settings{
			Setting: string(models.SIGNUP),
			Value:   "Forbidden",
		}

		if err := h.sr.Create(settings); err != nil {
			return true
		}
	} else if settings.Setting == string(models.SIGNUP) &&
		settings.Value == "Forbidden" {
		return true
	}
	return false
}
