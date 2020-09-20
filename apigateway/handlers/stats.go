package handlers

import (
	"net/http"

	"github.com/s3f4/go-load/apigateway/repository"
	. "github.com/s3f4/mu"
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
	responses, err := h.repository.List(nil)
	if err != nil {
		R500(w, err)
		return
	}
	R200(w, map[string]interface{}{"data": responses})
}
