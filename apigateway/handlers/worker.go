package handlers

import (
	"context"
	"net/http"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/s3f4/go-load/apigateway/library"
	"github.com/s3f4/go-load/apigateway/library/log"
	res "github.com/s3f4/go-load/apigateway/library/response"
	"github.com/s3f4/go-load/apigateway/models"
)

// WorkerHandler interface
type WorkerHandler interface {
	List(w http.ResponseWriter, r *http.Request)
	Stop(w http.ResponseWriter, r *http.Request)
}

type workerHandler struct {
}

// NewWorkerHandler returns new workerHandler object
func NewWorkerHandler() WorkerHandler {
	return &workerHandler{}
}

func (h *workerHandler) Stop(w http.ResponseWriter, r *http.Request) {
	var worker models.Worker
	if err := parse(r, &worker); err != nil {
		log.Debug(err)
		res.R400(w, r, library.ErrBadRequest)
		return
	}

	cli, err := client.NewEnvClient()
	if err != nil {
		log.Errorf("docker client err: %s", err)
		res.R500(w, r, library.ErrInternalServerError)
		return
	}

	if err := cli.ContainerStop(context.Background(), worker.ID, nil); err != nil {
		log.Errorf("docker client stop err: %s", err)
		res.R500(w, r, library.ErrInternalServerError)
		return
	}

	res.R200(w, r, "Container was stopped")
}

func (h *workerHandler) List(w http.ResponseWriter, r *http.Request) {
	cli, err := client.NewEnvClient()
	if err != nil {
		log.Debug(err)
		res.R500(w, r, library.ErrInternalServerError)
		return
	}

	//Retrieve a list of containers
	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{})
	if err != nil {
		log.Debug(err)
		res.R500(w, r, library.ErrInternalServerError)
		return
	}

	res.R200(w, r, map[string]interface{}{
		"containers": containers,
	})
}
