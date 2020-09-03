package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/s3f4/go-load/apigateway/models"
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
	var worker models.Worker
	err := json.NewDecoder(r.Body).Decode(&worker)
	fmt.Println(worker)

	cli, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}

	err = cli.ContainerStop(context.Background(), worker.Id, nil)
	if err != nil {
		R500(w, "internal server error")
		return
	}
	R200(w, "Container was stopped")
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
