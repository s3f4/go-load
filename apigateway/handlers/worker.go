package handlers

import "net/http"

type workerHandlerInterface interface {
	List(w http.ResponseWriter, r *http.Request)
}

type workerHandler struct{}

var (
	//WorkerHandler is handler.
	WorkerHandler workerHandlerInterface = &workerHandler{}
)

func (wh *workerHandler) List(w http.ResponseWriter, r *http.Request) {
	// todo list
}
