package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/s3f4/go-load/apigateway/library"
	"github.com/s3f4/go-load/apigateway/mocks"
	"github.com/s3f4/go-load/apigateway/models"
	"github.com/s3f4/go-load/apigateway/services"
	"github.com/stretchr/testify/assert"
)

func Test_Auth_Signup(t *testing.T) {
	userRepository := new(mocks.UserRepository)
	authService := new(mocks.AuthService)
	tokenService := new(mocks.TokenService)
	user := &models.User{Email: "email@email.com", Password: "123456"}
	userStr, _ := json.Marshal(user)

	userRepository.On("Create", user).Return(nil)
	authService.On("CreateAuthCache", &models.AccessToken{}, &models.RefreshToken{}).Return(nil)

	authHandler := NewAuthHandler(userRepository, authService, tokenService)

	req := httptest.NewRequest(http.MethodPost, "/auth/signup", strings.NewReader(`{"email":"email@email.com","password":"123456"}`))
	tokenService.On("CreateToken", req, user).Return(&models.AccessToken{}, &models.RefreshToken{}, nil)

	w := httptest.NewRecorder()
	authHandler.Signup(w, req)
	res := w.Result()
	body, _ := ioutil.ReadAll(res.Body)

	assert.Equal(t, fmt.Sprintf(`{"status":true,"data":{"token":"","user":%s}}`, userStr), string(body))
	assert.Equal(t, http.StatusOK, res.StatusCode, "%d status is not equal to %d", http.StatusOK, res.StatusCode)
}
func Test_Auth_Signup_ParseError(t *testing.T) {
	userRepository := new(mocks.UserRepository)
	authService := new(mocks.AuthService)
	tokenService := new(mocks.TokenService)
	user := &models.User{Email: "email@email.com", Password: "123456"}

	userRepository.On("Create", user).Return(nil)
	authService.On("CreateAuthCache", &models.AccessToken{}, &models.RefreshToken{}).Return(nil)

	authHandler := NewAuthHandler(userRepository, authService, tokenService)

	req := httptest.NewRequest(http.MethodPost, "/auth/signup", strings.NewReader(`{"email":"email@email.com","password":"123456"`))
	tokenService.On("CreateToken", req, user).Return(&models.AccessToken{}, &models.RefreshToken{}, nil)

	w := httptest.NewRecorder()
	authHandler.Signup(w, req)
	res := w.Result()
	body, _ := ioutil.ReadAll(res.Body)

	assert.Equal(t, fmt.Sprintf(`{"status":false,"message":"%s"}`, library.ErrBadRequest), string(body))
	assert.Equal(t, http.StatusBadRequest, res.StatusCode, "%d status is not equal to %d", http.StatusBadRequest, res.StatusCode)
}

func Test_Auth_Signup_DBError(t *testing.T) {
	userRepository := new(mocks.UserRepository)
	authService := new(mocks.AuthService)
	tokenService := new(mocks.TokenService)

	user := &models.User{Email: "email@email.com", Password: "123456"}
	userRepository.On("Create", user).Return(errors.New(""))
	authService.On("CreateAuthCache", &models.AccessToken{}, &models.RefreshToken{}).Return(nil)

	authHandler := NewAuthHandler(userRepository, authService, tokenService)

	req := httptest.NewRequest(http.MethodPost, "/auth/signup", strings.NewReader(`{"email":"email@email.com","password":"123456"}`))
	tokenService.On("CreateToken", req, user).Return(&models.AccessToken{}, &models.RefreshToken{}, nil)

	w := httptest.NewRecorder()
	authHandler.Signup(w, req)
	res := w.Result()
	body, _ := ioutil.ReadAll(res.Body)

	assert.Equal(t, fmt.Sprintf(`{"status":false,"message":"%s"}`, library.ErrInternalServerError), string(body))
	assert.Equal(t, http.StatusInternalServerError, res.StatusCode, "%d status is not equal to %d", http.StatusInternalServerError, res.StatusCode)
}

func Test_Auth_Signup_CreateTokenError(t *testing.T) {
	userRepository := new(mocks.UserRepository)
	authService := new(mocks.AuthService)
	tokenService := new(mocks.TokenService)

	user := &models.User{Email: "email@email.com", Password: "123456"}
	userRepository.On("Create", user).Return(nil)
	authService.On("CreateAuthCache", &models.AccessToken{}, &models.RefreshToken{}).Return(nil)

	authHandler := NewAuthHandler(userRepository, authService, tokenService)

	req := httptest.NewRequest(http.MethodPost, "/auth/signup", strings.NewReader(`{"email":"email@email.com","password":"123456"}`))
	tokenService.On("CreateToken", req, user).Return(nil, nil, errors.New(""))

	w := httptest.NewRecorder()
	authHandler.Signup(w, req)
	res := w.Result()
	body, _ := ioutil.ReadAll(res.Body)

	assert.Equal(t, fmt.Sprintf(`{"status":false,"message":"%s"}`, library.ErrUnauthorized), string(body))
	assert.Equal(t, http.StatusUnauthorized, res.StatusCode, "%d status is not equal to %d", http.StatusUnauthorized, res.StatusCode)
}

func Test_Auth_Signup_CreateAuthCacheError(t *testing.T) {
	userRepository := new(mocks.UserRepository)
	authService := new(mocks.AuthService)
	tokenService := new(mocks.TokenService)

	user := &models.User{Email: "email@email.com", Password: "123456"}
	userRepository.On("Create", user).Return(nil)
	authService.On("CreateAuthCache", &models.AccessToken{}, &models.RefreshToken{}).Return(errors.New(""))

	authHandler := NewAuthHandler(userRepository, authService, tokenService)

	req := httptest.NewRequest(http.MethodPost, "/auth/signup", strings.NewReader(`{"email":"email@email.com","password":"123456"}`))
	tokenService.On("CreateToken", req, user).Return(&models.AccessToken{}, &models.RefreshToken{}, nil)

	w := httptest.NewRecorder()
	authHandler.Signup(w, req)
	res := w.Result()
	body, _ := ioutil.ReadAll(res.Body)

	assert.Equal(t, fmt.Sprintf(`{"status":false,"message":"%s"}`, library.ErrUnauthorized), string(body))
	assert.Equal(t, http.StatusUnauthorized, res.StatusCode, "%d status is not equal to %d", http.StatusUnauthorized, res.StatusCode)
}

func Test_Auth_Signin(t *testing.T) {
	userRepository := new(mocks.UserRepository)
	authService := new(mocks.AuthService)
	tokenService := new(mocks.TokenService)
	user := &models.User{Email: "email@email.com", Password: "123456"}
	userStr, _ := json.Marshal(user)

	userRepository.On("GetByEmailAndPassword", user).Return(user, nil)
	authService.On("CreateAuthCache", &models.AccessToken{}, &models.RefreshToken{}).Return(nil)

	authHandler := NewAuthHandler(userRepository, authService, tokenService)

	req := httptest.NewRequest(http.MethodPost, "/auth/signin", strings.NewReader(`{"email":"email@email.com","password":"123456"}`))
	tokenService.On("CreateToken", req, user).Return(&models.AccessToken{}, &models.RefreshToken{}, nil)

	w := httptest.NewRecorder()
	authHandler.Signin(w, req)
	res := w.Result()
	body, _ := ioutil.ReadAll(res.Body)

	assert.Equal(t, fmt.Sprintf(`{"status":true,"data":{"token":"","user":%s}}`, userStr), string(body))
	assert.Equal(t, http.StatusOK, res.StatusCode, "%d status is not equal to %d", http.StatusOK, res.StatusCode)
}

func Test_Auth_Signin_ParseError(t *testing.T) {
	userRepository := new(mocks.UserRepository)
	authService := new(mocks.AuthService)
	tokenService := new(mocks.TokenService)
	user := &models.User{Email: "email@email.com", Password: "123456"}

	userRepository.On("GetByEmailAndPassword", user).Return(user, nil)
	authService.On("CreateAuthCache", &models.AccessToken{}, &models.RefreshToken{}).Return(nil)

	authHandler := NewAuthHandler(userRepository, authService, tokenService)

	req := httptest.NewRequest(http.MethodPost, "/auth/signin", strings.NewReader(`{"email":"email@email.com","password":"123456"`))
	tokenService.On("CreateToken", req, user).Return(&models.AccessToken{}, &models.RefreshToken{}, nil)

	w := httptest.NewRecorder()
	authHandler.Signin(w, req)
	res := w.Result()
	body, _ := ioutil.ReadAll(res.Body)

	assert.Equal(t, fmt.Sprintf(`{"status":false,"message":"%s"}`, library.ErrBadRequest), string(body))
	assert.Equal(t, http.StatusBadRequest, res.StatusCode, "%d status is not equal to %d", http.StatusBadRequest, res.StatusCode)
}

func Test_Auth_Signin_GetByEmailAndPasswordError(t *testing.T) {
	userRepository := new(mocks.UserRepository)
	authService := new(mocks.AuthService)
	tokenService := new(mocks.TokenService)
	user := &models.User{Email: "email@email.com", Password: "123456"}

	userRepository.On("GetByEmailAndPassword", user).Return(nil, errors.New(""))
	authService.On("CreateAuthCache", &models.AccessToken{}, &models.RefreshToken{}).Return(nil)

	authHandler := NewAuthHandler(userRepository, authService, tokenService)

	req := httptest.NewRequest(http.MethodPost, "/auth/signin", strings.NewReader(`{"email":"email@email.com","password":"123456"}`))
	tokenService.On("CreateToken", req, user).Return(&models.AccessToken{}, &models.RefreshToken{}, nil)

	w := httptest.NewRecorder()
	authHandler.Signin(w, req)
	res := w.Result()
	body, _ := ioutil.ReadAll(res.Body)

	assert.Equal(t, fmt.Sprintf(`{"status":false,"message":"%s"}`, library.ErrNotFound), string(body))
	assert.Equal(t, http.StatusNotFound, res.StatusCode, "%d status is not equal to %d", http.StatusNotFound, res.StatusCode)
}

func Test_Auth_Signin_CreateTokenError(t *testing.T) {
	userRepository := new(mocks.UserRepository)
	authService := new(mocks.AuthService)
	tokenService := new(mocks.TokenService)
	user := &models.User{Email: "email@email.com", Password: "123456"}

	userRepository.On("GetByEmailAndPassword", user).Return(user, nil)
	authService.On("CreateAuthCache", &models.AccessToken{}, &models.RefreshToken{}).Return(nil)

	authHandler := NewAuthHandler(userRepository, authService, tokenService)

	req := httptest.NewRequest(http.MethodPost, "/auth/signin", strings.NewReader(`{"email":"email@email.com","password":"123456"}`))
	tokenService.On("CreateToken", req, user).Return(nil, nil, errors.New(""))

	w := httptest.NewRecorder()
	authHandler.Signin(w, req)
	res := w.Result()
	body, _ := ioutil.ReadAll(res.Body)

	assert.Equal(t, fmt.Sprintf(`{"status":false,"message":"%s"}`, library.ErrUnauthorized), string(body))
	assert.Equal(t, http.StatusUnauthorized, res.StatusCode, "%d status is not equal to %d", http.StatusUnauthorized, res.StatusCode)
}

func Test_Auth_Signin_CreateAuthCacheError(t *testing.T) {
	userRepository := new(mocks.UserRepository)
	authService := new(mocks.AuthService)
	tokenService := new(mocks.TokenService)
	user := &models.User{Email: "email@email.com", Password: "123456"}

	userRepository.On("GetByEmailAndPassword", user).Return(user, nil)
	authService.On("CreateAuthCache", &models.AccessToken{}, &models.RefreshToken{}).Return(errors.New(""))

	authHandler := NewAuthHandler(userRepository, authService, tokenService)

	req := httptest.NewRequest(http.MethodPost, "/auth/signin", strings.NewReader(`{"email":"email@email.com","password":"123456"}`))
	tokenService.On("CreateToken", req, user).Return(&models.AccessToken{}, &models.RefreshToken{}, nil)

	w := httptest.NewRecorder()
	authHandler.Signin(w, req)
	res := w.Result()
	body, _ := ioutil.ReadAll(res.Body)

	assert.Equal(t, fmt.Sprintf(`{"status":false,"message":"%s"}`, library.ErrUnauthorized), string(body))
	assert.Equal(t, http.StatusUnauthorized, res.StatusCode, "%d status is not equal to %d", http.StatusUnauthorized, res.StatusCode)
}

func Test_Auth_Signout(t *testing.T) {
	userRepository := new(mocks.UserRepository)
	authService := new(mocks.AuthService)
	tokenService := new(mocks.TokenService)

	authService.On("DeleteAuthCache", "").Return(nil)
	authHandler := NewAuthHandler(userRepository, authService, tokenService)
	req := httptest.NewRequest(http.MethodPost, "/auth/signout", strings.NewReader(`{"email":"email@email.com","password":"123456"}`))
	tokenService.On("GetDetailsFromToken", req, "rt").Return(&models.Details{}, nil)

	w := httptest.NewRecorder()
	authHandler.Signout(w, req)
	res := w.Result()
	body, _ := ioutil.ReadAll(res.Body)

	assert.Equal(t, `{"status":true,"data":"Successfully logged out"}`, string(body))
	assert.Equal(t, http.StatusOK, res.StatusCode, "%d status is not equal to %d", http.StatusOK, res.StatusCode)
}

func Test_Auth_Signout_RefreshTokenError(t *testing.T) {
	userRepository := new(mocks.UserRepository)
	authService := new(mocks.AuthService)
	tokenService := new(mocks.TokenService)

	authService.On("DeleteAuthCache", "").Return(nil)
	authHandler := NewAuthHandler(userRepository, authService, tokenService)
	req := httptest.NewRequest(http.MethodPost, "/auth/signout", strings.NewReader(`{"email":"email@email.com","password":"123456"}`))
	tokenService.On("GetDetailsFromToken", req, "rt").Return(nil, errors.New(""))

	w := httptest.NewRecorder()
	authHandler.Signout(w, req)
	res := w.Result()
	body, _ := ioutil.ReadAll(res.Body)

	assert.Equal(t, fmt.Sprintf(`{"status":false,"message":"%s"}`, library.ErrUnauthorized), string(body))
	assert.Equal(t, http.StatusUnauthorized, res.StatusCode, "%d status is not equal to %d", http.StatusUnauthorized, res.StatusCode)
}

func Test_Auth_Signout_DeleteAuthCacheError(t *testing.T) {
	userRepository := new(mocks.UserRepository)
	authService := new(mocks.AuthService)
	tokenService := new(mocks.TokenService)

	authService.On("DeleteAuthCache", "").Return(errors.New(""))
	authHandler := NewAuthHandler(userRepository, authService, tokenService)
	req := httptest.NewRequest(http.MethodPost, "/auth/signout", strings.NewReader(`{"email":"email@email.com","password":"123456"}`))
	tokenService.On("GetDetailsFromToken", req, "rt").Return(&models.Details{}, nil)

	w := httptest.NewRecorder()
	authHandler.Signout(w, req)
	res := w.Result()
	body, _ := ioutil.ReadAll(res.Body)

	assert.Equal(t, fmt.Sprintf(`{"status":false,"message":"%s"}`, library.ErrUnauthorized), string(body))
	assert.Equal(t, http.StatusUnauthorized, res.StatusCode, "%d status is not equal to %d", http.StatusUnauthorized, res.StatusCode)
}

func Test_Auth_RefreshToken(t *testing.T) {
	realTokenService := services.NewTokenService()
	os.Setenv("AT_EXPIRE_MINUTES", "10")
	os.Setenv("RT_EXPIRE_MINUTES", "10")
	os.Setenv("AUTH_REFRESH_SECRET", "absdfs")

	user := &models.User{ID: 1, Email: "test@test.com", Password: "123456"}
	req := httptest.NewRequest(http.MethodPost, "/auth/_rt", strings.NewReader(`{"email":"email@email.com","password":"123456"}`))
	_, rt, _ := realTokenService.CreateToken(req, user)

	w := httptest.NewRecorder()

	userRepository := new(mocks.UserRepository)
	authService := new(mocks.AuthService)
	tokenService := new(mocks.TokenService)

	tokenService.On("TokenFromCookie", req, "rt").Return(rt.Token)

	authService.On("DeleteAuthCache", rt.UUID).Return(nil)
	authHandler := NewAuthHandler(userRepository, authService, tokenService)

	tokenService.On("GetDetailsFromToken", req, "rt").Return(&models.Details{}, nil)
	at := &models.AccessToken{}
	rt = &models.RefreshToken{}
	tokenService.On("CreateToken", req, &models.User{ID: user.ID}).Return(at, rt, nil)
	authService.On("CreateAuthCache", at, rt).Return(nil)
	userRepository.On("Get", user.ID).Return(user, nil)
	authHandler.RefreshToken(w, req)
	res := w.Result()
	body, _ := ioutil.ReadAll(res.Body)
	userStr, _ := json.Marshal(user)
	assert.Equal(t, fmt.Sprintf(`{"status":true,"data":{"token":"","user":%s}}`, userStr), string(body))
	assert.Equal(t, http.StatusOK, res.StatusCode, "%d status is not equal to %d", http.StatusOK, res.StatusCode)
}

func Test_Auth_RefreshToken_TokenParseError(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, "/auth/_rt", strings.NewReader(`{"email":"email@email.com","password":"123456"}`))

	w := httptest.NewRecorder()

	userRepository := new(mocks.UserRepository)
	authService := new(mocks.AuthService)
	tokenService := new(mocks.TokenService)

	tokenService.On("TokenFromCookie", req, "rt").Return("")
	authService.On("DeleteAuthCache", "abc").Return(nil)
	authHandler := NewAuthHandler(userRepository, authService, tokenService)

	authHandler.RefreshToken(w, req)
	res := w.Result()
	body, _ := ioutil.ReadAll(res.Body)

	assert.Equal(t, fmt.Sprintf(`{"status":false,"data":"%s"}`, "Refresh token expired"), string(body))
	assert.Equal(t, http.StatusUnauthorized, res.StatusCode, "%d status is not equal to %d", http.StatusUnauthorized, res.StatusCode)
}

