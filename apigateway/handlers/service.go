package handlers

import (
	"context"
	"net/http"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/s3f4/go-load/apigateway/library/log"
	res "github.com/s3f4/go-load/apigateway/library/response"
)

type serviceHandlerInterface interface {
	List(w http.ResponseWriter, r *http.Request)
}

type serviceHandler struct {
}

var (
	// ServiceHandler is used to show containers/services
	//and it can start and stop containers/services
	ServiceHandler serviceHandlerInterface = &serviceHandler{}
)

func (h *serviceHandler) List(w http.ResponseWriter, r *http.Request) {
	cli, err := client.NewEnvClient()
	if err != nil {
		log.Debug(err)
		res.R500(w, r, err)
		return
	}

	//Retrieve a list of containers
	services, err := cli.ServiceList(context.Background(), types.ServiceListOptions{})
	if err != nil {
		log.Debug(err)
		res.R500(w, r, err)
		return
	}

	res.R200(w, r, services)
}
