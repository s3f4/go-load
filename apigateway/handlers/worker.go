package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	res "github.com/s3f4/go-load/apigateway/library/response"
	"github.com/s3f4/go-load/apigateway/models"
	"github.com/s3f4/mu/log"
)

type workerHandlerInterface interface {
	List(w http.ResponseWriter, r *http.Request)
	Stop(w http.ResponseWriter, r *http.Request)
}

type workerHandler struct {
}

var (
	// WorkerHandler is used to show containers/services
	//and it can start and stop containers/services
	WorkerHandler workerHandlerInterface = &workerHandler{}
)

func (h *workerHandler) Stop(w http.ResponseWriter, r *http.Request) {
	var worker models.Worker
	if err := json.NewDecoder(r.Body).Decode(&worker); err != nil {
		res.R400(w, r, "Bad Request")
		return
	}

	cli, err := client.NewEnvClient()
	if err != nil {
		log.Errorf("docker client err: %s", err)
	}

	if err := cli.ContainerStop(context.Background(), worker.ID, nil); err != nil {
		log.Errorf("docker client stop err: %s", err)
		res.R500(w, r, "internal server error")
		return
	}

	res.R200(w, r, "Container was stopped")
}

func (h *workerHandler) List(w http.ResponseWriter, r *http.Request) {
	cli, err := client.NewEnvClient()
	if err != nil {
		res.R500(w, r, "internal server error")
		return
	}

	//Retrieve a list of containers
	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{})
	if err != nil {
		res.R500(w, r, "internal server error")
		return
	}

	res.R200(w, r, map[string]interface{}{
		"containers": containers,
	})
}
