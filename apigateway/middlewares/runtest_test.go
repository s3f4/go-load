package middlewares

import (
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi"
	"github.com/s3f4/go-load/apigateway/library"
	"github.com/s3f4/go-load/apigateway/mocks"
	"github.com/s3f4/go-load/apigateway/models"
	"github.com/stretchr/testify/assert"
)

func Test_RunTest(t *testing.T) {
	rtr := new(mocks.RunTestRepository)
	m := NewMiddleware(nil, nil, nil, nil, rtr)
	rtr.On("Get", uint(1)).Return(&models.RunTest{ID: 1}, nil)

	var next http.HandlerFunc
	next = func(w http.ResponseWriter, r *http.Request) {
		val, ok := r.Context().Value(RunTestCtxKey).(*models.RunTest)
		if !ok {
			t.Error("Run test not found")
		}
		assert.Equal(t, uint(1), val.ID)
	}

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	routeContext := chi.NewRouteContext()
	routeContext.URLParams.Add("ID", "1")
	ctx := context.WithValue(req.Context(), chi.RouteCtxKey, routeContext)

	req = req.WithContext(ctx)
	res := httptest.NewRecorder()

	test := m.RunTestCtx(next)
	test.ServeHTTP(res, req)
}

func Test_RunTest_ParamError(t *testing.T) {
	rtr := new(mocks.RunTestRepository)
	m := NewMiddleware(nil, nil, nil, nil, rtr)

	var next http.HandlerFunc
	next = func(w http.ResponseWriter, r *http.Request) {
	}

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	routeContext := chi.NewRouteContext()
	routeContext.URLParams.Add("ID", "abc")
	ctx := context.WithValue(req.Context(), chi.RouteCtxKey, routeContext)

	req = req.WithContext(ctx)
	w := httptest.NewRecorder()

	test := m.RunTestCtx(next)
	test.ServeHTTP(w, req)

	res := w.Result()
	body, _ := ioutil.ReadAll(res.Body)

	assert.Equal(t, fmt.Sprintf(`{"status":false,"message":"%s"}`, library.ErrNotFound), string(body))
	assert.Equal(t, http.StatusNotFound, res.StatusCode, "%d status is not equal to %d", http.StatusNotFound, res.StatusCode)
}

func Test_RunTest_NotFound(t *testing.T) {
	rtr := new(mocks.RunTestRepository)
	m := NewMiddleware(nil, nil, nil, nil, rtr)
	rtr.On("Get", uint(1)).Return(nil, errors.New(""))

	var next http.HandlerFunc
	next = func(w http.ResponseWriter, r *http.Request) {
	}

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	routeContext := chi.NewRouteContext()
	routeContext.URLParams.Add("ID", "1")
	ctx := context.WithValue(req.Context(), chi.RouteCtxKey, routeContext)

	req = req.WithContext(ctx)
	w := httptest.NewRecorder()

	test := m.RunTestCtx(next)
	test.ServeHTTP(w, req)

	res := w.Result()
	body, _ := ioutil.ReadAll(res.Body)

	assert.Equal(t, fmt.Sprintf(`{"status":false,"message":"%s"}`, library.ErrNotFound), string(body))
	assert.Equal(t, http.StatusNotFound, res.StatusCode, "%d status is not equal to %d", http.StatusNotFound, res.StatusCode)
}
