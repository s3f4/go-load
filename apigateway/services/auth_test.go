package services

import (
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/s3f4/go-load/apigateway/mocks"
	"github.com/s3f4/go-load/apigateway/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_Auth_NewAuthService(t *testing.T) {
	r := new(mocks.RedisRepository)
	as := NewAuthService(r)
	assert.NotNil(t, as)
}

func Test_Auth_CreateAuthCache(t *testing.T) {
	r := new(mocks.RedisRepository)
	rt := &models.RefreshToken{UUID: "abc", Expire: time.Now().Add(time.Hour).Unix()}
	at := &models.AccessToken{Token: "abc"}
	fmt.Println(&r)
	r.On("Set", rt.UUID, at.Token, mock.Anything).Return(nil)
	as := NewAuthService(r)
	err := as.CreateAuthCache(at, rt)
	fmt.Println(err)
	assert.Nil(t, err)
}

func Test_Auth_CreateAuthCache_Error(t *testing.T) {
	r := new(mocks.RedisRepository)
	rt := &models.RefreshToken{UUID: "abc", Expire: time.Now().Add(time.Hour).Unix()}
	at := &models.AccessToken{Token: "abc"}
	r.On("Set", rt.UUID, at.Token, mock.Anything).Return(errors.New(""))

	as := NewAuthService(r)
	err := as.CreateAuthCache(at, rt)
	fmt.Println(err)
	assert.NotNil(t, err)
}

func Test_Auth_GetAuthCache(t *testing.T) {
	r := new(mocks.RedisRepository)
	r.On("Get", "").Return("token", nil)

	as := NewAuthService(r)
	s, _ := as.GetAuthCache("")
	assert.Equal(t, s, "token")
}

func Test_Auth_GetAuthCache_Error(t *testing.T) {
	r := new(mocks.RedisRepository)
	r.On("Get", "").Return("", errors.New(""))
	as := NewAuthService(r)
	str, _ := as.GetAuthCache("")
	assert.Equal(t, "", str)
}
