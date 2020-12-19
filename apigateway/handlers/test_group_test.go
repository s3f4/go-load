package handlers

import (
	"fmt"
	"net/http"
	"strings"
	"testing"

	"github.com/s3f4/go-load/apigateway/library"
	"github.com/s3f4/go-load/apigateway/mocks"
	"github.com/s3f4/go-load/apigateway/models"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestGroup_Create(t *testing.T) {
	repository := new(mocks.TestGroupRepository)
	repository.On("Create", &models.TestGroup{Name: "test"}).Return(nil)
	testGroupHandler := NewTestGroupHandler(repository)

	res, body := makeRequest("/test_group", http.MethodPost, testGroupHandler.Create, strings.NewReader(`{"name":"test", "url":"url"}`))
	assert.Equal(t, `{"status":true,"data":{"id":0,"name":"test","tests":null}}`, string(body))
	assert.Equal(t, res.StatusCode, http.StatusOK, "%d status is not equal to %d", res.StatusCode, http.StatusOK)
}

func TestGroup_Create_ParseError(t *testing.T) {
	repository := new(mocks.TestGroupRepository)
	testGroupHandler := NewTestGroupHandler(repository)

	res, body := makeRequest("/test_group", http.MethodPost, testGroupHandler.Create, strings.NewReader(`{"name":"test", "":"url"`))
	assert.Equal(t, fmt.Sprintf(`{"status":false,"message":"%s"}`, library.ErrBadRequest), string(body))
	assert.Equal(t, http.StatusBadRequest, res.StatusCode, "%d status is not equal to %d", http.StatusBadRequest, res.StatusCode)
}

func TestGroup_Create_DBError(t *testing.T) {
	repository := new(mocks.TestGroupRepository)
	repository.On("Create", &models.TestGroup{Name: "test"}).Return(gorm.ErrNotImplemented)

	testGroupHandler := NewTestGroupHandler(repository)
	res, body := makeRequest("/test_group", http.MethodPost, testGroupHandler.Create, strings.NewReader(`{"name":"test"}`))

	assert.Equal(t, fmt.Sprintf(`{"status":false,"message":"%s"}`, library.ErrInternalServerError), string(body))
	assert.Equal(t, http.StatusInternalServerError, res.StatusCode, "%d status is not equal to %d", http.StatusInternalServerError, res.StatusCode)
}
