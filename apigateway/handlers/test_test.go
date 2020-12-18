package handlers

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/s3f4/go-load/apigateway/library"
	"github.com/s3f4/go-load/apigateway/middlewares"
	"github.com/s3f4/go-load/apigateway/mocks"
	"github.com/s3f4/go-load/apigateway/models"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
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
	assert.Equal(t, http.StatusBadRequest, res.StatusCode, "%d status is not equal to %d", http.StatusBadRequest, res.StatusCode)
}

func Test_TestCreate_DBError(t *testing.T) {
	service := new(mocks.TestService)
	repository := new(mocks.TestRepository)

	repository.On("Create", &models.Test{}).Return(gorm.ErrNotImplemented)
	testHandler := NewTestHandler(service, repository)

	res, body := makeRequest("/test", http.MethodPost, testHandler.Create, strings.NewReader(`{"":"test"}`))

	assert.Equal(t, fmt.Sprintf(`{"status":false,"message":"%s"}`, library.ErrInternalServerError), string(body))
	assert.Equal(t, http.StatusInternalServerError, res.StatusCode, "%d status is not equal to %d", http.StatusInternalServerError, res.StatusCode)
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
	assert.Equal(t, http.StatusOK, res.StatusCode, "%d status is not equal to %d", http.StatusOK, res.StatusCode)
}

func Test_TestDelete_NotFound(t *testing.T) {
	service := new(mocks.TestService)
	repository := new(mocks.TestRepository)

	repository.On("Delete", &models.Test{}).Return(nil)
	testHandler := NewTestHandler(service, repository)

	res, body := makeRequest("/test/1", http.MethodDelete, testHandler.Delete, strings.NewReader(`{"":"test"}`))

	assert.Equal(t, fmt.Sprintf(`{"status":false,"message":"%s"}`, library.ErrUnprocessableEntity), string(body))
	assert.Equal(t, http.StatusUnprocessableEntity, res.StatusCode, "%d status is not equal to %d", http.StatusUnprocessableEntity, res.StatusCode)
}

func Test_TestDelete_DBError(t *testing.T) {
	service := new(mocks.TestService)
	repository := new(mocks.TestRepository)

	repository.On("Delete", &models.Test{}).Return(gorm.ErrNotImplemented)
	testHandler := NewTestHandler(service, repository)
	test := &models.Test{}
	req := httptest.NewRequest(http.MethodDelete, "/test/1", nil)
	ctx := context.WithValue(req.Context(), middlewares.TestCtxKey, test)
	req = req.WithContext(ctx)

	w := httptest.NewRecorder()
	testHandler.Delete(w, req)
	res := w.Result()
	body, _ := ioutil.ReadAll(res.Body)

	assert.Equal(t, fmt.Sprintf(`{"status":false,"message":"%s"}`, library.ErrInternalServerError), string(body))
	assert.Equal(t, http.StatusInternalServerError, res.StatusCode, "%d status is not equal to %d", http.StatusInternalServerError, res.StatusCode)
}

func Test_TestUpdate(t *testing.T) {
	service := new(mocks.TestService)
	repository := new(mocks.TestRepository)

	test := &models.Test{ID: 1, Name: "test"}
	newTest := &models.Test{ID: 1, Name: "test2"}

	repository.On("Update", newTest).Return(nil)
	testHandler := NewTestHandler(service, repository)
	req := httptest.NewRequest(http.MethodPut, "/test/1", strings.NewReader(`{"id":1,"name":"test2"}`))
	ctx := context.WithValue(req.Context(), middlewares.TestCtxKey, test)
	req = req.WithContext(ctx)

	w := httptest.NewRecorder()
	testHandler.Update(w, req)
	res := w.Result()
	body, _ := ioutil.ReadAll(res.Body)

	newTestString, _ := json.Marshal(newTest)

	assert.Equal(t, fmt.Sprintf(`{"status":true,"data":%s}`, newTestString), string(body))
	assert.Equal(t, http.StatusOK, res.StatusCode, "%d status is not equal to %d", http.StatusOK, res.StatusCode)
}

func Test_TestUpdate_NotFound(t *testing.T) {
	repository := new(mocks.TestRepository)
	service := new(mocks.TestService)

	testHandler := NewTestHandler(service, repository)
	repository.On("Update", &models.Test{}).Return(nil)
	res, body := makeRequest("/test/1", http.MethodPut, testHandler.Update, strings.NewReader(`{"id":1,"name":"test2"}`))

	assert.Equal(t, fmt.Sprintf(`{"status":false,"message":"%s"}`, library.ErrUnprocessableEntity), string(body))
	assert.Equal(t, http.StatusUnprocessableEntity, res.StatusCode, "%d status is not equal to %d", http.StatusUnprocessableEntity, res.StatusCode)
}

func Test_TestUpdate_ParseError(t *testing.T) {
	service := new(mocks.TestService)
	repository := new(mocks.TestRepository)

	test := &models.Test{ID: 1, Name: "test"}
	newTest := &models.Test{ID: 1, Name: "test2"}

	repository.On("Update", newTest).Return(nil)
	testHandler := NewTestHandler(service, repository)
	req := httptest.NewRequest(http.MethodPut, "/test/1", strings.NewReader(`{"id":1,"name":"test2"`))
	ctx := context.WithValue(req.Context(), middlewares.TestCtxKey, test)
	req = req.WithContext(ctx)

	w := httptest.NewRecorder()
	testHandler.Update(w, req)
	res := w.Result()
	body, _ := ioutil.ReadAll(res.Body)

	assert.Equal(t, fmt.Sprintf(`{"status":false,"message":"%s"}`, library.ErrBadRequest), string(body))
	assert.Equal(t, http.StatusBadRequest, res.StatusCode, "%d status is not equal to %d", res.StatusCode, http.StatusOK)
}

func Test_TestUpdate_DBError(t *testing.T) {
	service := new(mocks.TestService)
	repository := new(mocks.TestRepository)

	test := &models.Test{ID: 1, Name: "test"}
	newTest := &models.Test{ID: 1, Name: "test2"}

	repository.On("Update", newTest).Return(gorm.ErrNotImplemented)
	testHandler := NewTestHandler(service, repository)
	req := httptest.NewRequest(http.MethodPut, "/test/1", strings.NewReader(`{"id":1,"name":"test2"}`))
	ctx := context.WithValue(req.Context(), middlewares.TestCtxKey, test)
	req = req.WithContext(ctx)

	w := httptest.NewRecorder()
	testHandler.Update(w, req)
	res := w.Result()
	body, _ := ioutil.ReadAll(res.Body)

	assert.Equal(t, fmt.Sprintf(`{"status":false,"message":"%s"}`, library.ErrInternalServerError), string(body))
	assert.Equal(t, http.StatusInternalServerError, res.StatusCode, "%d status is not equal to %d", http.StatusInternalServerError, res.StatusCode)
}

func Test_TestGet(t *testing.T) {
	repository := new(mocks.TestRepository)
	service := new(mocks.TestService)

	test := &models.Test{ID: 1, Name: "test"}
	testHandler := NewTestHandler(service, repository)
	req := httptest.NewRequest(http.MethodGet, "/test/1", nil)
	ctx := context.WithValue(req.Context(), middlewares.TestCtxKey, test)
	req = req.WithContext(ctx)

	w := httptest.NewRecorder()
	testHandler.Get(w, req)
	res := w.Result()
	body, _ := ioutil.ReadAll(res.Body)

	testString, _ := json.Marshal(test)
	assert.Equal(t, fmt.Sprintf(`{"status":true,"data":%s}`, testString), string(body))
	assert.Equal(t, http.StatusOK, res.StatusCode, "%d status is not equal to %d", http.StatusOK, res.StatusCode)
}

func Test_TestGet_NotFound(t *testing.T) {
	repository := new(mocks.TestRepository)
	service := new(mocks.TestService)

	testHandler := NewTestHandler(service, repository)
	req := httptest.NewRequest(http.MethodGet, "/test/1", nil)

	w := httptest.NewRecorder()
	testHandler.Get(w, req)
	res := w.Result()
	body, _ := ioutil.ReadAll(res.Body)

	assert.Equal(t, fmt.Sprintf(`{"status":false,"message":"%s"}`, library.ErrUnprocessableEntity), string(body))
	assert.Equal(t, http.StatusUnprocessableEntity, res.StatusCode, "%d status is not equal to %d", http.StatusUnprocessableEntity, res.StatusCode)
}

func Test_TestList(t *testing.T) {
	repository := new(mocks.TestRepository)
	service := new(mocks.TestService)

	tests := []models.Test{}
	qb := &library.QueryBuilder{}
	total := int64(2)
	repository.On("List", qb, "").Return(tests, total, nil)
	testHandler := NewTestHandler(service, repository)
	req := httptest.NewRequest(http.MethodGet, "/test", nil)
	ctx := context.WithValue(req.Context(), middlewares.QueryCtxKey, qb)
	req = req.WithContext(ctx)

	w := httptest.NewRecorder()
	testHandler.List(w, req)
	res := w.Result()
	body, _ := ioutil.ReadAll(res.Body)

	testsString, _ := json.Marshal(tests)
	assert.Equal(t, fmt.Sprintf(`{"status":true,"data":{"data":%s,"total":%d}}`, testsString, total), string(body))
	assert.Equal(t, http.StatusOK, res.StatusCode, "%d status is not equal to %d", http.StatusOK, res.StatusCode)
}

func Test_TestList_QBNotFound(t *testing.T) {
	repository := new(mocks.TestRepository)
	service := new(mocks.TestService)

	repository.On("List", nil, "").Return(nil, 0, nil)
	testHandler := NewTestHandler(service, repository)
	req := httptest.NewRequest(http.MethodGet, "/test", nil)

	w := httptest.NewRecorder()
	testHandler.List(w, req)
	res := w.Result()
	body, _ := ioutil.ReadAll(res.Body)

	assert.Equal(t, fmt.Sprintf(`{"status":false,"message":"%s"}`, library.ErrUnprocessableEntity), string(body))
	assert.Equal(t, http.StatusUnprocessableEntity, res.StatusCode, "%d status is not equal to %d", http.StatusUnprocessableEntity, res.StatusCode)
}

func Test_TestList_DBError(t *testing.T) {
	repository := new(mocks.TestRepository)
	service := new(mocks.TestService)

	qb := &library.QueryBuilder{}
	repository.On("List", qb, "").Return(nil, int64(0), gorm.ErrNotImplemented)
	testHandler := NewTestHandler(service, repository)
	req := httptest.NewRequest(http.MethodGet, "/test", nil)
	ctx := context.WithValue(req.Context(), middlewares.QueryCtxKey, qb)
	req = req.WithContext(ctx)

	w := httptest.NewRecorder()
	testHandler.List(w, req)
	res := w.Result()
	body, _ := ioutil.ReadAll(res.Body)

	assert.Equal(t, fmt.Sprintf(`{"status":false,"message":"%s"}`, library.ErrInternalServerError), string(body))
	assert.Equal(t, http.StatusInternalServerError, res.StatusCode, "%d status is not equal to %d", http.StatusInternalServerError, res.StatusCode)
}

func Test_TestList_DBRecordNotFound(t *testing.T) {
	repository := new(mocks.TestRepository)
	service := new(mocks.TestService)

	qb := &library.QueryBuilder{}
	repository.On("List", qb, "").Return(nil, int64(0), gorm.ErrRecordNotFound)
	testHandler := NewTestHandler(service, repository)
	req := httptest.NewRequest(http.MethodGet, "/test", nil)
	ctx := context.WithValue(req.Context(), middlewares.QueryCtxKey, qb)
	req = req.WithContext(ctx)

	w := httptest.NewRecorder()
	testHandler.List(w, req)
	res := w.Result()
	body, _ := ioutil.ReadAll(res.Body)

	assert.Equal(t, fmt.Sprintf(`{"status":false,"message":"%s"}`, library.ErrNotFound), string(body))
	assert.Equal(t, http.StatusNotFound, res.StatusCode, "%d status is not equal to %d", http.StatusInternalServerError, res.StatusCode)
}

func Test_ListByTestGroupID(t *testing.T) {
	repository := new(mocks.TestRepository)
	service := new(mocks.TestService)

	tests := []models.Test{}
	qb := &library.QueryBuilder{}
	testGroup := &models.TestGroup{ID: 1}
	total := int64(2)
	repository.On("List", qb, "test_group_id=?", testGroup.ID).Return(tests, total, nil)
	testHandler := NewTestHandler(service, repository)

	req := httptest.NewRequest(http.MethodGet, "/test_group/1/tests", nil)
	ctxQB := context.WithValue(req.Context(), middlewares.QueryCtxKey, qb)
	ctx := context.WithValue(ctxQB, middlewares.TestGroupCtxKey, testGroup)
	req = req.WithContext(ctx)

	w := httptest.NewRecorder()
	testHandler.ListByTestGroupID(w, req)
	res := w.Result()
	body, _ := ioutil.ReadAll(res.Body)

	testsString, _ := json.Marshal(tests)
	assert.Equal(t, fmt.Sprintf(`{"status":true,"data":{"data":%s,"total":%d}}`, testsString, total), string(body))
	assert.Equal(t, http.StatusOK, res.StatusCode, "%d status is not equal to %d", http.StatusOK, res.StatusCode)

}

func Test_ListByTestGroupID_QBNotFound(t *testing.T) {
	repository := new(mocks.TestRepository)
	service := new(mocks.TestService)

	tests := []models.Test{}
	qb := &library.QueryBuilder{}
	testGroup := &models.TestGroup{ID: 1}
	total := int64(2)
	repository.On("List", qb, "test_group_id=?", testGroup.ID).Return(tests, total, nil)
	testHandler := NewTestHandler(service, repository)

	req := httptest.NewRequest(http.MethodGet, "/test_group/1/tests", nil)
	ctx := context.WithValue(req.Context(), middlewares.TestGroupCtxKey, testGroup)
	req = req.WithContext(ctx)

	w := httptest.NewRecorder()
	testHandler.ListByTestGroupID(w, req)
	res := w.Result()
	body, _ := ioutil.ReadAll(res.Body)

	assert.Equal(t, fmt.Sprintf(`{"status":false,"message":"%s"}`, library.ErrUnprocessableEntity), string(body))
	assert.Equal(t, http.StatusUnprocessableEntity, res.StatusCode, "%d status is not equal to %d", http.StatusUnprocessableEntity, res.StatusCode)
}

func Test_ListByTestGroupID_TGNotFound(t *testing.T) {
	repository := new(mocks.TestRepository)
	service := new(mocks.TestService)

	tests := []models.Test{}
	qb := &library.QueryBuilder{}
	testGroup := &models.TestGroup{ID: 1}
	total := int64(2)
	repository.On("List", qb, "test_group_id=?", testGroup.ID).Return(tests, total, nil)
	testHandler := NewTestHandler(service, repository)

	req := httptest.NewRequest(http.MethodGet, "/test_group/1/tests", nil)
	ctx := context.WithValue(req.Context(), middlewares.QueryCtxKey, qb)
	req = req.WithContext(ctx)

	w := httptest.NewRecorder()
	testHandler.ListByTestGroupID(w, req)
	res := w.Result()
	body, _ := ioutil.ReadAll(res.Body)

	assert.Equal(t, fmt.Sprintf(`{"status":false,"message":"%s"}`, library.ErrUnprocessableEntity), string(body))
	assert.Equal(t, http.StatusUnprocessableEntity, res.StatusCode, "%d status is not equal to %d", http.StatusUnprocessableEntity, res.StatusCode)
}

func Test_ListByTestGroupID_DBRNotFound(t *testing.T) {
	repository := new(mocks.TestRepository)
	service := new(mocks.TestService)

	qb := &library.QueryBuilder{}
	testGroup := &models.TestGroup{ID: 1}
	total := int64(2)
	repository.On("List", qb, "test_group_id=?", testGroup.ID).Return(nil, total, gorm.ErrRecordNotFound)
	testHandler := NewTestHandler(service, repository)

	req := httptest.NewRequest(http.MethodGet, "/test_group/1/tests", nil)
	ctxQB := context.WithValue(req.Context(), middlewares.QueryCtxKey, qb)
	ctx := context.WithValue(ctxQB, middlewares.TestGroupCtxKey, testGroup)
	req = req.WithContext(ctx)

	w := httptest.NewRecorder()
	testHandler.ListByTestGroupID(w, req)
	res := w.Result()
	body, _ := ioutil.ReadAll(res.Body)

	assert.Equal(t, fmt.Sprintf(`{"status":false,"message":"%s"}`, library.ErrNotFound), string(body))
	assert.Equal(t, http.StatusNotFound, res.StatusCode, "%d status is not equal to %d", http.StatusInternalServerError, res.StatusCode)
}

func Test_ListByTestGroupID_DBError(t *testing.T) {
	repository := new(mocks.TestRepository)
	service := new(mocks.TestService)

	qb := &library.QueryBuilder{}
	testGroup := &models.TestGroup{ID: 1}
	total := int64(2)
	repository.On("List", qb, "test_group_id=?", testGroup.ID).Return(nil, total, gorm.ErrNotImplemented)
	testHandler := NewTestHandler(service, repository)

	req := httptest.NewRequest(http.MethodGet, "/test_group/1/tests", nil)
	ctxQB := context.WithValue(req.Context(), middlewares.QueryCtxKey, qb)
	ctx := context.WithValue(ctxQB, middlewares.TestGroupCtxKey, testGroup)
	req = req.WithContext(ctx)

	w := httptest.NewRecorder()
	testHandler.ListByTestGroupID(w, req)
	res := w.Result()
	body, _ := ioutil.ReadAll(res.Body)

	assert.Equal(t, fmt.Sprintf(`{"status":false,"message":"%s"}`, library.ErrInternalServerError), string(body))
	assert.Equal(t, http.StatusInternalServerError, res.StatusCode, "%d status is not equal to %d", http.StatusInternalServerError, res.StatusCode)
}

func Test_Start(t *testing.T) {
	repository := new(mocks.TestRepository)
	service := new(mocks.TestService)

	runTest := &models.RunTest{ID: 1, TestID: 1}
	test := &models.Test{ID: 1, Name: "test"}
	service.On("Start", test).Return(runTest, nil)
	testHandler := NewTestHandler(service, repository)
	req := httptest.NewRequest(http.MethodGet, "/test/1", nil)
	ctx := context.WithValue(req.Context(), middlewares.TestCtxKey, test)
	req = req.WithContext(ctx)

	w := httptest.NewRecorder()
	testHandler.Start(w, req)
	res := w.Result()
	body, _ := ioutil.ReadAll(res.Body)

	runTestString, _ := json.Marshal(runTest)
	assert.Equal(t, fmt.Sprintf(`{"status":true,"data":%s}`, runTestString), string(body))
	assert.Equal(t, http.StatusOK, res.StatusCode, "%d status is not equal to %d", http.StatusOK, res.StatusCode)
}

func Test_Start_TCtxNotfound(t *testing.T) {
	repository := new(mocks.TestRepository)
	service := new(mocks.TestService)

	runTest := &models.RunTest{ID: 1, TestID: 1}
	test := &models.Test{ID: 1, Name: "test"}
	service.On("Start", test).Return(runTest, nil)
	testHandler := NewTestHandler(service, repository)
	req := httptest.NewRequest(http.MethodGet, "/test/1", nil)

	w := httptest.NewRecorder()
	testHandler.Start(w, req)
	res := w.Result()
	body, _ := ioutil.ReadAll(res.Body)

	assert.Equal(t, fmt.Sprintf(`{"status":false,"message":"%s"}`, library.ErrUnprocessableEntity), string(body))
	assert.Equal(t, http.StatusUnprocessableEntity, res.StatusCode, "%d status is not equal to %d", http.StatusUnprocessableEntity, res.StatusCode)
}

func Test_Start_ServiceError(t *testing.T) {
	repository := new(mocks.TestRepository)
	service := new(mocks.TestService)

	test := &models.Test{ID: 1, Name: "test"}
	service.On("Start", test).Return(nil, errors.New(library.ErrInternalServerError.Error()))
	testHandler := NewTestHandler(service, repository)
	req := httptest.NewRequest(http.MethodGet, "/test/1", nil)
	ctx := context.WithValue(req.Context(), middlewares.TestCtxKey, test)
	req = req.WithContext(ctx)

	w := httptest.NewRecorder()
	testHandler.Start(w, req)
	res := w.Result()
	body, _ := ioutil.ReadAll(res.Body)

	assert.Equal(t, fmt.Sprintf(`{"status":false,"message":"%s"}`, library.ErrInternalServerError), string(body))
	assert.Equal(t, http.StatusInternalServerError, res.StatusCode, "%d status is not equal to %d", http.StatusInternalServerError, res.StatusCode)
}
