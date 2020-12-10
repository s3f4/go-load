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
	"gorm.io/gorm"
)

type runTestHandlerInterface interface {
	Create(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
	Get(w http.ResponseWriter, r *http.Request)
	List(w http.ResponseWriter, r *http.Request)
}

type runTestHandler struct {
	repository repository.RunTestRepository
}

var (
	//RunTestHandler .
	RunTestHandler runTestHandlerInterface = &runTestHandler{
		repository: repository.NewRunTestRepository(),
	}
)

func (h *runTestHandler) Create(w http.ResponseWriter, r *http.Request) {
	var runTest models.RunTest
	if err := parse(r, &runTest); err != nil {
		log.Error(err)
		res.R400(w, r, library.ErrBadRequest)
		return
	}

	err := h.repository.Create(&runTest)
	if err != nil {
		log.Info(err)
		res.R500(w, r, err)
		return
	}
	res.R200(w, r, runTest)
}

func (h *runTestHandler) Delete(w http.ResponseWriter, r *http.Request) {
	var runTest models.RunTest
	if err := parse(r, &runTest); err != nil {
		log.Info(err)
		res.R400(w, r, library.ErrBadRequest)
		return
	}

	err := h.repository.Delete(&runTest)
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
		res.R422(w, r, library.ErrUnprocessableEntity)
		return
	}
	res.R200(w, r, runTest)
}

func (h *runTestHandler) List(w http.ResponseWriter, r *http.Request) {
	runTest, err := h.repository.List()
	if err != nil {
		log.Info(err)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			res.R404(w, r, library.ErrNotFound)
			return
		}
		res.R500(w, r, err)
		return
	}
	res.R200(w, r, runTest)
}
