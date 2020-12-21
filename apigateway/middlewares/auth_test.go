package middlewares

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/dgrijalva/jwt-go"
	"github.com/s3f4/go-load/apigateway/library"
	"github.com/s3f4/go-load/apigateway/mocks"
	"github.com/s3f4/go-load/apigateway/models"
	"github.com/stretchr/testify/assert"
)

func Test_Auth(t *testing.T) {
	ts := new(mocks.TokenService)
	as := new(mocks.AuthService)
	m := NewMiddleware(ts, as, nil, nil, nil)

	var next http.HandlerFunc
	next = func(w http.ResponseWriter, r *http.Request) {
		val, ok := r.Context().Value(UserIDCtxKey).(uint)
		if !ok {
			t.Error("User ID not found")
		}
		assert.Equal(t, uint(1), val)
	}

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	res := httptest.NewRecorder()

	token := new(jwt.Token)
	details := &models.Details{UUID: "abc", UserID: 1}
	ts.On("VerifyToken", req, "at").Return(token, nil)
	ts.On("GetDetailsFromToken", req, "at").Return(details, nil)
	as.On("GetAuthCache", details.UUID).Return("", nil)

	test := m.AuthCtx(next)
	test.ServeHTTP(res, req)
}

func Test_Auth_VerifyTokenError(t *testing.T) {
	ts := new(mocks.TokenService)
	as := new(mocks.AuthService)
	m := NewMiddleware(ts, as, nil, nil, nil)

	var next http.HandlerFunc
	next = func(w http.ResponseWriter, r *http.Request) {
	}

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()

	details := &models.Details{UUID: "abc", UserID: 1}
	ts.On("VerifyToken", req, "at").Return(nil, errors.New(""))
	ts.On("GetDetailsFromToken", req, "at").Return(details, nil)
	as.On("GetAuthCache", details.UUID).Return("", nil)

	test := m.AuthCtx(next)
	test.ServeHTTP(w, req)

	res := w.Result()
	body, _ := ioutil.ReadAll(res.Body)

	assert.Equal(t, fmt.Sprintf(`{"status":false,"message":"%s"}`, library.ErrUnauthorized), string(body))
	assert.Equal(t, http.StatusUnauthorized, res.StatusCode, "%d status is not equal to %d", http.StatusUnauthorized, res.StatusCode)
}

func Test_Auth_GetDetailsFromToken_Error(t *testing.T) {
	ts := new(mocks.TokenService)
	as := new(mocks.AuthService)
	m := NewMiddleware(ts, as, nil, nil, nil)

	var next http.HandlerFunc
	next = func(w http.ResponseWriter, r *http.Request) {
	}

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()

	token := new(jwt.Token)
	details := &models.Details{UUID: "abc", UserID: 1}
	ts.On("VerifyToken", req, "at").Return(token, nil)
	ts.On("GetDetailsFromToken", req, "at").Return(nil, errors.New(""))
	as.On("GetAuthCache", details.UUID).Return("", nil)

	test := m.AuthCtx(next)
	test.ServeHTTP(w, req)

	res := w.Result()
	body, _ := ioutil.ReadAll(res.Body)

	assert.Equal(t, fmt.Sprintf(`{"status":false,"message":"%s"}`, library.ErrUnauthorized), string(body))
	assert.Equal(t, http.StatusUnauthorized, res.StatusCode, "%d status is not equal to %d", http.StatusUnauthorized, res.StatusCode)
}

func Test_Auth_GetAuthCache_Error(t *testing.T) {
	ts := new(mocks.TokenService)
	as := new(mocks.AuthService)
	m := NewMiddleware(ts, as, nil, nil, nil)

	var next http.HandlerFunc
	next = func(w http.ResponseWriter, r *http.Request) {
	}

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()

	token := new(jwt.Token)
	details := &models.Details{UUID: "abc", UserID: 1}
	ts.On("VerifyToken", req, "at").Return(token, nil)
	ts.On("GetDetailsFromToken", req, "at").Return(details, nil)
	as.On("GetAuthCache", details.UUID).Return("", errors.New(""))

	test := m.AuthCtx(next)
	test.ServeHTTP(w, req)

	res := w.Result()
	body, _ := ioutil.ReadAll(res.Body)

	assert.Equal(t, fmt.Sprintf(`{"status":false,"message":"%s"}`, library.ErrUnauthorized), string(body))
	assert.Equal(t, http.StatusUnauthorized, res.StatusCode, "%d status is not equal to %d", http.StatusUnauthorized, res.StatusCode)
}

func Test_Auth_IP_DOMAIN_Error(t *testing.T) {
	ts := new(mocks.TokenService)
	as := new(mocks.AuthService)
	m := NewMiddleware(ts, as, nil, nil, nil)

	var next http.HandlerFunc
	next = func(w http.ResponseWriter, r *http.Request) {
	}

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()

	token := new(jwt.Token)
	details := &models.Details{UUID: "abc", UserID: 1}
	ts.On("VerifyToken", req, "at").Return(token, nil)
	ts.On("GetDetailsFromToken", req, "at").Return(details, nil)
	as.On("GetAuthCache", details.UUID).Return("", nil)

	test := m.AuthCtx(next)
	os.Setenv("APP_ENV", "production")
	os.Setenv("DOMAIN", "abc")
	test.ServeHTTP(w, req)

	res := w.Result()
	body, _ := ioutil.ReadAll(res.Body)

	assert.Equal(t, fmt.Sprintf(`{"status":false,"message":"%s"}`, library.ErrUnauthorized), string(body))
	assert.Equal(t, http.StatusUnauthorized, res.StatusCode, "%d status is not equal to %d", http.StatusUnauthorized, res.StatusCode)
}
