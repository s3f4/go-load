package handlers

import (
	"errors"
	"net/http"

	"github.com/s3f4/go-load/apigateway/library"
	"github.com/s3f4/go-load/apigateway/library/log"
	res "github.com/s3f4/go-load/apigateway/library/response"
	"github.com/s3f4/go-load/apigateway/middlewares"
	"github.com/s3f4/go-load/apigateway/models"
	"github.com/s3f4/go-load/apigateway/repository"
	"github.com/s3f4/go-load/apigateway/services"
	"gorm.io/gorm"
)

// TestHandler interface
type TestHandler interface {
	Create(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Get(w http.ResponseWriter, r *http.Request)
	List(w http.ResponseWriter, r *http.Request)
	ListByTestGroupID(w http.ResponseWriter, r *http.Request)
	Start(w http.ResponseWriter, r *http.Request)
}

type testHandler struct {
	service    services.TestService
	repository repository.TestRepository
}

// NewTestHandler returns a new testHandler object
func NewTestHandler(service services.TestService, repository repository.TestRepository) TestHandler {
	return &testHandler{
		service:    service,
		repository: repository,
	}
}

func (h *testHandler) Create(w http.ResponseWriter, r *http.Request) {
	var test models.Test
	if err := parse(r, &test); err != nil {
		log.Debug(err)
		res.R400(w, r, library.ErrBadRequest)
		return
	}

	err := h.repository.Create(&test)
	if err != nil {
		log.Debug(err)
		res.R500(w, r, library.ErrInternalServerError)
		return
	}
	res.R200(w, r, test)
}

func (h *testHandler) Delete(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	test, ok := ctx.Value(middlewares.TestCtxKey).(*models.Test)
	if !ok {
		res.R422(w, r, library.ErrUnprocessableEntity)
		return
	}
	if err := h.repository.Delete(test); err != nil {
		log.Debug(err)
		res.R500(w, r, library.ErrInternalServerError)
		return
	}

	res.R200(w, r, test)
}

func (h *testHandler) Update(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	test, ok := ctx.Value(middlewares.TestCtxKey).(*models.Test)
	if !ok {
		res.R422(w, r, library.ErrUnprocessableEntity)
		return
	}

	var newTest models.Test
	if err := parse(r, &newTest); err != nil {
		log.Debug(err)
		res.R400(w, r, library.ErrBadRequest)
		return
	}

	if err := h.repository.Update(&newTest); err != nil {
		log.Debug(err)
		res.R500(w, r, library.ErrInternalServerError)
		return
	}
	res.R200(w, r, test)
}

func (h *testHandler) Get(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	test, ok := ctx.Value(middlewares.TestCtxKey).(*models.Test)
	if !ok {
		res.R422(w, r, library.ErrUnprocessableEntity)
		return
	}
	res.R200(w, r, test)
}

func (h *testHandler) List(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	query, ok := ctx.Value(middlewares.QueryCtxKey).(*library.QueryBuilder)
	if !ok {
		res.R422(w, r, library.ErrUnprocessableEntity)
		return
	}

	tests, total, err := h.repository.List(query, "")

	if err != nil {
		log.Debug(err)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			res.R404(w, r, library.ErrNotFound)
			return
		}
		res.R500(w, r, library.ErrInternalServerError)
		return
	}
	res.R200(w, r, map[string]interface{}{
		"total": total,
		"data":  tests,
	})
}

func (h *testHandler) ListByTestGroupID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	query, ok := ctx.Value(middlewares.QueryCtxKey).(*library.QueryBuilder)
	if !ok {
		res.R422(w, r, library.ErrUnprocessableEntity)
		return
	}

	testGroup, ok := ctx.Value(middlewares.TestGroupCtxKey).(*models.TestGroup)
	if !ok {
		res.R422(w, r, library.ErrUnprocessableEntity)
		return
	}

	tests, total, err := h.repository.List(query, "test_group_id=?", testGroup.ID)
	if err != nil {
		log.Debug(err)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			res.R404(w, r, library.ErrNotFound)
			return
		}
		res.R500(w, r, library.ErrInternalServerError)
		return
	}
	res.R200(w, r, map[string]interface{}{
		"total": total,
		"data":  tests,
	})
}

func (h *testHandler) Start(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	test, ok := ctx.Value(middlewares.TestCtxKey).(*models.Test)
	if !ok {
		res.R422(w, r, library.ErrUnprocessableEntity)
		return
	}
	runTest, err := h.service.Start(test)
	if err != nil {
		log.Debug(err)
		res.R500(w, r, library.ErrInternalServerError)
		return
	}

	res.R200(w, r, runTest)
}
