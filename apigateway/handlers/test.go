package handlers

import (
	"encoding/json"
	"net/http"

	res "github.com/s3f4/go-load/apigateway/library/response"
	"github.com/s3f4/go-load/apigateway/middlewares"
	"github.com/s3f4/go-load/apigateway/models"
	"github.com/s3f4/go-load/apigateway/repository"
	"github.com/s3f4/go-load/apigateway/services"
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
		res.R400(w, r, "Bad Request")
		return
	}

	err := h.tr.Create(&test)
	if err != nil {
		res.R500(w, r, err)
		return
	}
	res.R200(w, r, test)
}

func (h *testHandler) Delete(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	test, ok := ctx.Value(middlewares.TestCtxKey).(*models.Test)
	if !ok {
		res.R422(w, r, "unprocessable entity")
		return
	}
	if err := h.tr.Delete(test); err != nil {
		res.R500(w, r, err)
		return
	}

	res.R200(w, r, test)
}

func (h *testHandler) Update(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	test, ok := ctx.Value(middlewares.TestCtxKey).(*models.Test)
	if !ok {
		res.R422(w, r, "unprocessable entity")
		return
	}

	if err := h.tr.Update(test); err != nil {
		res.R500(w, r, err)
		return
	}
	res.R200(w, r, test)
}

func (h *testHandler) Get(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	test, ok := ctx.Value(middlewares.TestCtxKey).(*models.Test)
	if !ok {
		res.R422(w, r, "unprocessable entity")
		return
	}
	res.R200(w, r, test)
}

func (h *testHandler) List(w http.ResponseWriter, r *http.Request) {
	tests, err := h.tr.List()
	if err != nil {
		res.R500(w, r, err)
		return
	}
	res.R200(w, r, tests)
}

func (h *testHandler) Start(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	test, ok := ctx.Value(middlewares.TestCtxKey).(*models.Test)
	if !ok {
		res.R422(w, r, "unprocessable entity")
		return
	}

	if err := h.service.Start(test); err != nil {
		res.R500(w, r, err)
		return
	}

	res.R200(w, r, "Test has been started.")
}
