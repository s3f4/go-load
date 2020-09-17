package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/s3f4/go-load/apigateway/models"
	. "github.com/s3f4/mu"
	"github.com/s3f4/mu/log"
)

type workerHandlerInterface interface {
	Run(w http.ResponseWriter, r *http.Request)
	List(w http.ResponseWriter, r *http.Request)
	Stop(w http.ResponseWriter, r *http.Request)
}

type workerHandler struct{}

var (
	//WorkerHandler is handler.
	WorkerHandler workerHandlerInterface = &workerHandler{}
)

func (h *workerHandler) Run(w http.ResponseWriter, r *http.Request) {
	var run models.RunConfig
	if err := json.NewDecoder(r.Body).Decode(&run); err != nil {
		R400(w, "Bad Request")
		return
	}

}

func (h *workerHandler) Stop(w http.ResponseWriter, r *http.Request) {
	var worker models.Worker
	if err := json.NewDecoder(r.Body).Decode(&worker); err != nil {
		R400(w, "Bad Request")
		return
	}

	cli, err := client.NewEnvClient()
	if err != nil {
		log.Errorf("docker client err: %s", err)
	}

	if err := cli.ContainerStop(context.Background(), worker.ID, nil); err != nil {
		log.Errorf("docker client stop err: %s", err)
		R500(w, "internal server error")
		return
	}

	R200(w, "Container was stopped")
}

func (h *workerHandler) List(w http.ResponseWriter, r *http.Request) {
	cli, err := client.NewEnvClient()
	if err != nil {
		R500(w, "internal server error")
		return
	}

	//Retrieve a list of containers
	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{})
	if err != nil {
		R500(w, "internal server error")
		return
	}

	R200(w, map[string]interface{}{
		"containers": containers,
	})
}
