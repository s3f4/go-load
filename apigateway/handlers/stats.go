package handlers

import (
	"errors"
	"net/http"

	"github.com/s3f4/go-load/apigateway/library"
	res "github.com/s3f4/go-load/apigateway/library/response"
	"github.com/s3f4/go-load/apigateway/middlewares"
	"github.com/s3f4/go-load/apigateway/models"
	"github.com/s3f4/go-load/apigateway/repository"
	"gorm.io/gorm"
)

type statsHandlersInterface interface {
	List(w http.ResponseWriter, r *http.Request)
}

type statsHandler struct {
	repository repository.ResponseRepository
}

var (
	// StatsHandler ...
	StatsHandler statsHandlersInterface = &statsHandler{
		repository: repository.NewResponseRepository(),
	}
)

func (h *statsHandler) List(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	runTest, ok := ctx.Value(middlewares.RunTestCtxKey).(*models.RunTest)
	if !ok {
		res.R422(w, r, library.ErrUnprocessableEntity)
		return
	}

	responses, err := h.repository.List(runTest.ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			res.R404(w, r, library.ErrNotFound)
			return
		}
		res.R500(w, r, err)
		return
	}
	res.R200(w, r, responses)
}
