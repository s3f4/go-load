package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/s3f4/go-load/apigateway/library"
	"github.com/s3f4/go-load/apigateway/middlewares"
	"github.com/s3f4/go-load/apigateway/mocks"
	"github.com/s3f4/go-load/apigateway/models"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func Test_Stats_List(t *testing.T) {
	repository := new(mocks.ResponseRepository)

	runTest := &models.RunTest{ID: 1}
	responses := []models.Response{}
	qb := &library.QueryBuilder{}
	total := int64(2)
	repository.On("List", qb, "run_test_id=?", runTest.ID).Return(responses, total, nil)

	statsHandler := NewStatsHandler(repository)
	req := httptest.NewRequest(http.MethodGet, "/stats", nil)
	ctx := context.WithValue(req.Context(), middlewares.QueryCtxKey, qb)
	ctxQB := context.WithValue(ctx, middlewares.RunTestCtxKey, runTest)
	req = req.WithContext(ctxQB)

	w := httptest.NewRecorder()
	statsHandler.List(w, req)
	res := w.Result()
	body, _ := ioutil.ReadAll(res.Body)

	responseString, _ := json.Marshal(responses)
	assert.Equal(t, fmt.Sprintf(`{"status":true,"data":{"data":%s,"total":%d}}`, responseString, total), string(body))
	assert.Equal(t, http.StatusOK, res.StatusCode, "%d status is not equal to %d", http.StatusOK, res.StatusCode)
}

func Test_Stats_List_RunTestCTXNotFound(t *testing.T) {
	repository := new(mocks.ResponseRepository)

	runTest := &models.RunTest{ID: 1}
	responses := []models.Response{}
	qb := &library.QueryBuilder{}
	total := int64(2)
	repository.On("List", qb, "run_test_id=?", runTest.ID).Return(responses, total, nil)

	statsHandler := NewStatsHandler(repository)
	req := httptest.NewRequest(http.MethodGet, "/stats", nil)
	ctx := context.WithValue(req.Context(), middlewares.QueryCtxKey, qb)
	req = req.WithContext(ctx)

	w := httptest.NewRecorder()
	statsHandler.List(w, req)
	res := w.Result()
	body, _ := ioutil.ReadAll(res.Body)

	assert.Equal(t, fmt.Sprintf(`{"status":false,"message":"%s"}`, library.ErrUnprocessableEntity), string(body))
	assert.Equal(t, http.StatusUnprocessableEntity, res.StatusCode, "%d status is not equal to %d", http.StatusUnprocessableEntity, res.StatusCode)
}

func Test_Stats_List_QueryBuilderCTXNotFound(t *testing.T) {
	repository := new(mocks.ResponseRepository)

	runTest := &models.RunTest{ID: 1}
	responses := []models.Response{}
	qb := &library.QueryBuilder{}
	total := int64(2)
	repository.On("List", qb, "run_test_id=?", runTest.ID).Return(responses, total, nil)

	statsHandler := NewStatsHandler(repository)
	req := httptest.NewRequest(http.MethodGet, "/stats", nil)
	ctx := context.WithValue(req.Context(), middlewares.RunTestCtxKey, runTest)
	req = req.WithContext(ctx)

	w := httptest.NewRecorder()
	statsHandler.List(w, req)
	res := w.Result()
	body, _ := ioutil.ReadAll(res.Body)

	assert.Equal(t, fmt.Sprintf(`{"status":false,"message":"%s"}`, library.ErrUnprocessableEntity), string(body))
	assert.Equal(t, http.StatusUnprocessableEntity, res.StatusCode, "%d status is not equal to %d", http.StatusUnprocessableEntity, res.StatusCode)
}

func Test_Stats_List_DBError(t *testing.T) {
	repository := new(mocks.ResponseRepository)

	runTest := &models.RunTest{ID: 1}
	qb := &library.QueryBuilder{}
	repository.On("List", qb, "run_test_id=?", runTest.ID).Return(nil, int64(0), gorm.ErrNotImplemented)

	statsHandler := NewStatsHandler(repository)
	req := httptest.NewRequest(http.MethodGet, "/stats", nil)
	ctx := context.WithValue(req.Context(), middlewares.QueryCtxKey, qb)
	ctxQB := context.WithValue(ctx, middlewares.RunTestCtxKey, runTest)
	req = req.WithContext(ctxQB)

	w := httptest.NewRecorder()
	statsHandler.List(w, req)
	res := w.Result()
	body, _ := ioutil.ReadAll(res.Body)

	assert.Equal(t, fmt.Sprintf(`{"status":false,"message":"%s"}`, library.ErrInternalServerError), string(body))
	assert.Equal(t, http.StatusInternalServerError, res.StatusCode, "%d status is not equal to %d", http.StatusInternalServerError, res.StatusCode)
}

func Test_Stats_List_DBRecordNotFound(t *testing.T) {
	repository := new(mocks.ResponseRepository)

	runTest := &models.RunTest{ID: 1}
	qb := &library.QueryBuilder{}
	repository.On("List", qb, "run_test_id=?", runTest.ID).Return(nil, int64(0), gorm.ErrRecordNotFound)

	statsHandler := NewStatsHandler(repository)
	req := httptest.NewRequest(http.MethodGet, "/stats", nil)
	ctx := context.WithValue(req.Context(), middlewares.QueryCtxKey, qb)
	ctxQB := context.WithValue(ctx, middlewares.RunTestCtxKey, runTest)
	req = req.WithContext(ctxQB)

	w := httptest.NewRecorder()
	statsHandler.List(w, req)
	res := w.Result()
	body, _ := ioutil.ReadAll(res.Body)

	assert.Equal(t, fmt.Sprintf(`{"status":false,"message":"%s"}`, library.ErrNotFound), string(body))
	assert.Equal(t, http.StatusNotFound, res.StatusCode, "%d status is not equal to %d", http.StatusInternalServerError, res.StatusCode)
}
