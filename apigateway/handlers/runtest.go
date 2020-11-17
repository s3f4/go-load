package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/s3f4/go-load/apigateway/library"
	"github.com/s3f4/go-load/apigateway/library/log"
	res "github.com/s3f4/go-load/apigateway/library/response"
	"github.com/s3f4/go-load/apigateway/middlewares"
	"github.com/s3f4/go-load/apigateway/models"
	"github.com/s3f4/go-load/apigateway/services"
)

type runTestHandlerInterface interface {
	Create(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
	Get(w http.ResponseWriter, r *http.Request)
	List(w http.ResponseWriter, r *http.Request)
}

type runTestHandler struct {
	service services.RunTestService
}

var (
	//RunTestHandler .
	RunTestHandler runTestHandlerInterface = &runTestHandler{
		service: services.NewRunTestService(),
	}
)

func (h *runTestHandler) Create(w http.ResponseWriter, r *http.Request) {
	var runTest models.RunTest
	if err := json.NewDecoder(r.Body).Decode(&runTest); err != nil {
		log.Info(err)
		res.R400(w, r, library.ErrBadRequest)
		return
	}

	err := h.service.Create(&runTest)
	if err != nil {
		log.Info(err)
		res.R500(w, r, err)
		return
	}
	res.R200(w, r, runTest)
}

func (h *runTestHandler) Delete(w http.ResponseWriter, r *http.Request) {
	var runTest models.RunTest
	if err := json.NewDecoder(r.Body).Decode(&runTest); err != nil {
		log.Info(err)
		res.R400(w, r, library.ErrBadRequest)
		return
	}

	err := h.service.Delete(&runTest)
	if err != nil {
		log.Info(err)
		res.R500(w, r, err)
		return
	}
	res.R200(w, r, runTest)
}

func (h *runTestHandler) Get(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	runTest, ok := ctx.Value(middlewares.RunTestCtxKey).(*models.RunTest)
	if !ok {
		res.R422(w, r, "unprocessable entity")
		return
	}
	res.R200(w, r, runTest)
}
func (h *runTestHandler) List(w http.ResponseWriter, r *http.Request) {
	runTest, err := h.service.List()
	if err != nil {
		log.Info(err)
		res.R500(w, r, err)
		return
	}
	res.R200(w, r, runTest)
}
