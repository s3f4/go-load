package handlers

import (
	"context"
	"net/http"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	. "github.com/s3f4/mu"
)

type workerHandlerInterface interface {
	List(w http.ResponseWriter, r *http.Request)
}

type workerHandler struct{}

var (
	//WorkerHandler is handler.
	WorkerHandler workerHandlerInterface = &workerHandler{}
)

func (wh *workerHandler) List(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}

	containers, err := cli.ContainerList(ctx, types.ContainerListOptions{})
	if err != nil {
		panic(err)
	}

	R200(w, containers)
}
