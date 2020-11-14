package handlers

import (
	"net/http"

	res "github.com/s3f4/go-load/apigateway/library/response"
	"github.com/s3f4/go-load/apigateway/middlewares"
	"github.com/s3f4/go-load/apigateway/models"
	"github.com/s3f4/go-load/apigateway/repository"
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
		res.R422(w, r, "unprocessable entity")
		return
	}

	responses, err := h.repository.List(runTest.ID)
	if err != nil {
		res.R500(w, r, err)
		return
	}
	res.R200(w, r, responses)
}
