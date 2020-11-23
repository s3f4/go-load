package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
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

type testHandlerInterface interface {
	Create(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Get(w http.ResponseWriter, r *http.Request)
	List(w http.ResponseWriter, r *http.Request)
	Start(w http.ResponseWriter, r *http.Request)
}

type testHandler struct {
	service services.TestService
	tr      repository.TestRepository
}

var (
	//TestHandler is handler.
	TestHandler testHandlerInterface = &testHandler{
		service: services.NewTestService(),
		tr:      repository.NewTestRepository(),
	}
)

func (h *testHandler) Create(w http.ResponseWriter, r *http.Request) {
	var test models.Test
	if err := json.NewDecoder(r.Body).Decode(&test); err != nil {
		log.Debug(err)
		res.R400(w, r, library.ErrBadRequest)
		return
	}

	err := h.tr.Create(&test)
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
	if err := h.tr.Delete(test); err != nil {
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

	if err := h.tr.Update(test); err != nil {
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
	fmt.Printf("%#v", query)
	tests, total, err := h.tr.List(query)

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

	if err := h.service.Start(test); err != nil {
		log.Debug(err)
		res.R500(w, r, library.ErrInternalServerError)
		return
	}

	res.R200(w, r, "Test has been started.")
}
