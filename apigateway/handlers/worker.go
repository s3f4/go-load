package handlers

import (
	"context"
	"net/http"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/go-chi/chi"
	. "github.com/s3f4/mu"
)

type workerHandlerInterface interface {
	List(w http.ResponseWriter, r *http.Request)
	Stop(w http.ResponseWriter, r *http.Request)
}

type workerHandler struct{}

var (
	//WorkerHandler is handler.
	WorkerHandler workerHandlerInterface = &workerHandler{}
)

func (wh *workerHandler) Stop(w http.ResponseWriter, r *http.Request) {
	containerID := chi.URLParam(r, "ID")

	cli, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}

	err = cli.ContainerStop(context.Background(), containerID, nil)
	if err != nil {
		panic(err)
	}
}

func (wh *workerHandler) List(w http.ResponseWriter, r *http.Request) {
	cli, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}

	//Retrieve a list of containers
	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{})
	if err != nil {
		panic(err)
	}

	R200(w, map[string]interface{}{
		"containers": containers,
	})
}
