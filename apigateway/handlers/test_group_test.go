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

func TestGroup_Update(t *testing.T) {
	repository := new(mocks.TestGroupRepository)

	testGroup := &models.TestGroup{ID: 1, Name: "test"}
	newTestGroup := &models.TestGroup{ID: 1, Name: "test2"}
	repository.On("Update", newTestGroup).Return(nil)

	testGroupHandler := NewTestGroupHandler(repository)
	req := httptest.NewRequest(http.MethodPut, "/test_group/1", strings.NewReader(`{"id":1,"name":"test2"}`))
	ctx := context.WithValue(req.Context(), middlewares.TestGroupCtxKey, testGroup)
	req = req.WithContext(ctx)

	w := httptest.NewRecorder()
	testGroupHandler.Update(w, req)
	res := w.Result()
	body, _ := ioutil.ReadAll(res.Body)

	newTestGroupString, _ := json.Marshal(newTestGroup)

	assert.Equal(t, fmt.Sprintf(`{"status":true,"data":%s}`, newTestGroupString), string(body))
	assert.Equal(t, http.StatusOK, res.StatusCode, "%d status is not equal to %d", http.StatusOK, res.StatusCode)
}

func TestGroup_Update_TestGroupCTXNotFound(t *testing.T) {
	repository := new(mocks.TestGroupRepository)

	newTestGroup := &models.TestGroup{ID: 1, Name: "test2"}
	repository.On("Update", newTestGroup).Return(nil)

	testGroupHandler := NewTestGroupHandler(repository)
	req := httptest.NewRequest(http.MethodPut, "/test_group/1", strings.NewReader(`{"id":1,"name":"test2"}`))

	w := httptest.NewRecorder()
	testGroupHandler.Update(w, req)
	res := w.Result()
	body, _ := ioutil.ReadAll(res.Body)

	assert.Equal(t, fmt.Sprintf(`{"status":false,"message":"%s"}`, library.ErrUnprocessableEntity), string(body))
	assert.Equal(t, http.StatusUnprocessableEntity, res.StatusCode, "%d status is not equal to %d", http.StatusUnprocessableEntity, res.StatusCode)
}

func TestGroup_Update_ParseError(t *testing.T) {
	repository := new(mocks.TestGroupRepository)

	testGroup := &models.TestGroup{ID: 1, Name: "test"}
	newTestGroup := &models.TestGroup{ID: 1, Name: "test2"}
	repository.On("Update", newTestGroup).Return(nil)

	testGroupHandler := NewTestGroupHandler(repository)
	req := httptest.NewRequest(http.MethodPut, "/test_group/1", strings.NewReader(`{"id":1,"name":"test2`))
	ctx := context.WithValue(req.Context(), middlewares.TestGroupCtxKey, testGroup)
	req = req.WithContext(ctx)

	w := httptest.NewRecorder()
	testGroupHandler.Update(w, req)
	res := w.Result()
	body, _ := ioutil.ReadAll(res.Body)

	assert.Equal(t, fmt.Sprintf(`{"status":false,"message":"%s"}`, library.ErrBadRequest), string(body))
	assert.Equal(t, http.StatusBadRequest, res.StatusCode, "%d status is not equal to %d", res.StatusCode, http.StatusOK)
}

func TestGroup_Update_DBError(t *testing.T) {
	repository := new(mocks.TestGroupRepository)

	testGroup := &models.TestGroup{ID: 1, Name: "test"}
	newTestGroup := &models.TestGroup{ID: 1, Name: "test2"}
	repository.On("Update", newTestGroup).Return(gorm.ErrNotImplemented)

	testGroupHandler := NewTestGroupHandler(repository)
	req := httptest.NewRequest(http.MethodPut, "/test_group/1", strings.NewReader(`{"id":1,"name":"test2"}`))
	ctx := context.WithValue(req.Context(), middlewares.TestGroupCtxKey, testGroup)
	req = req.WithContext(ctx)

	w := httptest.NewRecorder()
	testGroupHandler.Update(w, req)
	res := w.Result()
	body, _ := ioutil.ReadAll(res.Body)

	assert.Equal(t, fmt.Sprintf(`{"status":false,"message":"%s"}`, library.ErrInternalServerError), string(body))
	assert.Equal(t, http.StatusInternalServerError, res.StatusCode, "%d status is not equal to %d", http.StatusInternalServerError, res.StatusCode)
}

func TestGroup_Delete(t *testing.T) {
	repository := new(mocks.TestGroupRepository)

	testGroup := &models.TestGroup{ID: 1, Name: "test"}
	repository.On("Delete", testGroup).Return(nil)

	testGroupHandler := NewTestGroupHandler(repository)
	req := httptest.NewRequest(http.MethodPut, "/test_group/1", strings.NewReader(`{"id":1,"name":"test2"}`))
	ctx := context.WithValue(req.Context(), middlewares.TestGroupCtxKey, testGroup)
	req = req.WithContext(ctx)

	w := httptest.NewRecorder()
	testGroupHandler.Delete(w, req)
	res := w.Result()
	body, _ := ioutil.ReadAll(res.Body)

	testGroupString, _ := json.Marshal(testGroup)

	assert.Equal(t, fmt.Sprintf(`{"status":true,"data":%s}`, testGroupString), string(body))
	assert.Equal(t, http.StatusOK, res.StatusCode, "%d status is not equal to %d", http.StatusOK, res.StatusCode)
}

func TestGroup_Delete_TestGroupCTXNotFound(t *testing.T) {
	repository := new(mocks.TestGroupRepository)

	testGroup := &models.TestGroup{ID: 1, Name: "test"}
	repository.On("Delete", testGroup).Return(nil)

	testGroupHandler := NewTestGroupHandler(repository)
	req := httptest.NewRequest(http.MethodPut, "/test_group/1", strings.NewReader(`{"id":1,"name":"test2"}`))

	w := httptest.NewRecorder()
	testGroupHandler.Delete(w, req)
	res := w.Result()
	body, _ := ioutil.ReadAll(res.Body)

	assert.Equal(t, fmt.Sprintf(`{"status":false,"message":"%s"}`, library.ErrUnprocessableEntity), string(body))
	assert.Equal(t, http.StatusUnprocessableEntity, res.StatusCode, "%d status is not equal to %d", http.StatusUnprocessableEntity, res.StatusCode)
}

func TestGroup_Delete_DBError(t *testing.T) {
	repository := new(mocks.TestGroupRepository)

	testGroup := &models.TestGroup{ID: 1, Name: "test"}
	repository.On("Delete", testGroup).Return(gorm.ErrNotImplemented)

	testGroupHandler := NewTestGroupHandler(repository)
	req := httptest.NewRequest(http.MethodPut, "/test_group/1", nil)
	ctx := context.WithValue(req.Context(), middlewares.TestGroupCtxKey, testGroup)
	req = req.WithContext(ctx)

	w := httptest.NewRecorder()
	testGroupHandler.Delete(w, req)
	res := w.Result()
	body, _ := ioutil.ReadAll(res.Body)

	assert.Equal(t, fmt.Sprintf(`{"status":false,"message":"%s"}`, library.ErrInternalServerError), string(body))
	assert.Equal(t, http.StatusInternalServerError, res.StatusCode, "%d status is not equal to %d", http.StatusInternalServerError, res.StatusCode)
}

func TestGroup_Get(t *testing.T) {
	repository := new(mocks.TestGroupRepository)

	testGroup := &models.TestGroup{ID: 1, Name: "test"}
	testGroupHandler := NewTestGroupHandler(repository)
	req := httptest.NewRequest(http.MethodGet, "/test_group/1", nil)
	ctx := context.WithValue(req.Context(), middlewares.TestGroupCtxKey, testGroup)
	req = req.WithContext(ctx)

	w := httptest.NewRecorder()
	testGroupHandler.Get(w, req)
	res := w.Result()
	body, _ := ioutil.ReadAll(res.Body)

	testString, _ := json.Marshal(testGroup)
	assert.Equal(t, fmt.Sprintf(`{"status":true,"data":%s}`, testString), string(body))
	assert.Equal(t, http.StatusOK, res.StatusCode, "%d status is not equal to %d", http.StatusOK, res.StatusCode)
}

func TestGroup_Get_TestGroupCTXNotFound(t *testing.T) {
	repository := new(mocks.TestGroupRepository)

	testGroupHandler := NewTestGroupHandler(repository)
	req := httptest.NewRequest(http.MethodGet, "/test_group/1", nil)

	w := httptest.NewRecorder()
	testGroupHandler.Get(w, req)
	res := w.Result()
	body, _ := ioutil.ReadAll(res.Body)

	assert.Equal(t, fmt.Sprintf(`{"status":false,"message":"%s"}`, library.ErrUnprocessableEntity), string(body))
	assert.Equal(t, http.StatusUnprocessableEntity, res.StatusCode, "%d status is not equal to %d", http.StatusUnprocessableEntity, res.StatusCode)
}

func TestGroup_List(t *testing.T) {
	repository := new(mocks.TestGroupRepository)

	testGroups := []models.TestGroup{}
	qb := &library.QueryBuilder{}
	total := int64(2)
	repository.On("List", qb, "").Return(testGroups, total, nil)

	testGroupHandler := NewTestGroupHandler(repository)
	req := httptest.NewRequest(http.MethodGet, "/test_group", nil)
	ctx := context.WithValue(req.Context(), middlewares.QueryCtxKey, qb)
	req = req.WithContext(ctx)

	w := httptest.NewRecorder()
	testGroupHandler.List(w, req)
	res := w.Result()
	body, _ := ioutil.ReadAll(res.Body)

	testGroupsString, _ := json.Marshal(testGroups)
	assert.Equal(t, fmt.Sprintf(`{"status":true,"data":{"data":%s,"total":%d}}`, testGroupsString, total), string(body))
	assert.Equal(t, http.StatusOK, res.StatusCode, "%d status is not equal to %d", http.StatusOK, res.StatusCode)
}

func TestGroup_List_QueryBuilderCTXNotFound(t *testing.T) {
	repository := new(mocks.TestGroupRepository)

	testGroups := []models.TestGroup{}
	qb := &library.QueryBuilder{}
	total := int64(2)
	repository.On("List", qb, "").Return(testGroups, total, nil)

	testGroupHandler := NewTestGroupHandler(repository)
	req := httptest.NewRequest(http.MethodGet, "/test_group", nil)

	w := httptest.NewRecorder()
	testGroupHandler.List(w, req)
	res := w.Result()
	body, _ := ioutil.ReadAll(res.Body)

	assert.Equal(t, fmt.Sprintf(`{"status":false,"message":"%s"}`, library.ErrUnprocessableEntity), string(body))
	assert.Equal(t, http.StatusUnprocessableEntity, res.StatusCode, "%d status is not equal to %d", http.StatusUnprocessableEntity, res.StatusCode)
}

func TestGroup_List_DBError(t *testing.T) {
	repository := new(mocks.TestGroupRepository)

	qb := &library.QueryBuilder{}
	repository.On("List", qb, "").Return(nil, int64(0), gorm.ErrNotImplemented)

	testGroupHandler := NewTestGroupHandler(repository)
	req := httptest.NewRequest(http.MethodGet, "/test_group", nil)
	ctx := context.WithValue(req.Context(), middlewares.QueryCtxKey, qb)
	req = req.WithContext(ctx)

	w := httptest.NewRecorder()
	testGroupHandler.List(w, req)
	res := w.Result()
	body, _ := ioutil.ReadAll(res.Body)

	assert.Equal(t, fmt.Sprintf(`{"status":false,"message":"%s"}`, library.ErrInternalServerError), string(body))
	assert.Equal(t, http.StatusInternalServerError, res.StatusCode, "%d status is not equal to %d", http.StatusInternalServerError, res.StatusCode)
}

func TestGroup_List_DBRecordNotFound(t *testing.T) {
	repository := new(mocks.TestGroupRepository)

	qb := &library.QueryBuilder{}
	repository.On("List", qb, "").Return(nil, int64(0), gorm.ErrRecordNotFound)

	testGroupHandler := NewTestGroupHandler(repository)
	req := httptest.NewRequest(http.MethodGet, "/test_group", nil)
	ctx := context.WithValue(req.Context(), middlewares.QueryCtxKey, qb)
	req = req.WithContext(ctx)

	w := httptest.NewRecorder()
	testGroupHandler.List(w, req)
	res := w.Result()
	body, _ := ioutil.ReadAll(res.Body)

	assert.Equal(t, fmt.Sprintf(`{"status":false,"message":"%s"}`, library.ErrNotFound), string(body))
	assert.Equal(t, http.StatusNotFound, res.StatusCode, "%d status is not equal to %d", http.StatusInternalServerError, res.StatusCode)
}
