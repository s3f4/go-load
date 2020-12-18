package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/jinzhu/gorm"
	"github.com/s3f4/go-load/apigateway/library"
	"github.com/s3f4/go-load/apigateway/middlewares"
	"github.com/s3f4/go-load/apigateway/mocks"
	"github.com/s3f4/go-load/apigateway/models"
	"github.com/stretchr/testify/assert"
)

func Test_TestCreate(t *testing.T) {
	service := new(mocks.TestService)
	repository := new(mocks.TestRepository)
	repository.On("Create", &models.Test{Name: "test", URL: "url"}).Return(nil)
	testHandler := NewTestHandler(service, repository)

	res, body := makeRequest("/test", http.MethodPost, testHandler.Create, strings.NewReader(`{"name":"test", "url":"url"}`))
	assert.Equal(t, `{"status":true,"data":{"id":0,"name":"test","test_group_id":0,"url":"url","method":"","request_count":0,"goroutine_count":0,"expected_response_code":0,"expected_response_body":"","expected_first_byte_time":0,"expected_connection_time":0,"expected_dns_time":0,"expected_tls_time":0,"expected_total_time":0,"transport_config":{"test_id":0,"disable_keep_alives":false},"test_group":null,"run_tests":null,"headers":null}}`, string(body))
	assert.Equal(t, res.StatusCode, http.StatusOK, "%d status is not equal to %d", res.StatusCode, http.StatusOK)
}

func Test_TestCreate_ParseError(t *testing.T) {
	service := new(mocks.TestService)
	repository := new(mocks.TestRepository)
	testHandler := NewTestHandler(service, repository)

	res, body := makeRequest("/test", http.MethodPost, testHandler.Create, strings.NewReader(`{"name":"test", "":"url"`))
	assert.Equal(t, fmt.Sprintf(`{"status":false,"message":"%s"}`, library.ErrBadRequest), string(body))
	assert.Equal(t, http.StatusBadRequest, res.StatusCode, "%d status is not equal to %d", res.StatusCode, http.StatusOK)
}

func Test_TestCreate_DBError(t *testing.T) {
	service := new(mocks.TestService)
	repository := new(mocks.TestRepository)

	repository.On("Create", &models.Test{}).Return(gorm.ErrInvalidSQL)
	testHandler := NewTestHandler(service, repository)

	res, body := makeRequest("/test", http.MethodPost, testHandler.Create, strings.NewReader(`{"":"test"}`))

	assert.Equal(t, fmt.Sprintf(`{"status":false,"message":"%s"}`, library.ErrInternalServerError), string(body))
	assert.Equal(t, http.StatusInternalServerError, res.StatusCode, "%d status is not equal to %d", res.StatusCode, http.StatusOK)
}

func Test_TestDelete(t *testing.T) {
	service := new(mocks.TestService)
	repository := new(mocks.TestRepository)

	test := &models.Test{ID: 1}
	repository.On("Delete", test).Return(nil)
	testHandler := NewTestHandler(service, repository)

	req := httptest.NewRequest(http.MethodDelete, "/test/1", nil)
	ctx := context.WithValue(req.Context(), middlewares.TestCtxKey, test)
	req = req.WithContext(ctx)

	w := httptest.NewRecorder()
	testHandler.Delete(w, req)
	res := w.Result()
	body, _ := ioutil.ReadAll(res.Body)

	testStr, _ := json.Marshal(test)
	assert.Equal(t, fmt.Sprintf(`{"status":true,"data":%s}`, string(testStr)), string(body))
	assert.Equal(t, http.StatusOK, res.StatusCode, "%d status is not equal to %d", res.StatusCode, http.StatusOK)
}

func Test_TestDelete_NotFound(t *testing.T) {
	service := new(mocks.TestService)
	repository := new(mocks.TestRepository)

	repository.On("Delete", &models.Test{}).Return(nil)
	testHandler := NewTestHandler(service, repository)

	res, body := makeRequest("/test/1", http.MethodDelete, testHandler.Delete, strings.NewReader(`{"":"test"}`))

	assert.Equal(t, fmt.Sprintf(`{"status":false,"message":"%s"}`, library.ErrUnprocessableEntity), string(body))
	assert.Equal(t, http.StatusUnprocessableEntity, res.StatusCode, "%d status is not equal to %d", res.StatusCode, http.StatusOK)
}
