package services

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/s3f4/go-load/apigateway/models"
	"github.com/stretchr/testify/assert"
)

func Test_Token_CreateToken(t *testing.T) {
	as := NewTokenService()
	user := &models.User{Email: "email@email.com", Password: "123456"}

	req := httptest.NewRequest(http.MethodPost, "/auth/signup", strings.NewReader(`{"email":"email@email.com","password":"123456"}`))
	w := httptest.NewRecorder()
	as.CreateToken(req, user)

	res := w.Result()

	assert.Equal(t, http.StatusOK, res.StatusCode, "%d status is not equal to %d", http.StatusOK, res.StatusCode)
}
