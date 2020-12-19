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

	"github.com/s3f4/go-load/apigateway/library"
	"github.com/s3f4/go-load/apigateway/middlewares"
	"github.com/s3f4/go-load/apigateway/mocks"
	"github.com/s3f4/go-load/apigateway/models"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func Test_RunTest_Create(t *testing.T) {
	repository := new(mocks.RunTestRepository)
	repository.On("Create", &models.RunTest{ID: 1}).Return(nil)
	runTestHandler := NewRunTestHandler(repository)

	res, body := makeRequest("/run_test", http.MethodPost, runTestHandler.Create, strings.NewReader(`{"id":1}`))
	assert.Equal(t, `{"status":true,"data":{"id":1,"test_id":0,"test":null,"start_time":null,"end_time":null,"passed":null}}`, string(body))
	assert.Equal(t, res.StatusCode, http.StatusOK, "%d status is not equal to %d", res.StatusCode, http.StatusOK)
}

func Test_RunTest_Create_ParseError(t *testing.T) {
	repository := new(mocks.RunTestRepository)
	runTestHandler := NewRunTestHandler(repository)

	res, body := makeRequest("/run_test", http.MethodPost, runTestHandler.Create, strings.NewReader(`{"id":1`))
	assert.Equal(t, fmt.Sprintf(`{"status":false,"message":"%s"}`, library.ErrBadRequest), string(body))
	assert.Equal(t, http.StatusBadRequest, res.StatusCode, "%d status is not equal to %d", http.StatusBadRequest, res.StatusCode)
}

func Test_RunTest_Create_DBError(t *testing.T) {
	repository := new(mocks.RunTestRepository)

	repository.On("Create", &models.RunTest{ID: 1}).Return(gorm.ErrNotImplemented)
	runTestHandler := NewRunTestHandler(repository)

	res, body := makeRequest("/run_test", http.MethodPost, runTestHandler.Create, strings.NewReader(`{"id":1}`))

	assert.Equal(t, fmt.Sprintf(`{"status":false,"message":"%s"}`, library.ErrInternalServerError), string(body))
	assert.Equal(t, http.StatusInternalServerError, res.StatusCode, "%d status is not equal to %d", http.StatusInternalServerError, res.StatusCode)
}

func Test_RunTest_Delete(t *testing.T) {
	repository := new(mocks.RunTestRepository)

	runTest := &models.RunTest{ID: 1}
	repository.On("Delete", runTest).Return(nil)
	runTestHandler := NewRunTestHandler(repository)

	req := httptest.NewRequest(http.MethodDelete, "/run_test/1", nil)
	ctx := context.WithValue(req.Context(), middlewares.RunTestCtxKey, runTest)
	req = req.WithContext(ctx)

	w := httptest.NewRecorder()
	runTestHandler.Delete(w, req)
	res := w.Result()
	body, _ := ioutil.ReadAll(res.Body)

	testStr, _ := json.Marshal(runTest)
	assert.Equal(t, fmt.Sprintf(`{"status":true,"data":%s}`, string(testStr)), string(body))
	assert.Equal(t, http.StatusOK, res.StatusCode, "%d status is not equal to %d", http.StatusOK, res.StatusCode)
}

func Test_RunTest_Delete_RunTestCTXNotFound(t *testing.T) {
	repository := new(mocks.RunTestRepository)

	runTestHandler := NewRunTestHandler(repository)
	res, body := makeRequest("/run_test/1", http.MethodDelete, runTestHandler.Delete, strings.NewReader(`{"":"test"}`))

	assert.Equal(t, fmt.Sprintf(`{"status":false,"message":"%s"}`, library.ErrUnprocessableEntity), string(body))
	assert.Equal(t, http.StatusUnprocessableEntity, res.StatusCode, "%d status is not equal to %d", http.StatusUnprocessableEntity, res.StatusCode)
}

func Test_RunTest_Delete_DBError(t *testing.T) {
	repository := new(mocks.RunTestRepository)

	repository.On("Delete", &models.RunTest{}).Return(gorm.ErrNotImplemented)
	runTestHandler := NewRunTestHandler(repository)
	test := &models.RunTest{}
	req := httptest.NewRequest(http.MethodDelete, "/run_test/1", nil)
	ctx := context.WithValue(req.Context(), middlewares.RunTestCtxKey, test)
	req = req.WithContext(ctx)

	w := httptest.NewRecorder()
	runTestHandler.Delete(w, req)
	res := w.Result()
	body, _ := ioutil.ReadAll(res.Body)

	assert.Equal(t, fmt.Sprintf(`{"status":false,"message":"%s"}`, library.ErrInternalServerError), string(body))
	assert.Equal(t, http.StatusInternalServerError, res.StatusCode, "%d status is not equal to %d", http.StatusInternalServerError, res.StatusCode)
}

func Test_RunTest_Get(t *testing.T) {
	repository := new(mocks.RunTestRepository)

	test := &models.RunTest{ID: 1}
	runTestHandler := NewRunTestHandler(repository)
	req := httptest.NewRequest(http.MethodGet, "/run_test/1", nil)
	ctx := context.WithValue(req.Context(), middlewares.RunTestCtxKey, test)
	req = req.WithContext(ctx)

	w := httptest.NewRecorder()
	runTestHandler.Get(w, req)
	res := w.Result()
	body, _ := ioutil.ReadAll(res.Body)

	runTestString, _ := json.Marshal(test)
	assert.Equal(t, fmt.Sprintf(`{"status":true,"data":%s}`, runTestString), string(body))
	assert.Equal(t, http.StatusOK, res.StatusCode, "%d status is not equal to %d", http.StatusOK, res.StatusCode)
}

func Test_RunTest_Get_RunTestCTXNotFound(t *testing.T) {
	repository := new(mocks.RunTestRepository)

	runTestHandler := NewRunTestHandler(repository)
	req := httptest.NewRequest(http.MethodGet, "/run_test/1", nil)

	w := httptest.NewRecorder()
	runTestHandler.Get(w, req)
	res := w.Result()
	body, _ := ioutil.ReadAll(res.Body)

	assert.Equal(t, fmt.Sprintf(`{"status":false,"message":"%s"}`, library.ErrUnprocessableEntity), string(body))
	assert.Equal(t, http.StatusUnprocessableEntity, res.StatusCode, "%d status is not equal to %d", http.StatusUnprocessableEntity, res.StatusCode)
}

func Test_RunTest_List(t *testing.T) {
	repository := new(mocks.RunTestRepository)

	tests := []models.RunTest{}
	qb := &library.QueryBuilder{}
	total := int64(2)
	repository.On("List", qb, "").Return(tests, total, nil)
	runTestHandler := NewRunTestHandler(repository)
	req := httptest.NewRequest(http.MethodGet, "/run_test", nil)
	ctx := context.WithValue(req.Context(), middlewares.QueryCtxKey, qb)
	req = req.WithContext(ctx)

	w := httptest.NewRecorder()
	runTestHandler.List(w, req)
	res := w.Result()
	body, _ := ioutil.ReadAll(res.Body)

	testsString, _ := json.Marshal(tests)
	assert.Equal(t, fmt.Sprintf(`{"status":true,"data":{"data":%s,"total":%d}}`, testsString, total), string(body))
	assert.Equal(t, http.StatusOK, res.StatusCode, "%d status is not equal to %d", http.StatusOK, res.StatusCode)
}

func Test_RunTest_List_QueryBuilderCTXNotFound(t *testing.T) {
	repository := new(mocks.RunTestRepository)

	repository.On("List", nil, "").Return(nil, 0, nil)
	runTestHandler := NewRunTestHandler(repository)
	req := httptest.NewRequest(http.MethodGet, "/run_test", nil)

	w := httptest.NewRecorder()
	runTestHandler.List(w, req)
	res := w.Result()
	body, _ := ioutil.ReadAll(res.Body)

	assert.Equal(t, fmt.Sprintf(`{"status":false,"message":"%s"}`, library.ErrUnprocessableEntity), string(body))
	assert.Equal(t, http.StatusUnprocessableEntity, res.StatusCode, "%d status is not equal to %d", http.StatusUnprocessableEntity, res.StatusCode)
}

func Test_RunTest_List_DBError(t *testing.T) {
	repository := new(mocks.RunTestRepository)

	qb := &library.QueryBuilder{}
	repository.On("List", qb, "").Return(nil, int64(0), gorm.ErrNotImplemented)
	runTestHandler := NewRunTestHandler(repository)
	req := httptest.NewRequest(http.MethodGet, "/run_test", nil)
	ctx := context.WithValue(req.Context(), middlewares.QueryCtxKey, qb)
	req = req.WithContext(ctx)

	w := httptest.NewRecorder()
	runTestHandler.List(w, req)
	res := w.Result()
	body, _ := ioutil.ReadAll(res.Body)

	assert.Equal(t, fmt.Sprintf(`{"status":false,"message":"%s"}`, library.ErrInternalServerError), string(body))
	assert.Equal(t, http.StatusInternalServerError, res.StatusCode, "%d status is not equal to %d", http.StatusInternalServerError, res.StatusCode)
}

func Test_RunTest_List_DBRecordNotFound(t *testing.T) {
	repository := new(mocks.RunTestRepository)

	qb := &library.QueryBuilder{}
	repository.On("List", qb, "").Return(nil, int64(0), gorm.ErrRecordNotFound)
	runTestHandler := NewRunTestHandler(repository)
	req := httptest.NewRequest(http.MethodGet, "/run_test", nil)
	ctx := context.WithValue(req.Context(), middlewares.QueryCtxKey, qb)
	req = req.WithContext(ctx)

	w := httptest.NewRecorder()
	runTestHandler.List(w, req)
	res := w.Result()
	body, _ := ioutil.ReadAll(res.Body)

	assert.Equal(t, fmt.Sprintf(`{"status":false,"message":"%s"}`, library.ErrNotFound), string(body))
	assert.Equal(t, http.StatusNotFound, res.StatusCode, "%d status is not equal to %d", http.StatusInternalServerError, res.StatusCode)
}

func Test_RunTest_ListByTestID(t *testing.T) {
	repository := new(mocks.RunTestRepository)

	tests := []models.RunTest{}
	qb := &library.QueryBuilder{}
	test := &models.Test{ID: 1}
	total := int64(2)
	repository.On("List", qb, "test_id=?", test.ID).Return(tests, total, nil)
	runTestHandler := NewRunTestHandler(repository)

	req := httptest.NewRequest(http.MethodGet, "/run_test_group/1/run_tests", nil)
	ctxQB := context.WithValue(req.Context(), middlewares.QueryCtxKey, qb)
	ctx := context.WithValue(ctxQB, middlewares.TestCtxKey, test)
	req = req.WithContext(ctx)

	w := httptest.NewRecorder()
	runTestHandler.ListByTestID(w, req)
	res := w.Result()
	body, _ := ioutil.ReadAll(res.Body)

	testsString, _ := json.Marshal(tests)
	assert.Equal(t, fmt.Sprintf(`{"status":true,"data":{"data":%s,"total":%d}}`, testsString, total), string(body))
	assert.Equal(t, http.StatusOK, res.StatusCode, "%d status is not equal to %d", http.StatusOK, res.StatusCode)

}

func Test_RunTest_ListByTestID_QueryBuidlerCTXNotFound(t *testing.T) {
	repository := new(mocks.RunTestRepository)

	tests := []models.RunTest{}
	qb := &library.QueryBuilder{}
	test := &models.Test{ID: 1}
	total := int64(2)
	repository.On("List", qb, "test_id=?", test.ID).Return(tests, total, nil)
	runTestHandler := NewRunTestHandler(repository)

	req := httptest.NewRequest(http.MethodGet, "/run_test_group/1/run_tests", nil)
	ctx := context.WithValue(req.Context(), middlewares.TestCtxKey, test)
	req = req.WithContext(ctx)

	w := httptest.NewRecorder()
	runTestHandler.ListByTestID(w, req)
	res := w.Result()
	body, _ := ioutil.ReadAll(res.Body)

	assert.Equal(t, fmt.Sprintf(`{"status":false,"message":"%s"}`, library.ErrUnprocessableEntity), string(body))
	assert.Equal(t, http.StatusUnprocessableEntity, res.StatusCode, "%d status is not equal to %d", http.StatusUnprocessableEntity, res.StatusCode)
}

func Test_RunTest_ListByTestID_testCTXNotFound(t *testing.T) {
	repository := new(mocks.RunTestRepository)

	tests := []models.RunTest{}
	qb := &library.QueryBuilder{}
	test := &models.Test{ID: 1}
	total := int64(2)
	repository.On("List", qb, "test_id=?", test.ID).Return(tests, total, nil)
	runTestHandler := NewRunTestHandler(repository)

	req := httptest.NewRequest(http.MethodGet, "/run_test_group/1/run_tests", nil)
	ctx := context.WithValue(req.Context(), middlewares.QueryCtxKey, qb)
	req = req.WithContext(ctx)

	w := httptest.NewRecorder()
	runTestHandler.ListByTestID(w, req)
	res := w.Result()
	body, _ := ioutil.ReadAll(res.Body)

	assert.Equal(t, fmt.Sprintf(`{"status":false,"message":"%s"}`, library.ErrUnprocessableEntity), string(body))
	assert.Equal(t, http.StatusUnprocessableEntity, res.StatusCode, "%d status is not equal to %d", http.StatusUnprocessableEntity, res.StatusCode)
}

func Test_RunTest_ListByTestID_DBRecordNotFound(t *testing.T) {
	repository := new(mocks.RunTestRepository)

	qb := &library.QueryBuilder{}
	test := &models.Test{ID: 1}
	total := int64(2)
	repository.On("List", qb, "test_id=?", test.ID).Return(nil, total, gorm.ErrRecordNotFound)
	runTestHandler := NewRunTestHandler(repository)

	req := httptest.NewRequest(http.MethodGet, "/run_test_group/1/run_tests", nil)
	ctxQB := context.WithValue(req.Context(), middlewares.QueryCtxKey, qb)
	ctx := context.WithValue(ctxQB, middlewares.TestCtxKey, test)
	req = req.WithContext(ctx)

	w := httptest.NewRecorder()
	runTestHandler.ListByTestID(w, req)
	res := w.Result()
	body, _ := ioutil.ReadAll(res.Body)

	assert.Equal(t, fmt.Sprintf(`{"status":false,"message":"%s"}`, library.ErrNotFound), string(body))
	assert.Equal(t, http.StatusNotFound, res.StatusCode, "%d status is not equal to %d", http.StatusInternalServerError, res.StatusCode)
}

func Test_RunTest_ListByTestID_DBError(t *testing.T) {
	repository := new(mocks.RunTestRepository)

	qb := &library.QueryBuilder{}
	test := &models.Test{ID: 1}
	total := int64(2)
	repository.On("List", qb, "test_id=?", test.ID).Return(nil, total, gorm.ErrNotImplemented)
	runTestHandler := NewRunTestHandler(repository)

	req := httptest.NewRequest(http.MethodGet, "/run_test_group/1/run_tests", nil)
	ctxQB := context.WithValue(req.Context(), middlewares.QueryCtxKey, qb)
	ctx := context.WithValue(ctxQB, middlewares.TestCtxKey, test)
	req = req.WithContext(ctx)

	w := httptest.NewRecorder()
	runTestHandler.ListByTestID(w, req)
	res := w.Result()
	body, _ := ioutil.ReadAll(res.Body)

	assert.Equal(t, fmt.Sprintf(`{"status":false,"message":"%s"}`, library.ErrInternalServerError), string(body))
	assert.Equal(t, http.StatusInternalServerError, res.StatusCode, "%d status is not equal to %d", http.StatusInternalServerError, res.StatusCode)
}
