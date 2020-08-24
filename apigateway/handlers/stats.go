package handlers

import "net/http"

type statsHandlersInterface interface {
	Get(w http.ResponseWriter, r *http.Request)
}

type statsHandler struct{}

var (
	StatsHandler statsHandlersInterface = &statsHandler{}
)

func (sh *statsHandler) Get(w http.ResponseWriter, r *http.Request) {
	// todo Get
}
