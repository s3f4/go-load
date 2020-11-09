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
)

type authHandlerInterface interface {
	Signin(w http.ResponseWriter, r *http.Request)
	Signup(w http.ResponseWriter, r *http.Request)
	Signout(w http.ResponseWriter, r *http.Request)
	RefreshToken(w http.ResponseWriter, r *http.Request)
	CurrentUser(w http.ResponseWriter, r *http.Request)
	ResponseWithCookie(http.ResponseWriter, *http.Request, *models.AccessToken, *models.RefreshToken)
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

	if err := h.as.CreateAuthCache(user.ID, at, rt); err != nil {
		library.R401(w, r, fmt.Errorf("Unauthorized"))
		return
	}

	h.ResponseWithCookie(w, r, at, rt)
}

func (h *authHandler) Signin(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := parse(r, &user); err != nil {
		library.R400(w, r, fmt.Errorf("Bad Request"))
		return
	}

	dbUser, err := h.ur.GetByEmailAndPassword(&user)
	if err != nil {
		library.R404(w, r, "User Not Found")
		return
	}

	at, rt, err := h.ts.CreateToken(r, dbUser)
	if err != nil {
		library.R401(w, r, fmt.Errorf("Unauthorized"))
		return
	}

	if err := h.as.CreateAuthCache(dbUser.ID, at, rt); err != nil {
		library.R401(w, r, fmt.Errorf("Unauthorized"))
		return
	}

	h.ResponseWithCookie(w, r, at, rt)
}

func (h *authHandler) Signout(w http.ResponseWriter, r *http.Request) {
	access, err := h.ts.GetDetailsFromToken(r, "header")

	if err != nil {
		library.R401(w, r, fmt.Errorf("Unauthorized"))
		return
	}

	refresh, err := h.ts.GetDetailsFromToken(r, "cookie")

	if err != nil {
		library.R401(w, r, fmt.Errorf("Unauthorized"))
		return
	}

	if err = h.as.DeleteAuthCache(access.UUID, refresh.UUID); err != nil {
		library.R401(w, r, err.Error())
		return
	}

	library.R200(w, r, "Successfully logged out")
}

func (h *authHandler) RefreshToken(w http.ResponseWriter, r *http.Request) {
	// Read refresh token
	refreshToken := h.ts.TokenFromCookie(r)
	fmt.Println(refreshToken)

	//verify the token
	token, err := jwt.Parse(refreshToken, func(token *jwt.Token) (interface{}, error) {
		//Make sure that the token method conform to "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("AUTH_REFRESH_SECRET")), nil
	})

	if err != nil {
		library.R401(w, r, "Refresh token expired")
		return
	}

	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		library.R401(w, r, err)
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		refreshUUID, ok := claims["uuid"].(string)
		if !ok {
			library.R422(w, r, err)
			return
		}
		userID, err := strconv.ParseUint(fmt.Sprintf("%.f", claims["user_id"]), 10, 64)
		if err != nil {
			library.R422(w, r, err)
			return
		}

		if err := h.as.DeleteAuthCache(refreshUUID, ""); err != nil {
			fmt.Println(err)
			library.R401(w, r, fmt.Errorf("Unauthorized"))
			return
		}

		at, rt, err := h.ts.CreateToken(r, &models.User{ID: uint(userID)})
		if err != nil {
			library.R403(w, r, err.Error())
			return
		}

		if err := h.as.CreateAuthCache(uint(userID), at, rt); err != nil {
			fmt.Println(err)
			library.R403(w, r, err.Error())
			return
		}

		h.ResponseWithCookie(w, r, at, rt)
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

func (h *authHandler) ResponseWithCookie(w http.ResponseWriter, r *http.Request, at *models.AccessToken, rt *models.RefreshToken) {
	rtCookie := http.Cookie{
		HttpOnly: true,
		Name:     "rt",
		Value:    rt.Token,
		Expires:  time.Unix(rt.Expire, 0),
	}

	if os.Getenv("APP_ENV") == "production" {
		rtCookie.Domain = os.Getenv("DOMAIN")
		rtCookie.Secure = true
	}

	http.SetCookie(w, &rtCookie)

	library.R200(w, r, map[string]string{
		"token": at.Token,
	})
}
