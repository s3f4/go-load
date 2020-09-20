package handlers

import (
	"net/http"

	"github.com/s3f4/go-load/apigateway/repository"
	. "github.com/s3f4/mu"
)

type statsHandlersInterface interface {
	List(w http.ResponseWriter, r *http.Request)
}

type statsHandler struct{}

var (
	// StatsHandler ...
	StatsHandler statsHandlersInterface = &statsHandler{}
)

func (sh *statsHandler) List(w http.ResponseWriter, r *http.Request) {
	rr := repository.NewResponseRepository()
	responses, err := rr.List(nil)
	if err != nil {
		R500(w, err)
		return
	}
	R200(w, responses)
}
